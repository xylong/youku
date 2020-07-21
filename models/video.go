package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Video 视频
type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int64
	Comment            int
	UserId             int
	IsRecommend        int
}

func init() {
	orm.RegisterModel(new(Video))
}

type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	Comment       int
}

// Episodes 情节
type Episode struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	AliyunVideoId string
}

func GetChannelHotList(channelId int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,episodes_count,is_end,add_time from video where status=1 and is_hot=1 and channel_id=? order by episodes_update_time desc limit 9", channelId).QueryRows(&videos)
	return
}

func GetChannelRecommendRegionList(channelId, regoinId int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,episodes_count,is_end,add_time from video where status=1 and is_recommend=1 and channel_id=? and region_id=? order by episodes_update_time desc limit 9", channelId, regoinId).QueryRows(&videos)
	return
}

func GetChannelRecommendTypeList(channelId, rtypeId int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,episodes_count,is_end,add_time from video where status=1 and is_recommend=1 and channel_id=? and type_id=? order by episodes_update_time desc limit 9", channelId, rtypeId).QueryRows(&videos)
	return
}

type VideoParam struct {
	// todo 配置详尽规则
	ChannelID int `valid:"Required;Min(1);"`
	RegionID  int `valid:Min(1);`
	TypeID    int `valid:Min(1)`
	End       string
	Sort      string
	Limit     int `valid:Min(1)`
	Offset    int `valid:Min(1)`
}

// GetChannelVideoList 频道视频列表
func GetChannelVideoList(p *VideoParam) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params

	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", p.ChannelID)
	qs = qs.Filter("status", 1)
	if p.RegionID > 0 {
		qs = qs.Filter("region_id", p.RegionID)
	}
	if p.TypeID > 0 {
		qs = qs.Filter("type_id", p.TypeID)
	}
	if p.End == "n" {
		qs = qs.Filter("is_end", 0)
	} else if p.End == "y" {
		qs = qs.Filter("is_end", 1)
	}
	if p.Sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time")
	} else if p.Sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if p.Sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}
	num, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(p.Limit, p.Offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	return num, videos, err

}

// GetVideoInfo 获取视频信息
func GetVideoInfo(id int) (video Video, err error) {
	o := orm.NewOrm()
	err = o.Raw("select * from video where id=? limit 1", id).QueryRow(&video)
	return
}

// GetVideoEpisodesList 剧集列表
func GetVideoEpisodesList(id int) (num int64, episodes []Episode, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,num,play_url,comment,add_time from video_episodes where video_id=? and status=1 order by num asc", id).QueryRows(&episodes)
	return
}

// GetChannelTop 排行榜
func GetChannelTop(channelId int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,add_time,episodes_count,is_end from video where status=1 and channel_id=? order by comment desc limit 10", channelId).QueryRows(&videos)
	return
}

// GetTypeTop 根据类型获取排行
func GetTypeTop(typeID int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,add_time,episodes_count,is_end from video where status=1 and type_id=? order by comment desc limit 10", typeID).QueryRows(&videos)
	return
}

// GetUserVideo 获取用户视频
func GetUserVideo(uid int) (num int64, videos []VideoData, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select id,title,sub_title,img,img1,add_time,episodes_count,is_end from video where user_id=? order by add_time desc", uid).QueryRows(&videos)
	return
}

func SaveVideo(title, subTitile, playUrl, aliyunVideoId string, channelId, regionId, typeId, userId int) error {
	o := orm.NewOrm()
	t := time.Now().Unix()
	video := Video{
		Title:              title,
		SubTitle:           subTitile,
		Img:                "",
		Img1:               "",
		EpisodesCount:      1,
		IsEnd:              1,
		AddTime:            0,
		ChannelId:          channelId,
		Status:             1,
		RegionId:           regionId,
		TypeId:             typeId,
		EpisodesUpdateTime: t,
		Comment:            0,
		UserId:             userId,
	}
	id, err := o.Insert(&video)
	if err == nil {
		o.Raw("INSERT INTO video_episodes (title,add_time,num,video_id,play_url,status,comment,aliyun_video_id) VALUES (?,?,?,?,?,?,?,?)", subTitile, t, 1, id, playUrl, 1, 0, aliyunVideoId).Exec()
	}
	return err
}
