package models

import (
	"github.com/astaxie/beego/orm"
)

type Region struct {
	Id   int
	Name string
}

type Type struct {
	Id   int
	Name string
}

func GetChannelRegion(channelId int) (num int64, regions []Region, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,name from channel_region where status=1 and channel_id=? order by sort desc", channelId).QueryRows(&regions)
	return
}

func GetChannelType(channelId int) (num int64, types []Type, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,name from channel_type where status=1 and channel_id=? order by sort desc", channelId).QueryRows(&types)
	return
}
