package mysql

import (
	"log"

	"github.com/007team/douyinapp/models"
)

//

// FindVideo todo 无登录态
func FindVideo() (videos []models.Video, err error) {

	//

	err = db.Preload("Author").Order("updated_at DESC").Limit(30).Find(&videos).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return videos, nil

}

// todo 有登录状态
func FindVideoByToken() {

}
