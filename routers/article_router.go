package routers

import (
	"goblog_server/api"
	"goblog_server/middlewares"
)

func (r RouterGroup) ArticleRouter() {
	adticle_api := api.Apigroup.ArticleApi
	r.POST("articles", middlewares.JwtAuth(), adticle_api.ArticleCreateView)
	r.GET("articles", adticle_api.ArticleListView)
	r.GET("articles/detail", adticle_api.ArticleDetailByTitleView)
	r.GET("articles/:id", adticle_api.ArticleDetailView)
}
