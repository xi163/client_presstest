package entry

import (
	"time"

	"github.com/cwloo/gonet/core/base/run"
	"github.com/cwloo/gonet/core/base/task"
	"github.com/cwloo/gonet/core/cb"
	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/packet"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/presstest/src/client_presstest/client"
	"github.com/cwloo/presstest/src/client_presstest/handler"
	"github.com/cwloo/presstest/src/client_presstest/logic"
	"github.com/cwloo/presstest/src/config"
	"github.com/cwloo/presstest/src/global"
)

type smain struct {
	c        run.Proc
	handlers cb.CmdCallbacks
	player   *logic.Player
}

func newsmain(c run.Proc) *smain {
	return &smain{
		c:        c,
		player:   logic.NewPlayer(c),
		handlers: cb.CmdCallbacks{},
	}
}

func (s *smain) onConnected(peer conn.Session, v ...any) {
	defer safe.Catch()
	c := v[0].(*client.Client)
	ctx := peer.GetContext("ctx").(*global.Ctx)
	ctx.Token = c.Ctx.Token
	ctx.Header.Set("Sec-WebSocket-Verify", c.Ctx.Token)
	//peer.WriteText("hello,world")
	handler.SendGameLogin(peer, ctx.Token)
}

func (s *smain) onClosed(peer conn.Session, reason conn.Reason, v ...any) {
	defer safe.Catch()
	ctx := peer.GetContext("ctx").(*global.Ctx)
	//peer.SetContext("cli", nil)
	//peer.SetContext("ctx", nil)
	//peer.SetContext("service", nil)
	switch ctx.Shutdown {
	case true:
		logs.Errorf("Shutdown")
	default:
		logs.Debugf("Reconnect[%v] -> [%v]", peer.LocalAddr(), peer.RemoteAddr())
		task.After(time.Duration(config.Config.Client.Interval[3])*time.Second, cb.NewFunctor00(func() {
			c := v[0].(*client.Client)
			c.Reconnect()
		}))
	}
}

func (s *smain) onRead(cmd uint32, msg any, peer conn.Session) {
	handler, ok := s.handlers[cmd]
	switch ok {
	case true:
		handler(msg, peer)
	default:
		mainId, subId := packet.Deword(int(cmd))
		logs.Errorf("unregistered cmd %v:%v", mainId, subId)
	}
}

func (s *smain) onCustom(cmd uint32, msg any, peer conn.Session) {
	handler, ok := s.handlers[cmd]
	switch ok {
	case true:
		handler(msg, peer)
	default:
		mainId, subId := packet.Deword(int(cmd))
		logs.Errorf("unregistered cmd %v:%v", mainId, subId)
	}
}

func (s *smain) OnTimer(timerID uint32, dt int32, args ...any) bool {
	return false
}

func (s *smain) initModuleHandlers() {
	s.player.RegisterModuleHandler(s.handlers)
}
