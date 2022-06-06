package redis

import (
	"log"
	"strconv"
)

// IsFavoriteVideo 判断视频是否被用户点赞
func IsFavoriteVideo(userId, videoId int64) (ok bool, err error) {
	ok, err = rdb.SIsMember(getKeyUserFavoriteSet(userId), videoId).Result()

	if err != nil {
		log.Println("rdb.SIsMember failed", err)
		return false, err
	}
	return ok, nil
}

// VideoFavoriteCount 视频的点赞数
func VideoFavoriteCount(videoId int64) (int64, error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()
	videoIdStr := strconv.Itoa(int(videoId))
	countF, err := rdb.ZScore(KeyAllVideoZSet, videoIdStr).Result()
	if err != nil {
		log.Println("rdb.ZScore(KeyAllVideoZSet, videoIdStr) failed", err)
		return 0, err
	}
	return int64(countF), nil
}

// VideoCommentCount 获取视频的评论数
func VideoCommentCount(videoId int64) (int64, error) {
	KeyAllVideoCommentCountZSet := getKeyAllVideoCommentCountZSet()
	videoIdStr := strconv.Itoa(int(videoId))
	countF, err := rdb.ZScore(KeyAllVideoCommentCountZSet, videoIdStr).Result()
	if err != nil {
		log.Println("rdb.ZScore(KeyAllVideoCommentCountZSet, videoIdStr) failed", err)
		return 0, err
	}
	return int64(countF), nil
}
