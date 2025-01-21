package common

import (
	"goblog_server/global"
	"goblog_server/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo          //传入的pageinfo
	Debug           bool     //是否需要debug查看MySQL查询语句
	Likes           []string // 需要模糊匹配的字段列表
	Where           *gorm.DB // 额外的查询
	Preload         []string // 预加载的字段列表
}

// ComList 返回一个T型的列表
// 传入任意类型的model，返回一个model型的列表
// 引入泛型T是为了更好的获取类型
func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		//将新的会话实例赋值给 DB，即当前会话将使用 global.MysqlLog 记录日志
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "create_at desc" // 默认按照时间从后往前排（降序）
	}

	DB.Model(model).Count(&count) //获取数量
	// count = DB.Select("id").Find(&list).RowsAffected  数据量大的话，这里会有点慢
	offset := (option.Page - 1) * option.Limit
	// 对offset上下值进行调整
	if offset < 0 {
		offset = 0
	}
	if offset > int(count) {
		offset = int(count)
	}

	// select * from banner_models limit 2 offset 0 ;
	// limit 表示展示的数量，offset表示偏移量（上面展示id为1，2的两条）
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
