package images_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models/res"
	"goblog_server/service/image_ser"
	"os"
)

var ImageService image_ser.ImageService

// ImageUploadView 上传单张/多张图片
func (i ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	filelist, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("File doesn't exist", c)
		return
	}

	// 判断路径是否存在
	// 不存在则创建
	info, err := os.Stat(global.Config.Uploads.Path)
	if os.IsNotExist(err) || !info.IsDir() { // 如果不存在或者为文件的话，则创建文件夹
		err := os.MkdirAll(global.Config.Uploads.Path, os.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
			res.FailWithMessage(err.Error(), c)
			return
		}
	} else if err != nil { //如果是其他错误
		global.Log.Error(err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}

	responseList := ImageService.ImageUploadService(filelist, c)

	res.OkWithData(responseList, c)
}
