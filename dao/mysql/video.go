package mysql

import (
	"fmt"
	"github.com/007team/douyinapp/models"
	"log"
)

//


// FindVideo todo 无登录态
func FindVideo()([]models.Video, error){
	var VideoList []models.Video
	//

	err := db.Preload("Author").Order("updated_at DESC").Limit(30).Find(&VideoList).Error

	if err!=nil{
		log.Println(err)
		return nil,err
	}

	return VideoList,nil

}


// GetLastId 获取最后一位视频id
func GetLastId(video *models.Video)(id int64){
	db.Last(&video)
	fmt.Println(video.Id)
	return video.Id
}
