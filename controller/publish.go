package controller

import (
	"log"
	"strconv"

	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	data, err := c.FormFile("data") // 视频数据
	if err != nil {
		ResponseFunc(c, 1, CodeInvalidParam)
	}
	title := c.PostForm("title")
	userId, ok := c.Get("user_id")
	if !ok {
		log.Fatalln("c.Get('user_id') failed")
		return
	}
	video := models.Video{
		UserId: userId.(int64),
		Title:  title,
	}

	// 业务处理
	if err := logic.Publish(c, &video, data); err != nil {
		ResponseFunc(c, 1, CodeServerBusy)
		return
	}
	ResponseFunc(c, 0, CodeSuccess)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	var videoArr []models.Video
	//userToken := c.Query("token")
	//newClaims, err := jwt.ParseToken(userToken)
	// token解析错误
	//if err != nil {
	//	VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
	//}
	// 解析user_id
	userId, err := strconv.ParseInt(c.Query("user_id"), 0, 64)
	// 用户id号非数字
	if err != nil {
		VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
		return
	}
	// token与用户id不符
	//if newClaims.UserID != userId {
	//	VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
	//}
	videoArr, err = logic.PublishList(userId)
	if err != nil {
		VideoListResponseFunc(c, 1, CodeServerBusy, videoArr)
		return
	}
	VideoListResponseFunc(c, 0, CodeSuccess, videoArr)
}
