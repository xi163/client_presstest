package global

type RouterParam struct {
	Node      string `json:"node" form:"node"`
	Type      int    `json:"type" form:"type"`
	Timestamp int64  `json:"timestamp" form:"timestamp"`
}

type LoginParam struct {
	Account   string `json:"account" form:"account"`
	Type      int    `json:"type" form:"type"`
	Timestamp int64  `json:"timestamp" form:"timestamp"`
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

type RouterResult struct {
	Node string       `json:"node" form:"node"`
	Type int          `json:"type" form:"type"`
	Data []ServerLoad `json:"data" form:"data"`
}
