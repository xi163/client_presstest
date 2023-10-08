package handler

import (
	"strconv"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/packet"
	gamecomm "github.com/cwloo/server/proto/game.comm"
	gamehall "github.com/cwloo/server/proto/game.hall"
)

// 登陆游戏大厅
func SendGameLogin(peer conn.Session, token string) {
	logs.Debugf(token)
	reqdata := &gamehall.LoginMessage{}
	val, _ := strconv.ParseUint("F5F5F5F5", 16, 32)
	reqdata.Header = &gamecomm.Header{}
	reqdata.Header.Sign = int32(val)
	reqdata.Session = token
	msg := packet.New(
		uint8(gamecomm.MAINID_MAIN_MESSAGE_CLIENT_TO_HALL),
		uint8(gamecomm.MESSAGE_CLIENT_TO_HALL_SUBID_CLIENT_TO_HALL_LOGIN_MESSAGE_REQ),
		reqdata)
	peer.Write(msg)
}
