package handler

import (
	"strconv"

	"github.com/xi163/libgo/core/net/conn"
	"github.com/xi163/libgo/logs"
	"github.com/xi163/libgo/utils/packet"
	"github.com/xi163/presstest/src/client_presstest/global"
	gamecomm "github.com/xi163/server/proto/game.comm"
	gamehall "github.com/xi163/server/proto/game.hall"
)

// 查询指定游戏节点
func ReqGameserverInfo(peer conn.Session, gameId, roomId int32) {
	ctx := peer.GetContext("ctx").(*global.Ctx)
	logs.Debugf("uid:%v %v %v", ctx.UserId, gameId, roomId)
	reqdata := &gamehall.GetGameServerMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.GameId = uint32(gameId)
	reqdata.RoomId = uint32(roomId)
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_SERVER_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}
