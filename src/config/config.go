package config

import (
	"flag"
	"strconv"
	"strings"
	"sync"

	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/cmd"
	"github.com/cwloo/gonet/utils/conv"
	libini "github.com/cwloo/gonet/utils/ini"
	"github.com/cwloo/presstest/src/global"
)

var (
	lock               = &sync.RWMutex{}
	ini    *libini.Ini = &libini.Ini{}
	Config *IniConfig
)

type IniConfig struct {
	Flag     int `json:"flag" form:"flag"`
	Interval int `json:"interval" form:"interval"`
	Path     struct {
		UpdateCfg string `json:"updatecfg" form:"updatecfg"`
		GetCfg    string `json:"getcfg" form:"getcfg"`
	} `json:"path" form:"path"`
	Log struct {
		Monitor struct {
			Dir      string `json:"dir" form:"dir"`
			Level    int    `json:"level" form:"level"`
			Mode     int    `json:"mode" form:"mode"`
			Style    int    `json:"style" form:"style"`
			Timezone int    `json:"timezone" form:"timezone"`
		} `json:"monitor" form:"monitor"`
		Client struct {
			Dir      string `json:"dir" form:"dir"`
			Level    int    `json:"level" form:"level"`
			Mode     int    `json:"mode" form:"mode"`
			Style    int    `json:"style" form:"style"`
			Timezone int    `json:"timezone" form:"timezone"`
		} `json:"client" form:"client"`
	} `json:"log" form:"log"`
	Sub struct {
		Client struct {
			Num  int    `json:"num" form:"num"`
			Dir  string `json:"dir" form:"dir"`
			Exec string `json:"exec" form:"exec"`
		} `json:"client" form:"client"`
	} `json:"sub" form:"sub"`
	TgBot struct {
		Enable int    `json:"enable" form:"enable"`
		ChatId int64  `json:"chatId" form:"chatId"`
		Token  string `json:"token" form:"token"`
	} `json:"tg_bot" form:"tg_bot"`
	Monitor struct {
		Name        string `json:"name" form:"name"`
		Ip          string `json:"ip" form:"ip"`
		Port        []int  `json:"port" form:"port"`
		MaxConn     int    `json:"maxConn" form:"maxConn"`
		IdleTimeout int    `json:"idleTimeout" form:"idleTimeout"`
		Path        struct {
			Start   string `json:"start" form:"start"`
			Kill    string `json:"kill" form:"kill"`
			KillAll string `json:"killall" form:"killall"`
			SubList string `json:"sublist" form:"sublist"`
		} `json:"path" form:"path"`
	} `json:"monitor" form:"monitor"`
	Client struct {
		Name      string `json:"name" form:"name"`
		Mailbox   int    `json:"mailbox" form:"mailbox"`
		Router    string `json:"router" form:"router"`
		Url       string `json:"url" form:"url"`
		Dynamic   int    `json:"dynamic" form:"dynamic"`
		Wssl      int    `json:"wssl" form:"wssl"`
		Wsaddr    string `json:"wsaddr" form:"wsaddr"`
		Prefix    string `json:"prefix" form:"prefix"`
		BaseIndex int    `json:"base" form:"base"`
		Type      int    `json:"type" form:"type"`
		Decrypt   int    `json:"decrypt" form:"decrypt"`
		Clients   struct {
			Total int    `json:"total" form:"total"`
			Count [3]int `json:"count" form:"count"`
		} `json:"clients" form:"clients"`
		IdleTimeout int    `json:"idletimeout" form:"idletimeout"`
		Interval    [4]int `json:"interval" form:"interval"`
		AgentId     int    `json:"agentid" form:"agentid"`
		Md5Key      string `json:"md5code" form:"md5code"`
		AesKey      string `json:"descode" form:"descode"`
		GameId      int    `json:"gameid" form:"gameid"`
		RoomId      int    `json:"roomid" form:"roomid"`
	} `json:"client" form:"client"`
}

