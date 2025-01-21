package settings_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models/res"
)

// 配置信息的一些api
func (s SettingsApi) SettingsEmailView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
	//res.FailWithCode(res.SettingError, c)
}
