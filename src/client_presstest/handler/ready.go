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

func ReqPlayerReady(peer conn.Session, userId int64) {
	defer safe.Catch()
	reqdata := &gameserv.MSG_C2S_UserReadyMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_USER_READY_REQ),
		reqdata)
	peer.Write(msg)
}
