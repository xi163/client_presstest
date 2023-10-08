package handler

import (
	"strconv"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/json"
	"github.com/cwloo/gonet/utils/packet"
	"github.com/cwloo/gonet/utils/safe"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gameserv "github.com/cwloo/server/proto/game.serv"
)

func ReqPlayerLeave(peer conn.Session, userId int64, gameId, roomId, Type int32) {
	defer safe.Catch()
	reqdata := &gameserv.MSG_C2S_UserLeftMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.UserId = uint32(userId)
	reqdata.GameId = gameId
	reqdata.RoomId = roomId
	reqdata.Type = Type
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_USER_LEFT_REQ),
		reqdata)
	peer.Write(msg)
}
