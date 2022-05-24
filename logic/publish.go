package logic

import (
	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/models"
)

func PublishList(userId int64) (VideoArr []models.Video) {
	VideoArr = mysql.GetVideoArr(userId)
	return
}
