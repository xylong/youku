package controllers

import (
	"youku/models"

	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// ChannelAdvert 广告
// @router /channel/advert [*]
func (c *VideoController) ChannelAdvert() {
	channelID, _ := c.GetInt("channelId")

	valid := validation.Validation{}
	valid.Required(channelID, "channel_id").Message("频道不能为空")
	valid.Min(channelID, 1, "channel_id").Message("频道错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelAdvert(channelID)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "数据请求失败，稍后重试~")
	}

	c.ServeJSON()
}

// ChannelHotList 频道页-获取正在热播
// @router /channel/hot [get]
func (c *VideoController) ChannelHotList() {
	channelID, _ := c.GetInt("channelId")

	valid := validation.Validation{}
	valid.Required(channelID, "channel_id").Message("必须指定频道")
	valid.Min(channelID, 1, "channel_id").Message("频道错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelHotList(channelID)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容~")
	}

	c.ServeJSON()
}

// ChannelRecommendList 频道页-根据地区获取推荐的视频
// @router /channel/recommend/region [get]
func (c *VideoController) ChannelRecommendList() {
	channelID, _ := c.GetInt("channelId")
	regionID, _ := c.GetInt("regionId")

	valid := validation.Validation{}
	valid.Required(channelID, "channel_id").Message("必须指定频道")
	valid.Min(channelID, 1, "channel_id").Message("频道错误")
	valid.Required(regionID, "region_id").Message("必须指定频道地区")
	valid.Min(regionID, 1, "region_id").Message("地区错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelID, regionID)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "数据请求失败，稍后重试~")
	}

	c.ServeJSON()
}

// GetChannelRecommendTypeList 频道页-根据类型获取视频
// @router /channel/recommend/type [get]
func (c *VideoController) GetChannelRecommendTypeList() {
	channelID, _ := c.GetInt("channelId")
	typeID, _ := c.GetInt("typeId")

	valid := validation.Validation{}
	valid.Required(channelID, "channel_id").Message("必须指定频道")
	valid.Min(channelID, 1, "channel_id").Message("频道错误")
	valid.Required(typeID, "type_id").Message("必须指定频道类型")
	valid.Min(typeID, 1, "type_id").Message("类型错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelID, typeID)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}

	c.ServeJSON()
}

// ChannelVideo 获取视频
// @router /channel/video [*]
func (c *VideoController) ChannelVideo() {
	p := models.VideoParam{}
	p.ChannelID, _ = c.GetInt("channelId")
	p.RegionID, _ = c.GetInt("regionId")
	p.TypeID, _ = c.GetInt("typeId")
	p.End = c.GetString("end")
	p.Sort = c.GetString("sort")
	p.Limit, _ = c.GetInt("limit")
	p.Offset, _ = c.GetInt("offset")

	valid := validation.Validation{}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelVideoList(&p)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}

	c.ServeJSON()
}

// VideoInfo 获取视频详情
// @router /video/info [*]
func (c *VideoController) VideoInfo() {
	id, _ := c.GetInt("videoId")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("必须指定视频")
	valid.Min(id, 1, "id").Message("视频不存在")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	video, err := models.GetVideoInfo(id)
	if err == nil {
		c.Data["json"] = Success(0, "success", video, 1)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}

	c.ServeJSON()
}

// VideoEpisodesList 视频剧集列表
// @router /video/episodes/list [*]
func (c *VideoController) VideoEpisodesList() {
	id, _ := c.GetInt("videoId")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("必须指定视频")
	valid.Min(id, 1, "id").Message("视频不存在")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, episodes, err := models.GetVideoEpisodesList(id)
	if err == nil {
		c.Data["json"] = Success(0, "success", episodes, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}

	c.ServeJSON()
}

// UserVideo 我的视频管理
// @router /user/video [*]
func (c *VideoController) UserVideo() {
	uid, _ := c.GetInt("uid")
	if uid == 0 {
		c.Data["json"] = Fail(4001, "必须指定用户")
		c.ServeJSON()
	}
	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}
	c.ServeJSON()
}

// 保存用户上传视频信息
// @router /video/save [*]
func (c VideoController) VideoSave() {
	playUrl := c.GetString("playUrl")
	title := c.GetString("title")
	subTitle := c.GetString("subTitle")
	channelId, _ := c.GetInt("channelId")
	typeId, _ := c.GetInt("typeId")
	regionId, _ := c.GetInt("regionId")
	uid, _ := c.GetInt("uid")
	if uid == 0 {
		c.Data["json"] = Fail(4001, "请先登录")
		c.ServeJSON()
	}
	if playUrl == "" {
		c.Data["json"] = Fail(4002, "视频地址不能为空")
		c.ServeJSON()
	}
	err := models.SaveVideo(title, subTitle, playUrl, "", channelId, regionId, typeId, uid)
	if err != nil {
		c.Data["json"] = Success(0, "success", nil, 1)
	} else {
		c.Data["json"] = Fail(5000, err)
	}
	c.ServeJSON()
}
