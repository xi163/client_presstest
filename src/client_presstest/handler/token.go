package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils/codec/base64"
	"github.com/xi123/libgo/utils/codec/uri"
	"github.com/xi123/libgo/utils/crypto/aes"
	"github.com/xi123/libgo/utils/crypto/md5"
	"github.com/xi123/libgo/utils/httpcli"
	"github.com/xi123/libgo/utils/json"
	"github.com/xi123/libgo/utils/result"
	"github.com/xi123/libgo/utils/safe"
	"github.com/xi123/libgo/utils/sign"
	"github.com/xi123/libgo/utils/validator"
	"github.com/cwloo/presstest/src/client_presstest/global"
)

func BuildPlatLogin(account string, Type int) (key, param string) {
	req := global.LoginParam{
		Account:   account,
		Type:      Type,
		Timestamp: time.Now().Unix(),
	}
	key = md5.Md5(fmt.Sprintf("%v%v%v%v", req.Account, req.Type, req.Timestamp, global.Md5Key), true)
	// rawParam := fmt.Sprintf("account=%v&type=%v&timestamp=%v", req.Account, req.Type, req.Timestamp)
	rawParam := json.String(req)
	encrypt := aes.ECBEncryptPKCS7([]byte(rawParam), []byte(global.AesKey), []byte(global.AesKey))
	strBase64 := base64.URLEncode(encrypt)
	param = uri.URLEncode(strBase64)
	return
}

func SendPlatLogin(httpaddr string, key, param string) (token, ipaddr, domain string) {
	defer safe.Catch()
	vec := strings.Split(httpaddr, "//")
	if len(vec) != 2 {
		return
	}
	proto := strings.Trim(vec[0], ":")
	host := vec[1]
	requrl := fmt.Sprintf("%v://%v/login?key=%v", proto, host, key)
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
		logs.Warnf("result: => %v", json.String(result))
		v, _ := sign.Decode(data["token"].(string), []byte(global.AesKey))
		if validator.Empty(data) {
			logs.Errorf("%v", json.String(result))
			return
		}
		res := global.LoginResult{}
		err := json.MapToStruct(v, &res)
		if err != nil {
			logs.Errorf("%v", json.String(result))
			return
		}
		token = data["token"].(string)
		ipaddr = res.Data[0].Host
		domain = res.Data[0].Domain
		logs.Debugf("loginRes => %v", json.String(res))
	}
	return
}

func GetPlatToken(httpaddr string, account string, Type int) (token, ipaddr, domain string) {
	defer safe.Catch()
	key, param := BuildPlatLogin(account, Type)
	token, ipaddr, domain = SendPlatLogin(httpaddr, key, param)
	return
}
