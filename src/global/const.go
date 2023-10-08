package global

type GameMode int

// 0-匹配场 1-竞技场-比赛 2-俱乐部 3-好友房-朋友圈 4-房卡
const (
	Match   GameMode = 0 //匹配场
	Compete          = 1 //竞技场-比赛
	Club             = 2 //俱乐部
	Friend           = 3 //好友房-朋友圈
)
