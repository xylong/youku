package models

import "github.com/astaxie/beego/orm"

type Barrage struct {
	Id int
	Content string
	CurrentTime int
	AddTime int64
	UserId int
	Status int
	EpisodesId int
	VideoId int
}

type BarrageData struct {
	Id int `json:"id"`
	Content string `json:"content"`
	CurrentTime int `json:"currentTime"`
}

func init() {
	orm.RegisterModel(new(Barrage))
}

func BarrageList(episodesId,start,end int) (num int64,barrages []BarrageData,err error) {
	o:=orm.NewOrm()
	num,err=o.Raw("select id,content,current_time from barrage where status=1 and episodes_id=? and current_time=? and current_time<? order by current_time asc",
		episodesId,start,end).QueryRows(&barrages)
	return
}