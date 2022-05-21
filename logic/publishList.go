package logic

import (
	"log"

	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/models"
)

func PublishList(userId int64) (VideoArr []models.Video) {
	VideoArr = mysql.GetVideoArr(userId)
	if VideoArr == nil {
		log.Fatalln("该用户不存在")
	}
	return
}
