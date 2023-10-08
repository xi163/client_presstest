package global

import "net/http"

type Ctx struct {
	Header   http.Header
	Retry_c  int
	GameId   uint32
	RoomId   uint32
	Token    string
	HeadId   uint32
	Account  string
	NickName string
	UserId   int64
	Score    int64
	GamePass []byte
	AgentId  uint32
	Shutdown bool
	ClubId   int64
	Type     GameMode
}
