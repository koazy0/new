package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"goblog_server/global"
)

// FullTextModel 全文表

type FullTextModel struct {
	ID string `json:"id" structs:"id"` // ES的ID
	// 在ES中为text文档属性
	Key   string `json:"key"`                   //文章关联的ID
	Title string `json:"title" structs:"title"` // 文章标题
	Body  string `json:"body" structs:"body"`   // 文章内容
	Slug  string `json:"slug" structs:"slug"`   //标题的跳转地址
}

func (FullTextModel) Index() string {
	return "full_text"
}

func (FullTextModel) Mapping() string {
	return `
{
  "settings": {
    "index": {
      "max_result_window": "100000"
    }
  },
  "mappings": {
    "properties": {
      "key": {
        "type": "keyword"
      },
      "title": {
        "type": "text"
      },
      "slug": {
        "type": "keyword"
      },
	  "body": {
        "type": "text"
      },
    }
  }
}`
}

// IndexExists 索引是否存在
func (demo FullTextModel) IndexExists() bool {
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
func (demo FullTextModel) CreateIndex() error {
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
func (demo FullTextModel) RemoveIndex() error {
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
