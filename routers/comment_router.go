package routers

import (
	"goblog_server/api"
	"goblog_server/middlewares"
)

func (r RouterGroup) CommentRouter() {
	comment_api := api.Apigroup.CommentApi
	r.POST("comments", middlewares.JwtAuth(), comment_api.CommentCreateView)
	r.GET("comments", comment_api.CommentListView)
	r.GET("comments/:id", comment_api.CommentDigg)
	r.DELETE("comments/:id", comment_api.CommentRemoveView)
}
