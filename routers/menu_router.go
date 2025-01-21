package routers

import "goblog_server/api"

func (r RouterGroup) MenusRouter() {
	menus_api := api.Apigroup.MenuApi
	r.GET("menus", menus_api.MenuListView)
	r.POST("menus", menus_api.MenuUpdateView)
	r.DELETE("menus", menus_api.MenuRemoveView) // todo jwt鉴权
	r.PUT("menus", menus_api.MenuUpdateView)
}
