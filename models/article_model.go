package models

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"goblog_server/global"
	"goblog_server/models/ctype"
)

// ArticleModel 文章表

type ArticleModel struct {
	ID        string `json:"id"`         // ES的ID
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"-"`          // 更新时间

	// 在ES中为text文档属性
	Title    string `json:"title"`              // 文章标题
	Keyword  string `json:"keyword,omit(list)"` //关键字
	Abstract string `json:"abstract"`           // 文章简介
	Content  string `json:"content,omit(list)"` // 文章内容
	// 指定了忽略特定字段
	// 当指定的过滤条件为list时，序列化时会忽略Content` 字段。
	// 这种功能通常用于动态返回不同的数据视图，以适应不同的业务场景。

	// 在ES中为int属性
	LookCount     int `json:"look_count"`     // 浏览量
	CommentCount  int `json:"comment_count"`  // 评论量
	DiggCount     int `json:"digg_count"`     // 点赞量
	CollectsCount int `json:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id"`        // 用户id
	UserNickName string `json:"user_nick_name"` // 用户昵称
	UserAvatar   string `json:"user_avatar"`    //用户头像

	Category string `json:"category"`          // 文章分类
	Source   string `json:"source,omit(list)"` // 文章来源
	Link     string `json:"link,omit(list)"`   // 原文链接

	BannerID  uint   `json:"banner_id"`  // 文章封面
	BannerURL string `json:"banner_url"` // 文章封面id

	Tags ctype.Array `json:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index": {
      "max_result_window": "100000"
    }
  },
  "mappings": {
    "properties": {
      "title": {
        "type": "text"
      },
      "keyword": {
        "type": "keyword"
      },
	  "abstract": {
        "type": "text"
      },
      "content": {
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": {
        "type": "text"
      },
      "user_avatar": {
        "type": "text"
      },
      "category": {
        "type": "text"
      },
      "source": {
        "type": "text"
      },
      "link": {
        "type": "text"
      },
      "banner_id": {
        "type": "integer"
      },
      "created_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}`
}

// IndexExists 索引是否存在
func (demo ArticleModel) IndexExists() bool {
	//Do() 方法负责实际发送请求到 Elasticsearch 集群，并返回请求的结果
	//IndexExists 本身只定义了请求的内容，但并未发送请求到 Elasticsearch。
	//Do() 是真正将请求执行并获取结果的关键步骤。
	exists, err := global.ESClient.
		IndexExists(demo.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
	}
	return exists
}

// CreateIndex 创建索引
func (demo ArticleModel) CreateIndex() error {
	if demo.IndexExists() {
		// 有索引，则删掉索引重新创建
		demo.RemoveIndex()
	}
	// 没有索引，则创建索引
	createIndex, err := global.ESClient.
		CreateIndex(demo.Index()).
		BodyString(demo.Mapping()).
		Do(context.Background())

	// 这个处理的是发送的错误
	if err != nil {
		global.Log.Errorf("发送错误:%s\n", err.Error())
		return err
	}
	//这个处理的是ES已接收但未执行的错误
	//当 Elasticsearch 确认该操作已成功执行，Acknowledged 才会是 true
	if !createIndex.Acknowledged {
		global.Log.Errorf("创建失败:%s\n", err.Error())
		return err
	}
	global.Log.Infof("索引 %s 创建成功\n", demo.Index())
	return nil
}

// RemoveIndex 删除索引
func (demo ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(demo.Index()).Do(context.Background())
	if err != nil {
		global.Log.Errorf("删除索引失败:%s\n", err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		global.Log.Errorf("删除索引失败\n")
		return err
	}
	global.Log.Info("索引删除成功")
	return nil
}

// Create 添加的方法
func (demo ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(demo.Index()).
		BodyJson(demo).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	demo.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (demo ArticleModel) ISExistData() bool {
	res, err := global.ESClient.
		Search(demo.Index()).
		Query(elastic.NewTermQuery("keyword", demo.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}
