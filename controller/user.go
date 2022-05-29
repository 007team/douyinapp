// Package controller 服务的入口
// 这一层用来进行参数校验，返回响应，请求转发
package controller

import (
	"errors"
	"log"
	"strconv"

	"github.com/007team/douyinapp/pkg/jwt"
	"gorm.io/gorm"

	"github.com/007team/douyinapp/dao/mysql"

	"github.com/007team/douyinapp/logic"

	"github.com/007team/douyinapp/models"

	"github.com/gin-gonic/gin"
)

var userIdSequence = int64(1)

// Register 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := models.User{
		Name:     username,
		Password: password,
	}
	if len(username) == 0 {
		UserResponseFunc(c, 1, CodeInvalidParam, user)
		return
	}

	token, err := logic.Register(&user)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			UserLoginResponseFunc(c, 1, CodeUserExist, user.Id, token) // 用户已存在
			return
		}
		UserLoginResponseFunc(c, 1, CodeServerBusy, user.Id, token) // 数据查询错误
		return
	}

	// 响应成功 ！
	UserLoginResponseFunc(c, 0, CodeSuccess, user.Id, token)
}

// Login 登录功能
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := models.User{
		Name:     username,
		Password: password,
	}

	// 进行业务处理
	if err := logic.Login(&user); err != nil {
		if err == gorm.ErrRecordNotFound {
			// 此用户不存在
			UserLoginResponseFunc(c, 1, CodeUserNotExist, 0, "")
			return
		}
		if err == mysql.ErrorInvalidUserPassword {
			// 用户密码错误
			UserLoginResponseFunc(c, 1, CodeInvalidPassword, 0, "")
			return
		}
		// mysql数据库查询错误
		log.Fatalln("logic.Logic  数据库查询错误")
		UserLoginResponseFunc(c, 1, CodeServerBusy, 0, "")
		return
	}

	// 生成token
	token, _, err := jwt.GenToken(user.Id)
	if err != nil {
		log.Fatalln("jwt,GenToken 生成token失败")
		return
	}

	// 响应成功 !
	UserLoginResponseFunc(c, 0, CodeSuccess, user.Id, token)

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

	// 响应成功 ！
	UserResponseFunc(c, 0, CodeSuccess, user)

}
