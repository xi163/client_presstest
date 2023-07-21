package global

type Ctx struct {
	GameId   int32
	RoomId   int32
	Token    string
	HeadId   uint32
	Account  string
	NickName string
	UserId   int64
	Score    int64
	GamePass []byte
	AgentId  uint32
}
