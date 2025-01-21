package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/res"
	"goblog_server/service/es_ser"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//list, count, err := es_ser.CommList(cr.Key, cr.Page, cr.Limit)
	list, count, err := es_ser.CommList(es_ser.Option{})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	// 以list为过滤字过滤掉content字段，对应代码为`json:"content,omit(list)"`
	retlist := filter.Omit("list", list)
	res.OKWithList(retlist, int64(count), c)
}
