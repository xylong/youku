package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
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