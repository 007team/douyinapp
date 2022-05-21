package controller

import (
	"strconv"

	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/007team/douyinapp/pkg/jwt"
	"github.com/gin-gonic/gin"
)

/**
 * 个人的视频列表
 */
func PublishList(c *gin.Context) {
	var videoArr []models.Video
	userToken := c.Query("token")
	newClaims, err := jwt.ParseToken(userToken)
	// token解析错误
	if err != nil {
		VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
	}
	// 解析user_id
	userId, err := strconv.ParseInt(c.Query("user_id"), 0, 64)
	// 用户id号非数字
	if err != nil {
		VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
	}
	// token与用户id不符
	if newClaims.UserID != userId {
		VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
	}
	videoArr = logic.PublishList(userId)
	VideoListResponseFunc(c, 0, CodeSuccess, videoArr)
	// c.JSON(http.StatusOK, &VideoListResponse{
	// 	Response: Response{
	// 		StatusCode: 0,
	// 	},
	// 	VideoList: videoArr,
	// })
}
