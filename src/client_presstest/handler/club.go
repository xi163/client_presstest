package handler

import (
	"strconv"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/json"
	"github.com/cwloo/gonet/utils/packet"
	"github.com/cwloo/gonet/utils/safe"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gamehallclub "github.com/cwloo/server/proto/game.hallclub"
)

// 获取俱乐部房间信息
func ReqGetRoomInfoClub(peer conn.Session, userId, clubId int64, gameId, roomId int32) {
	defer safe.Catch()
	reqdata := &gamehallclub.GetRoomInfoMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.GameId = int32(gameId)
	reqdata.RoomId = int32(roomId)
	reqdata.ClubId = clubId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_GET_ROOM_INFO_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 获取我的俱乐部
func ReqGetMyClubListHall(peer conn.Session, userId int64) {
	defer safe.Catch()
	reqdata := &gamehallclub.GetMyClubHallMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_GET_MY_CLUB_HALL_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 创建俱乐部
func ReqCreateClub(peer conn.Session, userId int64, clubName string, ratio, autopartnerratio int) {
	defer safe.Catch()
	reqdata := &gamehallclub.CreateClubMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.ClubName = clubName
	reqdata.Rate = int32(ratio)
	reqdata.AutoBCPartnerRate = int32(autopartnerratio)
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_CREATE_CLUB_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 用户通过邀请码加入俱乐部
func ReqJoinClub(peer conn.Session, userId int64, invitationcode int32) {
	defer safe.Catch()
	reqdata := &gamehallclub.JoinTheClubMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.InvitationCode = invitationcode
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_JOIN_THE_CLUB_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 代理发起人邀请加入俱乐部
func ReqInviteJoinClub(peer conn.Session, userId, clubId int64, inviteUserId int64) {
	defer safe.Catch()
	reqdata := &gamehallclub.JoinTheClubMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.ClubId = clubId
	reqdata.UserId = inviteUserId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_JOIN_THE_CLUB_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 退出俱乐部
func ReqExitClub(peer conn.Session, userId, clubId int64) {
	defer safe.Catch()
	reqdata := &gamehallclub.ExitTheClubMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.ClubId = clubId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_EXIT_THE_CLUB_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}

// 代理发起人开除俱乐部成员
func ReqFireClubUser(peer conn.Session, userId int64, clubId int64, fireUserId int64) {
	defer safe.Catch()
	reqdata := &gamehallclub.FireMemberMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.ClubId = clubId
	reqdata.UserId = fireUserId
	logs.Debugf("%v %v", userId, json.String(reqdata))
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL_CLUB),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_CLUB_SUBID_CLIENT_TO_HALL_CLUB_FIRE_MEMBER_REQ),
		reqdata)
	peer.Write(msg)
}
