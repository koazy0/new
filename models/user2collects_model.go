package models

import "time"

// User2CollectModel 自定义第三张表  记录用户什么时候收藏了什么文章
type User2CollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID uint      `gorm:"primaryKey"`
	CreatedAt time.Time
}
