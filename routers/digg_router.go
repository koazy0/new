package routers

import (
	"goblog_server/api"
	"goblog_server/middlewares"
)

func (r RouterGroup) DiggRouter() {
	digg_api := api.Apigroup.DiggApi
	r.POST("digg/article", middlewares.JwtAuth(), digg_api.DiggArticleView)

}
