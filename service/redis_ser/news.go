package redis_ser

import (
	"encoding/json"
	"fmt"
	"goblog_server/global"
	"time"
)

const newsIndex = "news_index"

type NewData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

// SetNews 设置某一个数据，重复执行，重复累加
//func SetNews(key string, newData []NewData) error {
//	byteData, _ := json.Marshal(newData)
//	err := global.Redis.HSet(newsIndex, key, byteData).Err()
//	return err
//}
//
//func GetNews(key string) (newData []NewData, err error) {
//	res := global.Redis.HGet(newsIndex, key).Val()
//	err = json.Unmarshal([]byte(res), &newData)
//	return
//}

// 但是这样会有一个问题
//
// 一直都是用的缓存中的数据
//
// 1. 定时器定时将键值删掉
// 2. 使用key-val的形式，每一个都有过期时间

// SetNews 设置某一个数据，重复执行，重复累加
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	err := global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 10*time.Second).Err() //设置一小时过期
	return err
}

func GetNews(key string) (newData []NewData, err error) {
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}
