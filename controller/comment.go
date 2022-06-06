package controller

import (
	"fmt"
	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	var commentReponse models.Comment
	// 获取参数部分
	userIdStr, ok := c.Get("user_id") // 获取用户id
	fmt.Println(userIdStr)
	if !ok {
		log.Println("user_id not exist")
	}
	userId := userIdStr.(int64)

	videoIdStr := c.Query("video_id") // 获取视频id
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println("controller.CommentAction videoid transform failed:", err)
		CommentListResponseFunc(c, 1, CommentError, []models.Comment{})
		return
	}

	actionTypeStr := c.Query("action_type")
	actionType, _ := strconv.Atoi(actionTypeStr)

	// 1:发布 or 2:删除？
	// 执行发布逻辑
	if actionType == 1 {
		// 获取评论内容
		Text := c.Query("comment_text")

		comment := models.Comment{
			UserId:  userId,
			VideoId: videoId,
			Content: Text,
		}

		err := logic.CreateComment(&comment)
		if err != nil {
			log.Println("comment上传失败")
			return
		}
		commentReponse = comment

		CommentResponseFunc(c, 0, CodeSuccess, commentReponse)

		return

		// 执行删除逻辑
	} else if actionType == 2 {
		commentIdStr := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)

		comment := models.Comment{
			Id: commentId,
		}

		err := logic.DeleteComment(comment, videoId)
		if err != nil {
			log.Println("comment删除失败")
			return
		}
		commentReponse = comment

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "删除成功",
		})
		return

	} else {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "请求信息错误",
		})
		return
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoIdStr := c.Query("video_id") // 获取视频id
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Fatalln("controller.CommentList videoid transform failed:", err)
		CommentListResponseFunc(c, 1, CommentError, nil)
		return
	}

	var comment []models.Comment

	comment = logic.GetCommentList(videoId)
	CommentListResponseFunc(c, 0, CodeSuccess, comment)
}
