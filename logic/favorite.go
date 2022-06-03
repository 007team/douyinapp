package logic

import (
	"strconv"

	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/redis"
	"github.com/007team/douyinapp/models"
)

/** 点赞(增加点赞数, 填入个人的点赞列表)
 *  参数: 用户id, 视频id, 点赞还是取消点赞(1则点赞, 2则取消点赞)
 */
func AddFavorite(userId int64, videoId int64, FavoriteJudge int64) (err error) {
	if FavoriteJudge == 1 {
		err = redis.AddFavoriteVideo(userId, videoId)
	} else if FavoriteJudge == 2 {
		err = redis.RemFavoriteVideo(userId, videoId)
	}
	return
}

// 获取点赞的视频队列
func GetFavoriteVideoList(userId int64) (videoList []models.Video, err error) {
	var (
		videoId       int64
		favoriteCount int64
	)

	// 获取视频id队列
	videoIdList, err := redis.GetUserAllFavoriteVideo(userId)
	if err != nil {
		return
	}
	videoList = make([]models.Video, len(videoIdList))
	// 查找对应的数据
	for key, value := range videoIdList {
		videoId, _ = strconv.ParseInt(value, 0, 64)
		// 依据视频id获取视频信息
		videoList[key], err = mysql.GetVideoByVideoId(videoId)
		if err != nil {
			videoList = nil
			return
		}
		// 获取该视频的点赞数
		favoriteCount, err = redis.GetVideoFavoriteCount(videoId)
		// 如果发生错误, 即该有序集合中不存在这个视频的点赞数
		if err == nil {
			videoList[key].FavoriteCount = favoriteCount
		} else {
			err = nil
		}
	}
	return
}
