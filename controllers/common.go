package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"time"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func Success(code int, msg, items interface{}, count int64) *JsonStruct {
	return &JsonStruct{
		Code:  code,
		Msg:   msg,
		Items: items,
		Count: count,
	}
}

func Fail(code int, msg interface{}) *JsonStruct {
	return &JsonStruct{
		Code: code,
		Msg:  msg,
	}
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

// DateFormat 时间戳->年-月-日
func DateFormat(tiemStamp int64) string {
	return time.Unix(tiemStamp, 0).Format("2006-01-02")
}
