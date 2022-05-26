// Package middlewares 中间件处理函数
package middlewares

import (
	"net/http"

	"github.com/007team/douyinapp/controller"
	"github.com/007team/douyinapp/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		token := c.Query("token")
		// 从 query 中获取 token
		if token == "" {
			c.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "token parse failed",
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(token) // 解析token
		if err != nil {
			c.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "token parse failed",
			})
			c.Abort()
			return
		}

		c.Set("user_id", mc.UserID)
		c.Set("token", token)
		c.Next() // 后续的处理函数可以用过c.Get("user_id") 或 c.Get("token") 来获取当前请求的用户信息
	}
}

func JWTAuthMiddlewareForPublish() func(c *gin.Context) {
	return func(c *gin.Context) {

		//token := c.Query("token")
		token := c.PostForm("token")
		// 从 query 中获取 token
		if token == "" {
			c.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "token parse failed",
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(token) // 解析token
		if err != nil {
			c.JSON(http.StatusOK, controller.Response{
				StatusCode: 1,
				StatusMsg:  "token parse failed",
			})
			c.Abort()
			return
		}

		c.Set("user_id", mc.UserID)
		c.Set("token", token)
		c.Next() // 后续的处理函数可以用过c.Get("user_id") 或 c.Get("token") 来获取当前请求的用户信息
	}
}
