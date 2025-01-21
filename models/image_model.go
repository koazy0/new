package models

import (
	"goblog_server/global"
	"goblog_server/models/ctype"
	"gorm.io/gorm"
	"os"
	"path"
)

// BannerModel banner表
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"defalut:1" json:"image_type"` // 是存储在本地还是oss,默认存储在本地
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(path.Join(b.Path, b.Name)) // 拼接路径和名字
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	//todo 删除七牛云的图片
	return nil
}
