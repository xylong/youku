package controllers

import (
	"github.com/astaxie/beego/validation"
	"youku/models"

	"github.com/astaxie/beego"
)

// CommentController 评论
type CommentController struct {
	beego.Controller
}

// CommentInfo 评论信息
type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
}

// 评论列表
// @router /comment/list [*]
func (c *CommentController) List() {
	id, _ := c.GetInt("episodesId")
	limit, _ := c.GetInt("limit", 12)
	offset, _ := c.GetInt("offset", 0)

	valid := validation.Validation{}
	valid.Required(id, "episodesId").Message("必须指定剧集")
	valid.Min(id, 1, "episodesId").Message("剧集错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, comments, err := models.GetCommentList(id, offset, limit)
	if err == nil {
		var (
			data        []CommentInfo
			commentInfo CommentInfo
		)
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.UserInfo, _ = models.GetUserInfo(v.UserId)
			data = append(data, commentInfo)
		}
		c.Data["json"] = Success(0, "success", data, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}
	c.ServeJSON()
}

// 提交评论
// @router /comment/save [*]
func (c *CommentController) Save() {
	content := c.GetString("content")
	uid, _ := c.GetInt("uid")
	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")

	valid := validation.Validation{}
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(uid, "uid").Message("请先登录")
	valid.Min(uid, 1, "uid").Message("请先登录")
	valid.Required(episodesId, "episodesId").Message("必须指定剧集")
	valid.Min(episodesId, 1, "episodesId").Message("剧集错误")
	valid.Required(videoId, "videoId").Message("必须指定视频")
	valid.Min(videoId, 1, "videoId").Message("视频错误")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	err := models.SaveComment(content, episodesId, uid, videoId)
	if err == nil {
		c.Data["json"] = Success(0, "success", "", 1)
	} else {
		c.Data["json"] = Fail(5000, err)
	}
	c.ServeJSON()
}
