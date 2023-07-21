package handler

import (
	"strconv"

	"github.com/xi123/libgo/core/net/conn"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils/packet"
	"github.com/xi123/presstest/src/client_presstest/global"
	gamecomm "github.com/xi123/server/proto/game.comm"
	gamehall "github.com/xi123/server/proto/game.hall"
)

// 查询游戏房间列表
func ReqGameListInfo(peer conn.Session) {
	ctx := peer.GetContext("ctx").(*global.Ctx)
	logs.Debugf("uid:%v", ctx.UserId)
	reqdata := &gamehall.GetGameMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_ROOM_INFO_REQ),
		reqdata)
	peer.Write(msg)
}
