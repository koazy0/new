package routers

import (
	"goblog_server/api"
)

func (r RouterGroup) NewRouter() {
	new_api := api.Apigroup.NewApi
	r.POST("news", new_api.NewListView)

}
