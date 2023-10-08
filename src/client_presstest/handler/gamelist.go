package handler

import (
	"strconv"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/json"
	"github.com/cwloo/gonet/utils/packet"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/presstest/src/global"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gamehall "github.com/cwloo/server/proto/game.hall"
)

// 获取所有游戏列表/匹配场
func ReqGameListInfo(peer conn.Session, userId int64) {
	defer safe.Catch()
	reqdata := &gamehall.GetGameMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Type = int32(global.Match)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_ROOM_INFO_REQ),
		reqdata)
	peer.Write(msg)
}

// 获取所有游戏列表/俱乐部
func ReqGameListInfoClub(peer conn.Session, userId, clubId int64) {
	defer safe.Catch()
	reqdata := &gamehall.GetGameMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Type = int32(global.Club)
	reqdata.ClubId = clubId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_ROOM_INFO_REQ),
		reqdata)
	peer.Write(msg)
}
