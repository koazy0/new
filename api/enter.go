package api

import (
	"goblog_server/api/advert_api"
	"goblog_server/api/article_api"
	"goblog_server/api/comment_api"
	"goblog_server/api/digg_api"
	"goblog_server/api/images_api"
	"goblog_server/api/menu_api"
	"goblog_server/api/message_api"
	"goblog_server/api/new_api"
	"goblog_server/api/settings_api"
	"goblog_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
	NewApi      new_api.NewApi
}

var Apigroup = new(ApiGroup)
