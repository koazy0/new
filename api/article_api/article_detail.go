package article_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/models/res"
	"goblog_server/service/es_ser"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

// ArticleDetailView 访问方式：http://127.0.0.1:8080/api/articles/Jo7DboYB6uoytGZAyrHz
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

// ArticleDetailByTitleView 访问方式：http://127.0.0.1:8080/api/articles/detail?title=python基础
func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 将title作为keyworld传入
	model, err := es_ser.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
