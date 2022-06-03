package mysql

import (
	"github.com/007team/douyinapp/models"
	"log"
)

func FavoriteList(es []string) (videos []models.Video, err error) {
	err = db.Find(&videos, es).Error
	if err != nil {
		log.Println("db.Find failed", err)
	}
	return videos, err
}
