package mysql

import (
	"github.com/007team/douyinapp/models"
	"log"
)


// 获取评论列表
func GetCommentList(videoId int64)(CommentList []models.Comment){

	err := db.Preload("Author").Where("video_id = ?",videoId).Order("updated_at DESC").Find(&CommentList).Error
	if err!=nil{
		log.Println("dao.GetCommentList error:",err)
	}
	return
}


// 增加评论
func AddComment(comment *models.Comment)(err error){
	if err = db.Preload("Author").Create(comment).Error; err!=nil{
		log.Println("mysql.comment.Addcomment error",err)
		return err
	}
	return nil
}

// 增加视频评论数
func AddVideoCommentCount(videoId int64)(err error){

	var video models.Video
	db.Preload("Author").Where("id = ?", videoId).First(&video)
	video.CommentCount=video.CommentCount+1

	db.Save(&video)

	return
}

// 删除评论
func DelComment(comment *models.Comment)(err error){
	if err = db.Where("id = ?",comment.Id).Delete(comment).Error; err!=nil{
		log.Println("mysql.comment.DelComment error",err)
		return err
	}

	return nil
}
