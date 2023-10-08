package logic

import (
	"time"

	"github.com/cwloo/gonet/core/base/task"
	"github.com/cwloo/gonet/core/cb"
	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/presstest/src/client_presstest/client"
	"github.com/cwloo/presstest/src/client_presstest/handler"
	"github.com/cwloo/presstest/src/config"
	"github.com/cwloo/presstest/src/global"
)

func (s *Player) retry(peer conn.Session) {
	defer safe.Catch()
	switch config.Config.Client.Url {
	case "":
	default:
		cli := peer.GetContext("cli").(*client.Client)
		token, ipaddr, _ := handler.GetPlatToken(config.Config.Client.Url, cli.Ctx.Account, config.Config.Client.Type)
		if token == "" || ipaddr == "" {
			task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
				s.retry(peer)
			}))
			return
		}
		switch peer.Connected() {
		case true:
			ctx := peer.GetContext("ctx").(*global.Ctx)
			ctx.Token = token
			ctx.Header.Set("Sec-WebSocket-Verify", token)
			peer.GetContext("cli").(*client.Client).Ctx.Token = token
			peer.GetContext("cli").(*client.Client).Ctx.Header.Set("Sec-WebSocket-Verify", token)
			logs.Debugf("token >>> %s", ctx.Token)
			handler.SendGameLogin(peer, ctx.Token)
		default:
			cli.Ctx.Token = token
			cli.Ctx.Header.Set("Sec-WebSocket-Verify", token)
		}
		return
	}
	switch config.Config.Client.Router {
	case "":
	default:
		ipaddr, domain := handler.GetRouter(config.Config.Client.Router, "game_login", 1)
		if ipaddr == "" || domain == "" {
			task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
				s.retry(peer)
			}))
			return
		}
		cli := peer.GetContext("cli").(*client.Client)
		token, ipaddr, _ := handler.GetPlatToken(ipaddr, cli.Ctx.Account, config.Config.Client.Type)
		if token == "" || ipaddr == "" {
			task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
				s.retry(peer)
			}))
			return
		}
		switch peer.Connected() {
		case true:
			ctx := peer.GetContext("ctx").(*global.Ctx)
			ctx.Token = token
			ctx.Header.Set("Sec-WebSocket-Verify", token)
			peer.GetContext("cli").(*client.Client).Ctx.Token = token
			peer.GetContext("cli").(*client.Client).Ctx.Header.Set("Sec-WebSocket-Verify", token)
			logs.Debugf("token >>> %s", ctx.Token)
			handler.SendGameLogin(peer, ctx.Token)
		default:
			cli.Ctx.Token = token
			cli.Ctx.Header.Set("Sec-WebSocket-Verify", token)
		}
	}
}
