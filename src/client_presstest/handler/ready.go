package handler

import (
	"strconv"

	"github.com/xi123/libgo/core/net/conn"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils/packet"
	"github.com/cwloo/presstest/src/client_presstest/global"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gameserv "github.com/cwloo/server/proto/game.serv"
)

func ReqPlayerReady(peer conn.Session) {
	ctx := peer.GetContext("ctx").(*global.Ctx)
	logs.Debugf("uid:%v", ctx.UserId)
	reqdata := &gameserv.MSG_C2S_UserReadyMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_USER_READY_REQ),
		reqdata)
	peer.Write(msg)
}
