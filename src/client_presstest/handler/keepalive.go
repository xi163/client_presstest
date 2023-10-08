package handler

import (
	"strconv"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/json"
	"github.com/cwloo/gonet/utils/packet"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	"google.golang.org/protobuf/proto"
)

// 发送心跳包
func SendKeepAlive(peer conn.Session, token string, v bool) {
	reqdata := &gamecomm.KeepAliveMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Session = token
	switch v {
	case true:
		// logs.Debugf("[game]%s", token)
		msg := packet.New(
			uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
			uint8(gamecomm.MESSAGE_CLIENT_TO_SERVER_SUBID_KEEP_ALIVE_REQ),
			reqdata)
		peer.Write(msg)
	default:
		//logs.Debugf("[hall]%s", token)
		msg := packet.New(
			uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
			uint8(gamecomm.MESSAGE_CLIENT_TO_SERVER_SUBID_KEEP_ALIVE_REQ),
			reqdata)
		peer.Write(msg)
	}
}

func ResultKeepAlive(msg any, peer conn.Session) {
	switch b := msg.(type) {
	case []byte:
		rspdata := gamecomm.KeepAliveMessageResponse{}
		err := proto.Unmarshal(b, &rspdata)
		if err != nil {
			panic(err.Error())
		}
		if rspdata.RetCode == 0 {
		} else {
			logs.Errorf("%v", json.String(&rspdata))
			// peer.Close()
		}
	}
}
