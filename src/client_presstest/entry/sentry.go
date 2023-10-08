package entry

import (
	"github.com/cwloo/gonet/core/base/run"
	"github.com/cwloo/gonet/core/base/run/cell"
	"github.com/cwloo/gonet/core/net/conn"
)

type sentry struct {
	cell.NetWorker
	main *smain
}

func newsentry(c run.Proc) cell.Worker {
	return &sentry{
		main: newsmain(c),
	}
}

func (s *sentry) OnInit() {
	s.main.initModuleHandlers()
}

func (s *sentry) OnConnected(peer conn.Session, v ...any) {
	s.main.onConnected(peer, v...)
}

func (s *sentry) OnClosed(peer conn.Session, reason conn.Reason, v ...any) {
	s.main.onClosed(peer, reason, v...)
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

func NewSentryCreator() cell.WorkerCreator {
	return &SentryCreator{}
}

func (s *SentryCreator) Create(c run.Proc, args ...any) cell.Worker {
	return newsentry(c)
}
