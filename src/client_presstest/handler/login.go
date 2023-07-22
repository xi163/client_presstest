package handler

import (
	"strconv"

	"github.com/xi163/libgo/core/net/conn"
	"github.com/xi163/libgo/logs"
	"github.com/xi163/libgo/utils/packet"
	gamecomm "github.com/xi163/server/proto/game.comm"
	gamehall "github.com/xi163/server/proto/game.hall"
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
