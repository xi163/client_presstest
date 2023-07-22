package client

import (
	"encoding/binary"
	"time"

	"github.com/xi163/libgo/core/base/pipe"
	"github.com/xi163/libgo/core/base/task"
	"github.com/xi163/libgo/core/cb"
	"github.com/xi163/libgo/core/net/conn"
	"github.com/xi163/libgo/core/net/tcp/tcpclient"
	"github.com/xi163/libgo/core/net/transmit"
	"github.com/xi163/libgo/logs"
	"github.com/xi163/libgo/utils/packet"
	"github.com/xi163/libgo/utils/timestamp"
	"github.com/xi163/presstest/src/client_presstest/global"
	"github.com/xi163/presstest/src/client_presstest/handler"
	"github.com/xi163/presstest/src/client_presstest/service"
	gamecomm "github.com/xi163/server/proto/game.comm"
	gameserv "github.com/xi163/server/proto/game.serv"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Ctx    global.Ctx
	client tcpclient.TCPClient
}

func New(addr ...string) *Client {
	s := &Client{}
	s.client = tcpclient.NewTCPClient("Client", addr...)
	s.client.SetProtocolCallback(s.onProtocol)
	s.client.SetConnectedCallback(s.onConnected)
	s.client.SetClosedCallback(s.onClosed)
	s.client.SetMessageCallback(s.onMessage)
	s.client.SetDialTimeout(10 * time.Second)
	s.client.SetIdleTimeout(global.IdleTimeout, global.Interval)
	s.client.EnableRetry(false)
	s.client.SetRetryInterval(10 * time.Second)
	s.client.SetHoldType(conn.KHoldNone)
	return s
}

func (s *Client) Connect(addr ...string) {
	s.client.ConnectTCP(addr...)
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

func (s *Client) onConnected(peer conn.Session, v ...any) {
	if peer.Connected() {
		logs.Infof("[%v] -> [%v]", peer.LocalAddr(), peer.RemoteAddr())
		peer.SetContext("ctx", &s.Ctx)
		peer.SetContext("service", service.Main.Next())
		handler.SendGameLogin(peer, s.Ctx.Token)
	} else {
		panic("error")
	}
}

func (s *Client) onClosed(peer conn.Session, reason conn.Reason) {
	if peer.Connected() {
		panic("error")
	} else {
		logs.Tracef("[%v] -> [%v] %v", peer.LocalAddr(), peer.RemoteAddr(), reason.Msg)
		peer.SetContext("ctx", nil)
		peer.SetContext("service", nil)
	}
}

func (s *Client) onMessage(peer conn.Session, msg any, recvTime timestamp.T) {
	switch b := msg.(type) {
	case []byte:
		cmd, data, err := packet.Unpack(b, binary.LittleEndian)
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

func (s *Client) onData(peer conn.Session, cmd uint32, data []byte) {
	if cmd == uint32(packet.Enword(int(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL), int(gamecomm.MESSAGE_CLIENT_TO_SERVER_SUBID_KEEP_ALIVE_RES))) ||
		cmd == uint32(packet.Enword(int(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER), int(gamecomm.MESSAGE_CLIENT_TO_SERVER_SUBID_KEEP_ALIVE_RES))) {
		//大厅服，游戏服心跳包
		handler.ResultKeepAlive(data, peer)
		task.After(global.Interval, cb.NewFunctor00(func() {
			handler.SendKeepAlive(peer, s.Ctx.Token)
		}))
	} else {
		mainId, _ := packet.Deword(int(cmd))
		if mainId == int(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_LOGIC) {
			//子游戏逻辑消息
			msg := gameserv.MSG_CSC_Passageway{}
			err := proto.Unmarshal(data, &msg)
			if err != nil {
				panic(err.Error())
			}
			peer.GetContext("service").(pipe.Pipe).Slot().Proc().PostRead(cmd, msg.PassData, peer)
		} else {
			//框架消息 网关服/大厅服/游戏服
			peer.GetContext("service").(pipe.Pipe).Slot().Proc().PostRead(cmd, data, peer)
		}
	}
}
