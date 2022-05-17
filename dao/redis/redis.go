// Package redis
// 操作redis数据库
//
package redis

import (
	"fmt"
	"strconv"

	"github.com/007team/douyinapp/settings"
	"github.com/go-redis/redis"
)

// 这是操作redis的全局变量
// 使用方法：
var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, strconv.Itoa(cfg.Port)),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis ping failed", err)
		return err
	}
	return nil
}
func Close() {
	err := rdb.Close()
	if err != nil {
		fmt.Println("redis close failed", err)
		return
	}
}
