package routers

import (
	"goblog_server/api"
	"goblog_server/middlewares"
)

func (router RouterGroup) MessageRouter() {
	app := api.Apigroup.MessageApi
	router.POST("messages", middlewares.JwtAuth(), app.MessageCreateView)
	router.GET("messages_all", app.MessageListAllView)
	router.GET("messages", middlewares.JwtAuth(), app.MessageListView)
	router.POST("messages_record", middlewares.JwtAuth(), app.MessageRecordView)
	router.GET("message_users", middlewares.JwtAdmin(), app.MessageUserListView)
	router.DELETE("message_users", middlewares.JwtAuth(), app.MessageRecordRemoveView) // 删除聊天记录
	router.GET("message_users/user", middlewares.JwtAdmin(), app.MessageUserListByUserView)
	router.GET("message_users/record", middlewares.JwtAdmin(), app.MessageUserRecordView)
	router.GET("message_users/me", middlewares.JwtAuth(), app.MessageUserListByMeView)
	router.GET("message_users/record/me", middlewares.JwtAuth(), app.MessageUserRecordByMeView)
}
