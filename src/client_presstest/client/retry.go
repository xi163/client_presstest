package client

import (
	"time"

	"github.com/cwloo/gonet/core/base/task"
	"github.com/cwloo/gonet/core/cb"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/presstest/src/client_presstest/handler"
	"github.com/cwloo/presstest/src/config"
)

func (s *Client) retry() {
	defer safe.Catch()
	s.Ctx.Retry_c++
	if s.Ctx.Retry_c < 3 {
		//三次内重连登陆当前网关 - 使用的是当前token进行重连登陆(可能掉线了，token未过期)
		//即便token过期了，只需重新拉取平台token并登陆当前网关，不用重连新的网关
		task.After(time.Duration(config.Config.Client.Interval[3])*time.Second, cb.NewFunctor00(func() {
			s.Reconnect()
		}))
	} else {
		//重连超过了三次，重新拉取平台token并连接登陆新的网关
		switch config.Config.Client.Url {
		case "":
		default:
			//拉取平台token 最小负载网关
			token, ipaddr, domain := handler.GetPlatToken(config.Config.Client.Url, s.Ctx.Account, config.Config.Client.Type)
			if token == "" || ipaddr == "" {
				task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
					s.retry()
				}))
				return
			}
			s.Ctx.Token = token
			s.Ctx.Header.Set("Sec-WebSocket-Verify", token)
			logs.Debugf("token >>> %s", s.Ctx.Token)
			s.Connect(s.Ctx.Header, domain) //连接网关并登陆
			return
		}
		switch config.Config.Client.Router {
		case "":
		default:
			//路由登陆节点
			ipaddr, domain := handler.GetRouter(config.Config.Client.Router, "game_login", 1)
			if ipaddr == "" || domain == "" {
				task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
					s.retry()
				}))
				return
			}
			//拉取平台token 最小负载网关
			token, ipaddr, domain := handler.GetPlatToken(ipaddr, s.Ctx.Account, config.Config.Client.Type)
			if token == "" || ipaddr == "" {
				task.After(time.Duration(config.Config.Client.Interval[2])*time.Second, cb.NewFunctor00(func() {
					s.retry()
				}))
				return
			}
			s.Ctx.Token = token
			s.Ctx.Header.Set("Sec-WebSocket-Verify", token)
			logs.Debugf("token >>> %s", s.Ctx.Token)
			s.Connect(s.Ctx.Header, domain) //连接网关并登陆
		}
	}
}
