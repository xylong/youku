package models

import (
	"github.com/astaxie/beego/orm"
)

type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	AddTime       int64
}

func GetChannelHotList(channelId int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,episodes_count,is_end,add_time from video where status=1 and is_hot=1 and channel_id=? order by episodes_update_time desc limit 9", channelId).QueryRows(&videos)
	return
}
