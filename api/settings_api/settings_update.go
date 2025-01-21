package settings_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/config"
	"goblog_server/core"
	"goblog_server/global"
	"goblog_server/models/res"
)

//// 配置信息的一些api
//func (s SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
//
//	var cr config.SiteInfo
//
//	err := c.ShouldBindJSON(&cr)
//	if err != nil {
//		res.FailWithCode(res.ArgumentError, c)
//		return
//	}
//
//	global.Config.SiteInfo = cr
//	err = core.WriteConf()
//	if err != nil {
//		global.Log.Error(err)
//		res.FailWithMessage(err.Error(), c)
//		return
//	}
//	res.OkWithMessage("success", c)
//}

// SettingsInfoUpdateView 修改某一项的配置信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info

	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwts":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	err = core.WriteConf()
	if err != nil {
		global.Log.Errorln(err)
		return
	}
	res.OkWithMessage("Success", c)
}
