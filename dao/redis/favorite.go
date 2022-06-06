package redis

import (
	"log"
	"strconv"
)

func FavoriteAction(userId int64, videoId int64) (err error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()                  // 视频点赞数集合
	KeyUserFavoriteVideoSet := getKeyUserFavoriteSet(userId) //用户的点赞视频集合
	// 开启事务
	pipe := rdb.TxPipeline()
	// 对视频点赞数++
	videoIdstr := strconv.Itoa(int(videoId))
	err = pipe.ZIncrBy(KeyAllVideoZSet, 1, videoIdstr).Err()
	if err != nil {
		log.Println("pipe.ZIncrBy failed", err)
	}
	
	// 在用户的点赞视频列表里写入videoId
	err = pipe.SAdd(KeyUserFavoriteVideoSet, videoId).Err()
	if err != nil {
		log.Println("pipe.SAdd failed", err)
	}
	_, err = pipe.Exec()
	if err != nil {
		log.Println("pipe failed", err)
	}
	return err
}

func UnFavoriteAction(userId, videoId int64) (err error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()                  // 视频点赞数集合
	KeyUserFavoriteVideoSet := getKeyUserFavoriteSet(userId) //用户的点赞视频集合

	//开启事务
	pipe := rdb.TxPipeline()
	// 对视频点赞数--
	videoIdStr := strconv.Itoa(int(videoId))
	err = pipe.ZIncrBy(KeyAllVideoZSet, -1, videoIdStr).Err()
	if err != nil {
		log.Println("pipe.ZIncrBy failed", err)
	}
	// 在用户的点赞视频列表里删除videoId
	err = pipe.SRem(KeyUserFavoriteVideoSet, videoId).Err()
	if err != nil {
		log.Println("pipe.SRem failed", err)
	}
	_, err = pipe.Exec()
	if err != nil {
		log.Println("pipe failed", err)
	}
	return err
}

func FavoriteList(userId int64) (es []string, err error) {
	KeyUserFavoriteSet := getKeyUserFavoriteSet(userId)
	es, err = rdb.SMembers(KeyUserFavoriteSet).Result()
	if err != nil {
		log.Println("rdb.SMembers failed", err)
	}
	return es, err

}
