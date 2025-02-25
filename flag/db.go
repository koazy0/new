package flag

import (
	"goblog_server/global"
	"goblog_server/models"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.User2CollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuImageModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			//&models.ArticleModel{}, //用ES存
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuImageModel{},
			&models.FeedBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[success] 生成数据库表结构成功！")
}
