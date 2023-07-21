package service

import (
	"time"

	"github.com/xi123/libgo/core/base/mailbox"
	"github.com/xi123/libgo/core/base/pipe"
	"github.com/xi123/libgo/core/base/run/cell"
)

var Main = NewService("client.service")

type Service struct {
	m mailbox.Pipes
}

func NewService(name string) *Service {
	return &Service{
		m: mailbox.NewPipes(name),
	}
}

func (s *Service) Add(d time.Duration, creator cell.WorkerCreator, size, num int) {
	s.m.Add(d, creator, size, num)
}

func (s *Service) Next() (pipe pipe.Pipe) {
	return s.m.Next()
}
