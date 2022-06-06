package logic

import (
	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/redis"
	"github.com/007team/douyinapp/models"
	"log"
)

// GetVideo 获取视频列表
func NotLoggedInGetVideo() (videos []models.Video, err error) {

	videos, err = mysql.FindVideo()
	if err != nil {
		return nil, err
	}
	for i, video := range videos {
		// 展示视频的点赞数
		videos[i].FavoriteCount, err = redis.VideoFavoriteCount(video.Id)
		if err != nil {
			log.Println("redis.VideoFavoriteCount(video.Id) failed", err)
			return nil, err
		}
		// 展示视频的评论数
		videos[i].CommentCount, err = redis.VideoCommentCount(video.Id)
		if err != nil {
			log.Println("redis.VideoCommentCount(video.Id) failed", err)
			return nil, err
		}
	}
	return videos, nil

}

// LoggedInGetVideo 登录状态下获取feed流
func LoggedInGetVideo(userId int64) (videos []models.Video, err error) {
	videos, err = mysql.FindVideo()
	if err != nil {
		return nil, err
	}
	for i, video := range videos {
		// 是否被用户点赞
		ok, err := redis.IsFavoriteVideo(userId, video.Id)
		if err != nil {
			log.Println("redis.IsFavoriteVideo failed", err)
			return nil, err
		}
		if ok {
			videos[i].IsFavorite = true
		}
		// 展示视频的点赞数
		videos[i].FavoriteCount, err = redis.VideoFavoriteCount(video.Id)
		if err != nil {
			log.Println("redis.VideoFavoriteCount(video.Id) failed", err)
			return nil, err
		}

		// 展示视频的评论数
		videos[i].CommentCount, err = redis.VideoCommentCount(video.Id)
		if err != nil {
			log.Println("redis.VideoCommentCount(video.Id) failed", err)
			return nil, err
		}
	}
	return videos, nil
}
