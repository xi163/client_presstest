package main

import (
	"github.com/cwloo/gonet/utils/conv"
	"github.com/cwloo/presstest/src/config"
)

func ParallLoginRequest() {
	go func() {
		// timestart = time.Now().UnixMilli()
		for i := 0; i < config.Config.Client.Clients.Total; i++ {
			account := config.Config.Client.Prefix + conv.IntToStr(config.Config.Client.BaseIndex+i)
			go retry(account)
			//timenow = TimeNowMilliSec()
			// timdiff := TimeDiff(timenow, timestart)
			// if timdiff >= int32(*deltaTime) {
			// 	timestart = timenow
			// 	c := gSessMgr.Count()
			// 	delteConn := c - curConn
			// 	curConn = c
			// 	log.Printf("--- *** detla = %dms deltaClients = %03d", timdiff, delteConn)
			// }
		}
	}()
}
