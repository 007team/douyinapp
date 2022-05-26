package logic

import (
	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/models"
)

// GetVideo 获取视频列表
func GetVideo() ([]models.Video, error) {
	var Videolist []models.Video
	var err error

	Videolist, err = mysql.FindVideo()
	if err != nil {
		return nil, err
	}

	return Videolist, nil

}
