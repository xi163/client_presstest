package entry

import (
	"github.com/xi123/libgo/core/base/run"
	"github.com/xi123/libgo/core/base/run/cell"
	"github.com/xi123/libgo/core/net/conn"
)

type sentry struct {
	cell.NetWorker
	main *smain
}

func newsentry(c run.Proc) cell.Worker {
	p := &sentry{}
	p.main = newsmain(c)
	return p
}

func (s *sentry) OnInit() {
	s.main.initModuleHandlers()
}

func (s *sentry) OnRead(cmd uint32, msg any, peer conn.Session) {
	s.main.onRead(cmd, msg, peer)
}

func (s *sentry) OnCustom(cmd uint32, msg any, peer conn.Session) {
	s.main.onCustom(cmd, msg, peer)
}

func (s *sentry) OnTimer(timerID uint32, dt int32, args ...any) bool {
	return s.main.OnTimer(timerID, dt, args)
}

type SentryCreator struct {
	cell.NetWorkerCreator
}

func NewSentryCreator() *SentryCreator {
	return &SentryCreator{}
}

func (s *SentryCreator) Create(c run.Proc, args ...any) cell.Worker {
	return newsentry(c)
}
