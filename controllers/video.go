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
