package redis

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
)

// 给视频的点赞数加上一个数值(可增可减)
func incrNumToFavoriteCount(addNum float64, videoId *string) (err error) {
	err = rdb.ZIncrBy(getKeyAllVideoZSet(), addNum, *videoId).Err()
	if err != nil {
		log.Println("redis.incrNumToFavoriteCount 增加失败", err)
	}
	return
}

// 存入个人点赞了的视频列表并增加点赞数
func AddFavoriteVideo(userId int64, videoId int64) (err error) {
	videoIdStr := strconv.Itoa(int(videoId))
	UserFavoriteSetKey := getKeyUserFavoriteSet(userId)
	if ok, _ := rdb.SIsMember(UserFavoriteSetKey, videoIdStr).Result(); !ok {
		err = rdb.SAdd(getKeyUserFavoriteSet(userId), videoIdStr).Err()
		if err != nil {
			log.Println("redis.AddFavoriteVideo.SAdd 点赞视频添加到队列中失败了", err)
			return
		}
		err = incrNumToFavoriteCount(1, &videoIdStr)
	} else {
		log.Println("redis.AddFavoriteVideo 点赞操作重复执行了", err)
	}
	return
}

// 从个人的点赞视频列表中移除并减少点赞数
func RemFavoriteVideo(userId int64, videoId int64) (err error) {
	videoIdStr := strconv.Itoa(int(videoId))
	UserFavoriteSetKey := getKeyUserFavoriteSet(userId)
	if ok, _ := rdb.SIsMember(UserFavoriteSetKey, videoIdStr).Result(); ok {
		err = rdb.SRem(getKeyUserFavoriteSet(userId), videoIdStr).Err()
		if err != nil {
			log.Println("redis.RemFavoriteVideo.SRem 点赞视频从队列中删除失败了", err)
			return
		}
		err = incrNumToFavoriteCount(-1, &videoIdStr)
	} else {
		log.Println("redis.RemFavoriteVideo 点赞操作重复执行了", err)
	}
	return
}

// 获取视频的点赞数
func GetVideoFavoriteCount(videoId int64) (num int64, err error) {
	videoIdStr := strconv.Itoa(int(videoId))
	temp, err := rdb.ZScore(getKeyAllVideoZSet(), videoIdStr).Result()
	if err != nil && err != redis.Nil {
		log.Println("redis.GetVideoFavoriteCount 查询点赞数失败", err)
	}
	num = int64(temp)
	return
}

// 查询该用户是否点赞了视频
func IsVideoFavorite(userId int64, videoId int64) (ok bool, err error) {
	videoIdStr := strconv.Itoa(int(videoId))
	ok, err = rdb.SIsMember(getKeyUserFavoriteSet(userId), videoIdStr).Result()
	if err != nil {
		log.Println("redis.IsVideoFavorite 查询用户是否点赞某个视频失败", err)
	}
	return
}

// 查询该用户的所有点赞视频
func GetUserAllFavoriteVideo(userId int64) (videoIdList []string, err error) {
	videoIdList, err = rdb.SMembers(getKeyUserFavoriteSet(userId)).Result()
	if err != nil {
		log.Println("redis.GetUserAllFavoriteVideo 查询用户的点赞视频失败", err)
	}
	return
}
