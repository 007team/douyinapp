package logic

import (
	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/redis"
	"github.com/007team/douyinapp/models"
	"log"
)

func CreateComment(comment *models.Comment) (err error) {

	if err = mysql.AddComment(comment); err != nil {
		log.Println("AddComment failed")
		return
	}

	if err = mysql.AddVideoCommentCount(comment.VideoId); err != nil {
		log.Println("mysql.AddVideoCommentCount failed")
		return
	}
	// 对redis评论数zset中videoId ++
	redis.AddComment(comment)

	return nil
}

// DeleteComment 删除评论
func DeleteComment(comment *models.Comment, videoid int64) (err error) {

	if err = mysql.SubVideoCommentCount(videoid); err != nil {
		return
	}
	if err = mysql.DelComment(comment); err != nil {
		log.Println("DelComment failed")
		return
	}
	// 对redis评论数zset中videoId --
	redis.SubComment(videoid)

	return nil
}

func GetCommentList(videoId int64) (CommentArr []models.Comment) {
	CommentArr = mysql.GetCommentList(videoId)
	return
}
