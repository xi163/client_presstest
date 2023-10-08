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

// 查询指定游戏节点/游戏服务器IP
func ReqGameserverInfo(peer conn.Session, userId int64, gameId, roomId int32) {
	defer safe.Catch()
	reqdata := &gamehall.GetGameServerMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Type = int32(global.Match)
	reqdata.GameId = uint32(gameId)
	reqdata.RoomId = uint32(roomId)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_SERVER_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 获取游戏服务器IP
func ReqGameserverInfoClub(peer conn.Session, userId, clubId int64, gameId, roomId int32) {
	defer safe.Catch()
	reqdata := &gamehall.GetGameServerMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Type = int32(global.Club)
	reqdata.ClubId = clubId
	reqdata.GameId = uint32(gameId)
	reqdata.RoomId = uint32(roomId)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_SERVER_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

func ReqValidateGameserverInfoClub(peer conn.Session, userId, clubId int64, gameId, roomId int32, servId string, tableId int32) {
	defer safe.Catch()
	reqdata := &gamehall.GetGameServerMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Type = int32(global.Club)
	reqdata.ClubId = clubId
	reqdata.GameId = uint32(gameId)
	reqdata.RoomId = uint32(roomId)
	reqdata.TableId = uint32(tableId)
	reqdata.ServId = servId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_GET_GAME_SERVER_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}
