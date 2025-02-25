package models

import "time"

type MODEL struct {
	ID       uint      `gorm:"primarykey" json:"id"`            // 主键ID
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"` // 创建时间
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"-"`         // 更新时间
}

type PageInfo struct {
	Page  int    `form:"page"` // 第几页
	Key   string `form:"key"`
	Limit int    `form:"limit"` //展示数量
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id"`
}
