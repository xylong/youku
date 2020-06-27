package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Comment 评论
type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init() {
	orm.RegisterModel(new(Comment))
}

// GetCommentList 评论列表
func GetCommentList(episodeId, offset, limit int) (num int64, comments []Comment, err error) {
	o := orm.NewOrm()
	// ? 统计
	num, _ = o.Raw("select id from comment where status=1 and episodes_id=?", episodeId).QueryRows(&comments)
	_, err = o.Raw("select id,content,add_time,user_id,stamp,praise_count,episodes_id from comment where status=1 and episodes_id=? order by add_time desc limit ?,?",
		episodeId, offset, limit).QueryRows(&comments)
	return
}

// SaveComment 保存评论
func SaveComment(content string, episodesId, uid, videoId int) error {
	o := orm.NewOrm()
	comment := &Comment{
		Content:    content,
		AddTime:    time.Now().Unix(),
		UserId:     uid,
		Stamp:      0,
		Status:     1,
		EpisodesId: episodesId,
		VideoId:    videoId,
	}
	_, err := o.Insert(comment)
	if err == nil {
		// 修改视频总评论数
		_, _ = o.Raw("update video set comment=comment+1 where id=?", videoId).Exec()
		// 修改视频剧集评论数
		_, _ = o.Raw("update video_episodes set comment=comment+1 where id=?", episodesId).Exec()
	}
	return err
}
