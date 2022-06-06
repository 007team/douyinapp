package logic

import (
	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/redis"
	"github.com/007team/douyinapp/models"
)

// FavoriteAction 用户对视频点赞
func FavoriteAction(userId, videoId int64) (err error) {
	return redis.FavoriteAction(userId, videoId)
}

func UnFavoriteAction(userId, videoId int64) (err error) {
	return redis.UnFavoriteAction(userId, videoId)
}

func FavoriteList(userId int64) (videos []models.Video, err error) {
	// 从redis获取用户的点赞视频列表
	es, err := redis.FavoriteList(userId)
	if err != nil {
		return nil, nil
	}
	// mysql查询视频数据
	if len(es) != 0 {
		videos, err = mysql.FavoriteList(es)
	}
	return videos, err

}
