package controller

import (
	"net/http"
	"strconv"

	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userToken := c.Query("token")
	newClaims, err := jwt.ParseToken(userToken)
	if err != nil {
		// 错误处理
	}
	userId, err := strconv.ParseInt(c.Query("user_id"), 0, 64)
	if err != nil {
		// 该user_id非数字
	}
	if newClaims.UserID != userId {
		// 返回token错误
	}
	videoArr := logic.PublishList(userId)

	c.JSON(http.StatusOK, &VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoArr,
	})
}
