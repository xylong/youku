package controllers

import (
	"youku/models"

	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// @router /channel/advert [*]
func (c *VideoController) ChannelAdvert() {
	channelId, _ := c.GetInt("channelId")

	valid := validation.Validation{}
	valid.Required(channelId, "channel_id").Message("频道不能为空")
	valid.Min(channelId, 1, "channel_id").Message("频道错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelAdvert(channelId)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "数据请求失败，稍后重试~")
	}

	c.ServeJSON()
}

// 频道页-获取正在热播
// @router /channel/hot [get]
func (c *VideoController) ChannelHotList() {
	channelId, _ := c.GetInt("channelId")

	valid := validation.Validation{}
	valid.Required(channelId, "channel_id").Message("必须指定频道")
	valid.Min(channelId, 1, "channel_id").Message("频道错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelHotList(channelId)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容~")
	}

	c.ServeJSON()
}

// 频道页-根据地区获取推荐的视频
// @router /channel/recommend/region [get]
func (c *VideoController) ChannelRecommendList() {
	channelId, _ := c.GetInt("channelId")
	regionId, _ := c.GetInt("regionId")

	valid := validation.Validation{}
	valid.Required(channelId, "channel_id").Message("必须指定频道")
	valid.Min(channelId, 1, "channel_id").Message("频道错误")
	valid.Required(regionId, "region_id").Message("必须指定频道地区")
	valid.Min(regionId, 1, "region_id").Message("地区错误")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "数据请求失败，稍后重试~")
	}

	c.ServeJSON()
}

// 频道页-根据类型获取视频
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
