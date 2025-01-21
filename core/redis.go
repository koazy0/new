package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"goblog_server/global"
	"time"
)

func InitRedis() *redis.Client {
	// 默认连0号数据库
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr(),
		Password: global.Config.Redis.Password, // no password set
		DB:       db,                           // use default DB
		PoolSize: global.Config.Redis.PoolSize, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return rdb
}
