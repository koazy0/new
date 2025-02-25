package routers

import (
	"goblog_server/api"
	"goblog_server/middlewares"
)

func (r RouterGroup) ArticleRouter() {
	adticle_api := api.Apigroup.ArticleApi
	r.POST("articles", middlewares.JwtAuth(), adticle_api.ArticleCreateView)
	r.GET("articles", adticle_api.ArticleListView)
	r.PUT("articles", adticle_api.ArticleUpdateView)
	r.DELETE("articles", adticle_api.ArticleRemoveView)
	r.GET("articles/detail", adticle_api.ArticleDetailByTitleView)
	r.GET("articles/calendar", adticle_api.ArticleCalendarView)
	r.GET("articles/tags", adticle_api.ArticleTagListView)
	r.GET("articles/:id", adticle_api.ArticleDetailView)
	r.POST("articles/collects", middlewares.JwtAuth(), adticle_api.ArticleCollCreateView)
	r.GET("articles/collects", middlewares.JwtAuth(), adticle_api.ArticleCollListView)
	r.DELETE("articles/collects", middlewares.JwtAuth(), adticle_api.ArticleCollBatchRemoveView)

}
