package mysql

import (
	"github.com/007team/douyinapp/models"
	"log"
)

func FindVideo()([]models.Video, error){
	var VideoList []models.Video
	//var err error

	err := db.Select("video.id").Limit(30).Order("Id DESC").Joins("User").Find(&VideoList).Error
	if err!=nil{
		log.Println(err)
		return nil,err
	}


	return VideoList,nil



}
