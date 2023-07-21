package global

import "time"

var Md5Key = "334270F58E3E9DEC"
var AesKey = "111362EE140F157D"

var TotalClients = 1
var Httpaddr = "http://192.168.0.105:9787"
var BaseAccount = "test_"
var Type = 0
var Dynamic = true
var Wssl = false
var Wsaddr string

var IdleTimeout = 60 * time.Second
var Interval = 10 * time.Second

var SubGameId uint32 = 630
var SubRoomId uint32 = 6301
