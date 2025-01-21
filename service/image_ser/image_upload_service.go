package image_ser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/plugnis/qiniu"
	"goblog_server/utils"
	"io"
	"mime/multipart"
	"path"
)

// 对每张图片上传的响应
type FileUploadResponse struct {
	FileName  string `json:"fileName"`   // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Message   string `json:"message"`    // 返回消息
}

// ImageUploadService 文件上传的方法
func (s ImageService) ImageUploadService(filelist []*multipart.FileHeader, c *gin.Context) (responseList []FileUploadResponse) {

	for _, fileHeader := range filelist {

		tmpResponse := FileUploadResponse{}
		//fmt.Println(fileHeader.Size)     // 图片大小 单位kb
		//fmt.Println(fileHeader.Header)   // 请求头
		//fmt.Println(fileHeader.Filename) // 上传的文件名

		// 组合路径+文件名
		filepath := path.Join(global.Config.Uploads.Path, fileHeader.Filename)

		// 白名单校验
		if !utils.CheckWhiteImageList(fileHeader.Filename, global.WhiteImageList) {
			tmpResponse.IsSuccess = false
			tmpResponse.Message = "Not Images"
			responseList = append(responseList, tmpResponse)
			continue
		}

		// 判断大小是否符合要求
		size := float64(fileHeader.Size) / float64(1048576) // 转化为MB

		if size > float64(global.Config.Uploads.Size) {
			tmpResponse.IsSuccess = false
			tmpResponse.Message = fmt.Sprintf("File size: %.2fMB. Size limit: %dMB", size, global.Config.Uploads.Size) //保留两位小数
		} else {

			// 大小符合要求，可以上传到服务器了

			fileobj, err := fileHeader.Open()
			if err != nil {
				global.Log.Error(err.Error())
				tmpResponse.IsSuccess = false
				tmpResponse.Message = err.Error()
				continue
			}
			uploadByte, _ := io.ReadAll(fileobj)
			uploadMd5 := utils.MD5(uploadByte) // 用utils中的md5方法
			//fmt.Println(uploadMd5)             // 先计算出上传文件的md5

			// 在数据库中查找这个文件是否被上传过了
			var bannerModel models.BannerModel
			err = global.DB.Take(
				&bannerModel,
				"hash = ?",      //单纯查找是否存在用take方法，
				uploadMd5).Error // 若要要求顺序（主键排序）的话用fitst方法

			if err == nil {
				// 找到了
				tmpResponse.Message = "图片已存在"
				tmpResponse.FileName = bannerModel.Path
				tmpResponse.IsSuccess = false
				responseList = append(responseList, tmpResponse)
				continue
			}

			// 若打开了要上传到七牛云，则之上传到七牛云 不保存在本地
			success := true
			err = nil

			if global.Config.QiNiu.Enable {
				filepath, err = qiniu.UploadImage(uploadByte, fileHeader.Filename, "goblog")
				if err != nil {
					success = false
				}
			} else {
				err = c.SaveUploadedFile(fileHeader, filepath)
				if err != nil {
					success = false
				}
			}

			if success {
				tmpResponse.IsSuccess = true
				tmpResponse.Message = "Upload Success"
				tmpResponse.FileName = filepath
			} else {
				tmpResponse.IsSuccess = false
				tmpResponse.Message = err.Error()
				global.Log.Error(err.Error())
			}

			// 方法：
			global.DB.Create(&models.BannerModel{
				Path: filepath,
				Hash: uploadMd5,
				Name: fileHeader.Filename,
			})
		}
		responseList = append(responseList, tmpResponse)
	}
	return
}
