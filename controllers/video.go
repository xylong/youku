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
