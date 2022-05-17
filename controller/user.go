// Package controller 服务的入口
// 这一层用来进行参数校验，返回响应，请求转发
package controller

import (
	"fmt"

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
	//token := c.Query("token")

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
