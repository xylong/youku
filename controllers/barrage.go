package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader=websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId int
}

// 获取弹幕websocket
// @router /barrage/ws [*]
func (c *BarrageController) BarrageWs() {
	var (
		conn *websocket.Conn
		err error
		data []byte
	)

	if conn,err=upgrader.Upgrade(c.Ctx.ResponseWriter,c.Ctx.Request,nil);err!=nil {
		goto ERR
	}
	for  {
		if _,data,err=conn.ReadMessage();err!=nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal([]byte(data),&wsData)
		endTime:=wsData.CurrentTime+int(time.Second*60)
	}
	ERR:
		conn.Close()
}
