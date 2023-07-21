package global

type LoginParam struct {
	Account   string
	Type      int
	Timestamp int64
}

type Request struct {
	Key   string `json:"key" form:"key"`
	Param string `json:"param" form:"param"`
}

type LoginResult struct {
	Account string       `json:"account" form:"account"`
	Userid  int64        `json:"userid" form:"userid"`
	Data    []ServerLoad `json:"data" form:"data"`
}

type ServerLoad struct {
	Host       string `json:"host" form:"host"`
	Domain     string `json:"domain" form:"domain"`
	NumOfLoads int    `json:"numOfLoads" form:"numOfLoads"`
}
