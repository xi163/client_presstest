package handler

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/gonet/utils/codec/base64"
	"github.com/cwloo/gonet/utils/codec/uri"
	"github.com/cwloo/gonet/utils/crypto/aes"
	"github.com/cwloo/gonet/utils/crypto/md5"
	"github.com/cwloo/gonet/utils/httpcli"
	"github.com/cwloo/gonet/utils/json"
	"github.com/cwloo/gonet/utils/result"
	"github.com/cwloo/gonet/utils/safe"
	"github.com/cwloo/gonet/utils/sign"
	"github.com/cwloo/gonet/utils/validator"
	"github.com/cwloo/presstest/src/config"
	"github.com/cwloo/presstest/src/global"
)

func BuildRouter(node string, Type int) (key, param string) {
	req := global.RouterParam{
		Node:      node,
		Type:      Type,
		Timestamp: time.Now().Unix(),
	}
	key = md5.Md5(fmt.Sprintf("%v%v%v%v", req.Node, req.Type, req.Timestamp, config.Config.Client.Md5Key), true)
	rawParam := json.String(req)
	// logs.Infof("rawParam >>>>  %s", rawParam)
	encrypt := aes.ECBEncryptPKCS7([]byte(rawParam), []byte(config.Config.Client.AesKey), []byte(config.Config.Client.AesKey))
	strBase64 := base64.URLEncode(encrypt)
	// logs.Infof("strBase64 >>>>  %s", strBase64)
	param = uri.URLEncode(strBase64)
	// logs.Infof("URLEncode >>>>  %s", param)
	return
}

func SendRouter(httpaddr string, key, param string) (ipaddr, domain string) {
	defer safe.Catch()
	vec := strings.Split(httpaddr, "//")
	if len(vec) != 2 {
		return
	}
	proto := strings.Trim(vec[0], ":")
	host := vec[1]
	requrl := fmt.Sprintf("%v://%v/opt/router?key=%v", proto, host, key)
	body, err := httpcli.Post(requrl, global.Request{Key: key, Param: param}, 5, httpcli.New(5))
	if err != nil {
		logs.Errorf("%v", err.Error())
		return
	}
	result := result.R{}
	err = json.Parse(body, &result)
	if err != nil {
		logs.Errorf("%v", json.String(result))
		return
	}
	// logs.Errorf("==>>> %v", json.String(result))
	if !result.Ok() || result.Empty() {
		logs.Errorf("%v", json.String(result))
		return
	}
	switch data := result.Data.(type) {
	case map[string]any:
		_, ok := data["token"]
		if !ok {
			logs.Errorf("%v", json.String(result))
			return
		}
		// logs.Warnf("result: => %v", json.String(result))
		v, _ := sign.Decode(data["token"].(string), []byte(config.Config.Client.AesKey))
		if validator.Empty(data) {
			logs.Errorf("%v", json.String(result))
			return
		}
		res := global.RouterResult{}
		err := json.MapToStruct(v, &res)
		if err != nil {
			logs.Errorf("%v", json.String(result))
			return
		}
		i := rand.Int() % len(res.Data)
		// minLoads := 0
		// for k := range res.Data {
		// 	switch k {
		// 	case 0:
		// 		i = 0
		// 		minLoads = res.Data[k].NumOfLoads
		// 	default:
		// 		if minLoads > res.Data[k].NumOfLoads {
		// 			i = k
		// 			minLoads = res.Data[k].NumOfLoads
		// 		}
		// 	}
		// }
		ipaddr = res.Data[i].Host
		domain = res.Data[i].Domain
		logs.Debugf("routerRes => %v", json.String(res))
	default:
		logs.Errorf("error")
	}
	return
}

func GetRouter(httpaddr string, node string, Type int) (ipaddr, domain string) {
	defer safe.Catch()
	key, param := BuildRouter(node, Type)
	ipaddr, domain = SendRouter(httpaddr, key, param)
	return
}
