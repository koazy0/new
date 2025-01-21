package routers

//将每个api分文件拆开来，方便维护

import (
	"goblog_server/api"
)

func (r RouterGroup) SettingsRouter() {
	settingapi := api.Apigroup.SettingsApi
	//r.GET("settings", settingapi.SettingsInfoView)
	//r.PUT("settings", settingapi.SettingsInfoUpdateView)
	//r.GET("settings_email", settingapi.SettingsEmailView)
	//r.PUT("settings_email", settingapi.SettingsEmailUpdateView)
	r.GET("settings/:name", settingapi.SettingsInfoView)
	r.PUT("settings/:name", settingapi.SettingsInfoUpdateView)
}
