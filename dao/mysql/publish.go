package mysql

import (
	"fmt"
	"log"

	"github.com/007team/douyinapp/models"
)

/**
作用: 获取该用户所拥有的所有视频
传入: 用户名
返回: Video数组
*/
func GetVideoArr(userId int64) (videoArr []models.Video, err error) {
	err = db.Preload("Author").Find(&videoArr, "user_id = ?", userId).Error
	if err != nil {
		log.Fatalln("mysql.GetVideoArr failed", err)
	}

	return
}

// 根据视频id查询对应的视频信息
func GetVideoByVideoId(videoId int64) (videoMes models.Video, err error) {
	err = db.Preload("Author").Find(&videoMes, "video_id = ?", videoId).Error
	if err != nil {
		log.Fatalln("mysql.GetVideoByVideoId failed", err)
	}
	return
}

func CreateNewVideo(video *models.Video) (err error) {
	if err = db.Select("user_id", "play_url", "cover_url", "title").Create(video).Error; err != nil {
		log.Fatalln("mysql.CreateNewVideo failed", err)
		return
	}
	return nil
}

// GetLastId 获取最后一位视频id
func GetLastId(video *models.Video) (id int64) {
	db.Last(&video)
	fmt.Println(video.Id)
	return video.Id
}
