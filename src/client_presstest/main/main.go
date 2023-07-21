package main

import (
	"sync"
	"time"

	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils/conv"
	"github.com/cwloo/presstest/src/client_presstest/client"
	"github.com/cwloo/presstest/src/client_presstest/entry"
	"github.com/cwloo/presstest/src/client_presstest/global"
	"github.com/cwloo/presstest/src/client_presstest/handler"
	"github.com/cwloo/presstest/src/client_presstest/service"
)

func ParallLoginRequest() {
	go func() {
		// timestart = time.Now().UnixMilli()
		for i := 0; i < global.TotalClients; i++ {
			account := global.BaseAccount + conv.IntToStr(i)
			token, ipaddr, domain := handler.GetPlatToken(global.Httpaddr, account, global.Type)
			if token == "" || ipaddr == "" {
				continue
			}
			//timenow = TimeNowMilliSec()
			// timdiff := TimeDiff(timenow, timestart)
			// if timdiff >= int32(*deltaTime) {
			// 	timestart = timenow
			// 	c := gSessMgr.Count()
			// 	delteConn := c - curConn
			// 	curConn = c
			// 	log.Printf("--- *** detla = %dms deltaClients = %03d", timdiff, delteConn)
			// }
			domain = "ws://192.168.0.113:10000/"
			c := client.New()
			c.Ctx.Token = token
			c.Ctx.Account = account
			if !global.Dynamic {
				c.Connect(global.Wsaddr)
			} else {
				if !global.Wssl {
					c.Connect(domain)
					// c.Connect("ws://" + ipaddr + "/")
				} else {
					// c.Connect("wss://" + ipaddr + "/")
					c.Connect(domain)
				}
			}
		}
	}()
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	service.Main.Add(time.Second, entry.NewSentryCreator(), 0, 10)
	ParallLoginRequest()
	wg.Wait()
	logs.Close()
}