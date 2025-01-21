package settings_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/config"
	"goblog_server/core"
	"goblog_server/global"
	"goblog_server/models/res"
)

func (s SettingsApi) SettingsEmailUpdateView(c *gin.Context) {

	var cr config.Email

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.Email = cr
	err = core.WriteConf()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("success", c)
}
