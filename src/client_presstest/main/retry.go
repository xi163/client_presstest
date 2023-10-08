package main

import (
	"time"

	"github.com/cwloo/gonet/core/base/task"
	"github.com/cwloo/gonet/core/cb"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/presstest/src/client_presstest/client"
	"github.com/cwloo/presstest/src/client_presstest/handler"
	"github.com/cwloo/presstest/src/config"
)

func retry(account string) {
	defer safe.Catch()
	switch config.Config.Client.Url {
	case "":
	default:
		token, ipaddr, domain := handler.GetPlatToken(config.Config.Client.Url, account, config.Config.Client.Type)
		if token == "" || ipaddr == "" {
			task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
				retry(account)
			}))
			return
		}
		s := client.New()
		s.Ctx.Account = account
		s.Ctx.Token = token
		s.Ctx.Header.Set("Sec-WebSocket-Verify", token)
		if config.Config.Client.Dynamic == 0 {
			s.Connect(s.Ctx.Header, config.Config.Client.Wsaddr)
		} else {
			if config.Config.Client.Wssl == 0 {
				s.Connect(s.Ctx.Header, domain)
			} else {
				s.Connect(s.Ctx.Header, domain)
			}
		}
		return
	}
	switch config.Config.Client.Router {
	case "":
	default:
		ipaddr, domain := handler.GetRouter(config.Config.Client.Router, "game_login", 1)
		if ipaddr == "" || domain == "" {
			task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
				retry(account)
			}))
			return
		}
		token, ipaddr, domain := handler.GetPlatToken(ipaddr, account, config.Config.Client.Type)
		if token == "" || ipaddr == "" {
			task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
				retry(account)
			}))
			return
		}
		s := client.New()
		s.Ctx.Account = account
		s.Ctx.Token = token
		s.Ctx.Header.Set("Sec-WebSocket-Verify", token)
		if config.Config.Client.Dynamic == 0 {
			s.Connect(s.Ctx.Header, config.Config.Client.Wsaddr)
		} else {
			if config.Config.Client.Wssl == 0 {
				s.Connect(s.Ctx.Header, domain)
			} else {
				s.Connect(s.Ctx.Header, domain)
			}
		}
	}
}
