package redis

import (
	"github.com/007team/douyinapp/models"
	"log"
	"strconv"
)

func AddComment(comment *models.Comment) (err error) {
	KeyAllVideoCommentCountZSet := getKeyAllVideoCommentCountZSet()
	videoIdStr := strconv.Itoa(int(comment.VideoId))
	err = rdb.ZIncrBy(KeyAllVideoCommentCountZSet, 1, videoIdStr).Err()
	if err != nil {
		log.Println("rdb.ZIncrBy(KeyAllVideoCommentCountZSet, 1, videoIdStr) failed")
		return err
	}
	return
}
