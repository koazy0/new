package images_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/models"
	"goblog_server/models/res"
	"goblog_server/service/common"
)

// ImageListView 主要用于后台管理返回图片列表
func (i ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel

	//注意调用方法
	imageList, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OKWithList(imageList, count, c)

}
