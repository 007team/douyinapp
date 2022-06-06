package redis

import (
	"github.com/go-redis/redis"
	"log"
)

// 添加videoId 到点赞数zset 和 评论数zset
func Publish(videoId int64) (err error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()                         //点赞数zset
	KeyAllVideoCommentCountZSet := getKeyAllVideoCommentCountZSet() // 评论数zset
	err = rdb.ZAdd(KeyAllVideoZSet, redis.Z{0, videoId}).Err()
	if err != nil {
		log.Println("rdb.ZAdd(KeyAllVideoZSet, redis.Z{0, videoId}) failed", err)
		return err
	}
	err = rdb.ZAdd(KeyAllVideoCommentCountZSet, redis.Z{0, videoId}).Err()
	if err != nil {
		log.Println("rdb.ZAdd(KeyAllVideoCommentCountZSet,redis.Z{0,videoId}) failed", err)
		return err
	}
	return nil
}
