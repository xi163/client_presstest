package client

import (
	"encoding/binary"
	"net/http"
	"time"

	"github.com/cwloo/gonet/core/base/pipe"
	"github.com/cwloo/gonet/core/base/task"
	"github.com/cwloo/gonet/core/cb"
	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/core/net/tcp/tcpclient"
	"github.com/cwloo/gonet/core/net/transmit"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/packet"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/gonet/utils/timestamp"
	"github.com/cwloo/presstest/src/client_presstest/handler"
	"github.com/cwloo/presstest/src/client_presstest/service"
	"github.com/cwloo/presstest/src/config"
	"github.com/cwloo/presstest/src/global"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gameserv "github.com/cwloo/server/proto/game.serv"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Ctx    global.Ctx
	client tcpclient.TCPClient
}

func New(addr ...string) *Client {
	s := &Client{Ctx: global.Ctx{Header: http.Header{}}}
	s.Ctx.Header.Set("Origin", "*")
	s.Ctx.Header.Set("Sec-WebSocket-Protocol", "chat")
	//s.Ctx.Header.Set("Sec-WebSocket-Key", "")
	s.client = tcpclient.NewTCPClient("Client", addr...)
	s.client.SetProtocolCallback(s.onProtocol)
	s.client.SetConnectErrorCallback(s.onConnectError)
	s.client.SetConnectedCallback(s.onConnected)
	s.client.SetClosedCallback(s.onClosed)
	s.client.SetMessageCallback(s.onMessage)
	s.client.SetDialTimeout(10 * time.Second)
	s.client.SetIdleTimeout(
		time.Duration(config.Config.Client.IdleTimeout)*time.Second,
		time.Duration(config.Config.Client.Interval[0])*time.Second)
	//s.client.EnableRetry(true)
	s.client.SetRetryInterval(10 * time.Second)
	s.client.SetHoldType(conn.KHoldNone)
	return s
}

func (s *Client) Connect(header http.Header, addr ...string) {
	s.client.ConnectTCP(header, addr...)
}

func (s *Client) Reconnect() {
	s.client.Reconnect()
}

func (s *Client) onProtocol(proto string) transmit.Channel {
	switch proto {
	case "tcp":
		panic("tcp Channel undefine")
	case "ws", "wss":
		return packet.NewWSChannel()
	}
	panic("no proto setup")
}

func (s *Client) onConnectError(proto string, err error) {
	logs.Errorf(err.Error())
	s.retry()
}

func (s *Client) onConnected(peer conn.Session, v ...any) {
	defer safe.Catch()
	if peer.Connected() {
		logs.Infof("[%v] -> [%v]", peer.LocalAddr(), peer.RemoteAddr())
		s.Ctx.Retry_c = 0
		peer.SetContext("cli", s)
		ctx := global.Ctx{Account: s.Ctx.Account, Token: s.Ctx.Token, Header: s.Ctx.Header}
		peer.SetContext("ctx", &ctx)
		peer.SetContext("service", service.Main.Next())
		peer.GetContext("service").(pipe.Pipe).Slot().Proc().PostConnected(peer, s)
	} else {
		panic("error")
	}
}

func (s *Client) onClosed(peer conn.Session, reason conn.Reason, v ...any) {
	defer safe.Catch()
	if peer.Connected() {
		panic("error")
	} else {
		logs.Tracef("[%v] -> [%v] %v", peer.LocalAddr(), peer.RemoteAddr(), reason.Msg)
		peer.GetContext("service").(pipe.Pipe).Slot().Proc().PostClosed(peer, reason, s)
	}
}

func (s *Client) onMessage(peer conn.Session, msg any, msgType int, recvTime timestamp.T) {
	switch msgType {
	case websocket.TextMessage:
		switch msg := msg.(type) {
		case []byte:
			logs.Debugf("%s", string(msg))
		case string:
			logs.Debugf("%s", msg)
		default:
			panic("error")
		}
	case websocket.BinaryMessage:
		switch msg := msg.(type) {
		case []byte:
			cmd, data, err := packet.Unpack(msg, binary.LittleEndian)
			if err != nil {
				logs.Errorf(err.Error())
				peer.Close()
				return
			}
			s.onData(peer, cmd, data)
		default:
			panic("error")
		}
	}
}

func (s *Client) onData(peer conn.Session, cmd uint32, data []byte) {
	defer safe.Catch()
	if cmd == uint32(packet.Enword(int(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL), int(gamecomm.MESSAGE_CLIENT_TO_SERVER_SUBID_KEEP_ALIVE_RES))) {
		//大厅服心跳包
		handler.ResultKeepAlive(data, peer)
		ctx := peer.GetContext("ctx").(*global.Ctx)
		task.After(time.Duration(config.Config.Client.Interval[0])*time.Second, cb.NewFunctor10(func(args any) {
			handler.SendKeepAlive(peer, args.(string), false)
		}, ctx.Token))
	} else if cmd == uint32(packet.Enword(int(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER), int(gamecomm.MESSAGE_CLIENT_TO_SERVER_SUBID_KEEP_ALIVE_RES))) {
		//游戏服心跳包
		handler.ResultKeepAlive(data, peer)
		ctx := peer.GetContext("ctx").(*global.Ctx)
		task.After(time.Duration(config.Config.Client.Interval[0])*time.Second, cb.NewFunctor10(func(args any) {
			handler.SendKeepAlive(peer, args.(string), true)
		}, ctx.Token))
	} else {
		mainId, _ := packet.Deword(int(cmd))
		if mainId == int(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_LOGIC) {
			//游戏逻辑服消息
			msg := gameserv.MSG_CSC_Passageway{}
			err := proto.Unmarshal(data, &msg)
			if err != nil {
				logs.Errorf(err.Error())
				peer.Close()
				return
			}
			peer.GetContext("service").(pipe.Pipe).Slot().Proc().PostRead(cmd, msg.PassData, peer)
		} else {
			//框架消息 网关服/大厅服/游戏服
			peer.GetContext("service").(pipe.Pipe).Slot().Proc().PostRead(cmd, data, peer)
		}
	}
}
