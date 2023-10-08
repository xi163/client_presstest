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

func ReqEnterRoom(peer conn.Session, userId int64, gameId, roomId uint32, pwd []byte) {
	defer safe.Catch()
	reqdata := &gameserv.MSG_C2S_UserEnterMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.GameId = gameId
	reqdata.RoomId = roomId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	reqdata.DynamicPassword = pwd[:]
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_ENTER_ROOM_REQ),
		reqdata)
	peer.Write(msg)
}

func ReqEnterRoomClub(peer conn.Session, userId, clubId int64, gameId, roomId, tableId uint32, servId string, pwd []byte) {
	defer safe.Catch()
	reqdata := &gameserv.MSG_C2S_UserEnterMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.ClubId = clubId
	reqdata.GameId = gameId
	reqdata.RoomId = roomId
	reqdata.ServId = servId
	reqdata.TableId = tableId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	reqdata.DynamicPassword = pwd[:]
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_GAME_SERVER),
		uint8(gameserv.SUBID_SUB_C2S_ENTER_ROOM_REQ),
		reqdata)
	peer.Write(msg)
}