func readIni(filename string, cb func(*IniConfig) string) (c *IniConfig) {
	if err := ini.Load(filename); err != nil {
		logs.Fatalf(err.Error())
	}
	c = &IniConfig{}
	c.Flag = ini.GetInt("flag", "flag")
	s := ini.GetString("flag", "interval")
	sli := strings.Split(s, "*")
	va := 1
	for _, v := range sli {
		v = strings.ReplaceAll(v, " ", "")
		c, _ := strconv.Atoi(v)
		va *= c
	}
	c.Interval = va
	c.Path.UpdateCfg = ini.GetString("path", "updateconfig")
	c.Path.GetCfg = ini.GetString("path", "getconfig")
	c.Monitor.Name = ini.GetString("monitor", "name")
	c.Client.Name = ini.GetString("client", "name")
	setServiceName(cb, c)
	c.Log.Monitor.Dir = ini.GetString("log", "monitor.dir")
	c.Log.Monitor.Level = ini.GetInt("log", "monitor.level")
	c.Log.Monitor.Mode = ini.GetInt("log", "monitor.mode")
	c.Log.Monitor.Style = ini.GetInt("log", "monitor.style")
	c.Log.Monitor.Timezone = ini.GetInt("log", "monitor.timezone")
	c.Log.Client.Dir = ini.GetString("log", "client.dir")
	c.Log.Client.Level = ini.GetInt("log", "client.level")
	c.Log.Client.Mode = ini.GetInt("log", "client.mode")
	c.Log.Client.Style = ini.GetInt("log", "client.style")
	c.Log.Client.Timezone = ini.GetInt("log", "client.timezone")
	c.Sub.Client.Num = ini.GetInt("sub", "client.num")
	c.Sub.Client.Dir = ini.GetString("sub", "client.dir")
	c.Sub.Client.Exec = ini.GetString("sub", "client.execname")
	c.Monitor.Ip = ini.GetString("monitor", "ip")
	ports := strings.Split(ini.GetString("monitor", "port"), ",")
	for _, port := range ports {
		switch port == "" {
		case false:
			c.Monitor.Port = append(c.Monitor.Port, conv.StrToInt(port))
		}
	}
	c.Monitor.MaxConn = ini.GetInt("monitor", "maxConn")
	c.Monitor.IdleTimeout = ini.GetInt("monitor", "idleTimeout")
	c.Monitor.Path.Start = ini.GetString("path", "monitor.start")
	c.Monitor.Path.Kill = ini.GetString("path", "monitor.kill")
	c.Monitor.Path.KillAll = ini.GetString("path", "monitor.killall")
	c.Monitor.Path.SubList = ini.GetString("path", "monitor.sublist")
	c.Client.Mailbox = ini.GetInt("client", "mailboxs.num")
	c.Client.Router = ini.GetString("client", "router")
	c.Client.Url = ini.GetString("client", "url")
	c.Client.Dynamic = ini.GetInt("client", "dynamic")
	c.Client.Wssl = ini.GetInt("client", "wssl")
	c.Client.Wsaddr = ini.GetString("client", "wsaddr")
	c.Client.Prefix = ini.GetString("client", "prefix")
	c.Client.BaseIndex = ini.GetInt("client", "baseIndex")
	c.Client.Type = ini.GetInt("client", "type")
	c.Client.Decrypt = ini.GetInt("client", "decrypt")
	c.Client.Clients.Total = ini.GetInt("client", "totalClients")
	for i := range c.Client.Clients.Count {
		c.Client.Clients.Count[i] = ini.GetInt("client", "numClients"+conv.IntToStr(i+1))
	}
	c.Client.IdleTimeout = ini.GetInt("client", "idletimeout")
	for i := range c.Client.Interval {
		c.Client.Interval[i] = ini.GetInt("client", "interval"+conv.IntToStr(i))
	}
	c.Client.AgentId = ini.GetInt("client", "agentid")
	c.Client.Md5Key = ini.GetString("client", "md5code")
	c.Client.AesKey = ini.GetString("client", "descode")
	c.Client.GameId = ini.GetInt("client", "gameid")
	c.Client.RoomId = ini.GetInt("client", "roomid")
	return
}

func check() {
	switch ServiceName() {
	case Config.Monitor.Name:
		switch cmd.Log() == "" {
		case true:
			switch Config.Log.Monitor.Dir == "" {
			case true:
				Config.Log.Monitor.Dir = global.Dir + "log"
			default:
			}
		default:
			Config.Log.Monitor.Dir = cmd.Log()
		}
		switch Config.Log.Monitor.Timezone != int(logs.GetTimezone()) {
		case true:
			logs.SetTimezone(logs.Timezone(Config.Log.Monitor.Timezone))
		}
		switch Config.Log.Monitor.Mode != int(logs.GetMode()) {
		case true:
			logs.SetMode(logs.Mode(Config.Log.Monitor.Mode))
		}
		switch Config.Log.Monitor.Style != int(logs.GetStyle()) {
		case true:
			logs.SetStyle(logs.Style(Config.Log.Monitor.Style))
		}
		switch Config.Log.Monitor.Level != int(logs.GetLevel()) {
		case true:
			logs.SetLevel(logs.Level(Config.Log.Monitor.Level))
		}
	case Config.Client.Name:
		switch cmd.Log() == "" {
		case true:
			switch Config.Log.Client.Dir == "" {
			case true:
				Config.Log.Client.Dir = global.Dir + "log"
			default:
			}
		default:
			Config.Log.Client.Dir = cmd.Log()
		}
		switch Config.Log.Client.Timezone != int(logs.GetTimezone()) {
		case true:
			logs.SetTimezone(logs.Timezone(Config.Log.Client.Timezone))
		}
		switch Config.Log.Client.Mode != int(logs.GetMode()) {
		case true:
			logs.SetMode(logs.Mode(Config.Log.Client.Mode))
		}
		switch Config.Log.Client.Style != int(logs.GetStyle()) {
		case true:
			logs.SetStyle(logs.Style(Config.Log.Client.Style))
		}
		switch Config.Log.Client.Level != int(logs.GetLevel()) {
		case true:
			logs.SetLevel(logs.Level(Config.Log.Client.Level))
		}
	}
}

func read(conf string, cb func(*IniConfig) string) {
	Config = readIni(conf, cb)
	if Config == nil {
		logs.Fatalf("error")
	}
	// logs.Infof("%v %v", ServiceName(), conf)
}

func InitMonitorConfig(conf string) {
	read(conf, func(c *IniConfig) string {
		return c.Monitor.Name
	})
	switch Config.Flag {
	case 1:
		flag.Parse()
	default:
	}
	check()
}

func InitClientConfig(conf string) {
	read(conf, func(c *IniConfig) string {
		return c.Client.Name
	})
	switch Config.Flag {
	case 1:
		flag.Parse()
	default:
	}
	check()
}

func readConfig(conf string) {
	read(conf, nil)
	check()
}

func ReadConfig(conf string) {
	lock.RLock()
	readConfig(conf)
	lock.RUnlock()
}
