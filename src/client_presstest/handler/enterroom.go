package handler

import (
	"strconv"

	"github.com/xi123/libgo/core/net/conn"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils/packet"
	"github.com/xi123/presstest/src/client_presstest/global"
	gamecomm "github.com/xi123/server/proto/game.comm"
	gameserv "github.com/xi123/server/proto/game.serv"
)

func ReqEnterRoom(peer conn.Session, gameId, roomId int32, pwd []byte) {
	ctx := peer.GetContext("ctx").(*global.Ctx)
	logs.Debugf("uid:%v %v %v %v", ctx.UserId, gameId, roomId, string(pwd))
	reqdata := &gameserv.MSG_C2S_UserEnterMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.GameId = gameId
	reqdata.RoomId = roomId
	reqdata.DynamicPassword = pwd[:]
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_ENTER_ROOM_REQ),
		reqdata)
	peer.Write(msg)
}
