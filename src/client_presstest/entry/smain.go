package entry

import (
	"github.com/xi123/libgo/core/base/run"
	"github.com/xi123/libgo/core/cb"
	"github.com/xi123/libgo/core/net/conn"
	"github.com/xi123/libgo/logs"
	"github.com/cwloo/presstest/src/client_presstest/logic"
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

func (s *smain) onRead(cmd uint32, msg any, peer conn.Session) {
	if handler, ok := s.handlers[cmd]; ok {
		handler(msg, peer)
	} else {
		logs.Errorf("cmd:%v not exist", cmd)
	}
}

func (s *smain) onCustom(cmd uint32, msg any, peer conn.Session) {
	if handler, ok := s.handlers[cmd]; ok {
		handler(msg, peer)
	} else {
		logs.Errorf("cmd:%v not exist", cmd)
	}
}

func (s *smain) OnTimer(timerID uint32, dt int32, args ...any) bool {
	return false
}

func (s *smain) initModuleHandlers() {
	logs.Debugf("...")
	s.player.RegisterModuleHandler(s.handlers)
}
