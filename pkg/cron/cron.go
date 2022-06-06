package cron

import (
	"github.com/robfig/cron"
)

func task() {
	//mdb := mysql.DB()
	//rdb := redis.RDB()
	// 把点赞数和评论数写入mysql

}

func Cron() {
	c := cron.New()
	c.AddFunc("*/60 * * * *", task)

	go c.Start()
	defer c.Stop()
}
