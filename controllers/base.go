package controllers

import (
	"youku/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

// ChannelRegion 获取频道地区
// @router /channel/region [*]
func (c *BaseController) ChannelRegion() {
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

	num, regions, err := models.GetChannelRegion(channelID)
	if err == nil {
		c.Data["json"] = Success(0, "success", regions, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}

	c.ServeJSON()
}

// ChannelType 获取频道类型
// @router /channel/type [*]
func (c *BaseController) ChannelType() {
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

	num, types, err := models.GetChannelType(channelID)
	if err == nil {
		c.Data["json"] = Success(0, "success", types, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}

	c.ServeJSON()
}
