package models

import (
	"github.com/astaxie/beego/orm"
)

type Advert struct {
	Id        int
	Title     string
	SubTitile string
	Img       string
	Url       string
	AddTime   int64
}

func init() {
	orm.RegisterModel(new(Advert))
}

func GetChannelAdvert(channelId int) (num int64, adverts []Advert, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id, title, sub_title,img,add_time,url from advert where status=1 and channel_id=? order by sort desc limit 1", channelId).QueryRows(&adverts)
	return
}
