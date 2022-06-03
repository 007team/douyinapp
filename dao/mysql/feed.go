package mysql

import (
	"log"

	"github.com/007team/douyinapp/models"
)

//

// FindVideo todo 无登录态
func FindVideo() ([]models.Video, error) {
	var VideoList []models.Video

	err := db.Preload("Author").Order("updated_at DESC").Limit(30).Find(&VideoList).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return VideoList, nil

}

// todo 有登录状态
func FindVideoByToken() {

}
