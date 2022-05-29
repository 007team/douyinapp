package main

import (
	"fmt"

	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/redis"
	"github.com/007team/douyinapp/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置信息初始化
	if err := settings.Init(); err != nil {
		fmt.Println("settings init failed", err)
		return
	}

	// mysql 初始化
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("mysql init failed", err)
		return
	}

	//redis 初始化
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("redis init failed", err)
		return
	}
	defer redis.Close()

	r := gin.Default()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
