package models

import "goblog_server/models/ctype"

// MenuModel 菜单表(主页表)
type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                     // slogan
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                               // 简介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` // 菜单的图片列表
	// many2many:menu_banner_models 指定这是一个多对多关系；使用 GORM 的 Preload 功能来加载关联的 Banners
	// joinForeignKey:MenuID MenuID是中间表连接菜单表的外键
	// JoinReferences:BannerID BannerID是中间表连接图片表的外键
	//
	BannerTime int `json:"banner_time"`         // 菜单图片的切换时间 为 0 表示不切换
	Sort       int `gorm:"size:10" json:"sort"` // 菜单的顺序
}
