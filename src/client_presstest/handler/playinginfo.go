package handler

import (
	"strconv"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/json"
	"github.com/cwloo/gonet/utils/packet"
	"github.com/cwloo/gonet/utils/safe"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gamehall "github.com/cwloo/server/proto/game.hall"
)

// 查询正在玩的游戏/游戏服务器IP
func ReqPlayingGameInfo(peer conn.Session, userId int64) {
	defer safe.Catch()
	reqdata := &gamehall.GetPlayingGameInfoMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_PLAYING_GAME_INFO_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}
