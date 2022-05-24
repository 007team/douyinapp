package controller

import (
	"strconv"

	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	//token := c.PostForm("token")
	//
	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}
	//
	//data, err := c.FormFile("data")
	//if err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	//saveFile := filepath.Join("./public/", finalName)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, Response{
	//	StatusCode: 0,
	//	StatusMsg:  finalName + " uploaded successfully",
	//})
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
	}
	// token与用户id不符
	//if newClaims.UserID != userId {
	//	VideoListResponseFunc(c, 1, CodeInvalidParam, videoArr)
	//}
	videoArr = logic.PublishList(userId)
	VideoListResponseFunc(c, 0, CodeSuccess, videoArr)
	// c.JSON(http.StatusOK, &VideoListResponse{
	// 	Response: Response{
	// 		StatusCode: 0,
	// 	},
	// 	VideoList: videoArr,
	// })
}
