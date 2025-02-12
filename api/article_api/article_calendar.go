package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/res"
	"time"
)

// ES的时间聚合
// 不需要传参数，直接进行查询

/*
Elasticsearch（ES）的时间聚合（Date Histogram Aggregation）是一种聚合方式，用于按照时间区间（如每天、每周、每月）对文档进行分组，
并计算每个时间段内的文档数量（或其他统计信息）。这在时间序列数据分析中非常有用，例如：
	统计网站每日文章发布量
	计算每天的订单数量
	统计某个时间范围内的访问量变化趋势
*/

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	// 时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")

	// 时间段搜索
	// 从今天开始，到去年的今天
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)

	format := "2006-01-02 15:04:05"
	// lt 小于  gt 大于
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format(format)).
		Lte(now.Format(format))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())

	//global.ESClient.Search(models.ArticleModel{}.Index())：查询文章索引。
	//.Query(query)：添加时间范围查询（仅查询过去一年的数据）。
	//.Aggregation("calendar", agg)：添加时间聚合，命名为 "calendar"。
	//.Size(0)：不需要返回文章数据，只统计数量。
	//.Do(context.Background())：执行查询。
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)

	var resList = make([]CalendarResponse, 0)
	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
	}
	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")

		count, _ := DateCount[day]
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OkWithData(resList, c)

}
