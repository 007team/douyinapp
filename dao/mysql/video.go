package mysql

import (
	"github.com/007team/douyinapp/models"
	"log"
)

func FindVideo()([]models.Video, error){
	var VideoList []models.Video

	//err := db.Select("videos.id, play_url, cover_url, favorite_count, comment_count, is_favorite, title, user_id, users.id,users.name,users.follow_count,users.follower_count，users.is_follow").Limit(30).Order("Id DESC").Joins("User").Find(&VideoList).Error
	err := db.Preload("Author").Find(&VideoList).Error
	if err!=nil{
		log.Println(err)
		return nil,err
	}


	return VideoList,nil



}
