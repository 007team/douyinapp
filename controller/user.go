// Package controller 服务的入口
// 这一层用来进行参数校验，返回响应，请求转发
package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/007team/douyinapp/logic"

	"github.com/007team/douyinapp/models"

	"github.com/007team/douyinapp/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var userIdSequence = int64(1)

func Register(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")

	token, _, err := jwt.GenToken(userIdSequence)
	fmt.Println(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	//	})
	//} else {
	//	atomic.AddInt64(&userIdSequence, 1)
	//	newUser := User{
	//		Id:   userIdSequence,
	//		Name: username,
	//	}
	//	usersLoginInfo[token] = newUser
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   userIdSequence,
	//		Token:    token,
	//	})
	//}
}

func Login(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")

}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	userIdStr := c.Query("user_id") // 获取用户id
	userid, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Fatalln("UserInfo: user_id invalied", err)
		UserResponseFunc(c, 1, CodeInvalidParam, models.User{})
		return
	}
	user := models.User{
		Id: userid,
	}
	if err := logic.UserInfo(&user); err != nil {
		log.Fatalln("logic.UserInfo failed", err)
		UserResponseFunc(c, 1, CodeServerBusy, user)
		return
	}
	UserResponseFunc(c, 0, CodeSuccess, user)
}
