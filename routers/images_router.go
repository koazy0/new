package routers

import (
	"goblog_server/api"
)

func (r RouterGroup) ImagesRouter() {
	images_api := api.Apigroup.ImagesApi
	r.GET("images", images_api.ImageListView)
	r.POST("images", images_api.ImageUploadView)
	r.DELETE("images", images_api.ImageRemoveView) // todo jwt鉴权
	r.PUT("images", images_api.ImageUpdateView)
}
