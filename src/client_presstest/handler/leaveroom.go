package handler

import (
	"strconv"

	"github.com/xi163/libgo/core/net/conn"
	"github.com/xi163/libgo/logs"
	"github.com/xi163/libgo/utils/packet"
	"github.com/xi163/presstest/src/client_presstest/global"
	gamecomm "github.com/xi163/server/proto/game.comm"
	gameserv "github.com/xi163/server/proto/game.serv"
)

func ReqPlayerLeave(peer conn.Session, userId int64, gameId, roomId, Type int32) {
	ctx := peer.GetContext("ctx").(*global.Ctx)
	logs.Debugf("uid:%v %v %v", ctx.UserId, gameId, roomId)
	reqdata := &gameserv.MSG_C2S_UserLeftMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.UserId = uint32(userId)
	reqdata.GameId = gameId
	reqdata.RoomId = roomId
	reqdata.Type = Type
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_USER_LEFT_REQ),
		reqdata)
	peer.Write(msg)
}
