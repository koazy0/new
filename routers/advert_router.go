package routers

import "goblog_server/api"

func (r RouterGroup) AdvertRouter() {
	advert_api := api.Apigroup.AdvertApi
	r.POST("adverts", advert_api.AdvertCreateView)
	r.GET("adverts", advert_api.AdvertListView)
	r.DELETE("adverts", advert_api.AdvertRemoveView) // todo jwt鉴权
	r.PUT("adverts", advert_api.AdvertUpdateView)
}
