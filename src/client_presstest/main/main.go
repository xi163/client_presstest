package main

import (
	"runtime"
	"strings"
	"time"

	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/cmd"
	"github.com/cwloo/gonet/utils/console"
	"github.com/cwloo/gonet/utils/signal_handler"
	"github.com/cwloo/presstest/src/client_presstest/entry"
	"github.com/cwloo/presstest/src/client_presstest/service"
	"github.com/cwloo/presstest/src/config"
	"github.com/cwloo/presstest/src/global"
)

func init() {
	cmd.InitArgs(func(arg *cmd.ARG) {
		arg.SetConf("clientConfig_win/conf.ini")
	})
}

func onInput() {
	console.Read(func(str string) int {
		if str == "" {
			return 0
		}
		str = strings.ToLower(str)
		switch str[0] {
		case 'c':
			console.Clear[runtime.GOOS]()
		case 'q':
			console.Clear[runtime.GOOS]()
			signal_handler.Stop()
			return -1
		}
		return 0
	})
}

func main() {
	cmd.ParseArgs()
	config.InitClientConfig(cmd.Conf())
	logs.SetTimezone(logs.Timezone(config.Config.Log.Client.Timezone))
	logs.SetMode(logs.Mode(config.Config.Log.Client.Mode))
	logs.SetStyle(logs.Style(config.Config.Log.Client.Style))
	logs.SetLevel(logs.Level(config.Config.Log.Client.Level))
	logs.Init(config.Config.Log.Client.Dir, global.Exe, 100000000)
	start := time.Now()
	service.Main.Add(time.Second, entry.NewSentryCreator(), 0, config.Config.Client.Mailbox)
	logs.Debugf("Mailbox.Start = %d elapsed:%v", config.Config.Client.Mailbox, time.Since(start))
	ParallLoginRequest()
	go onInput()
	signal_handler.Wait(func() {
		logs.Close()
	})
}
