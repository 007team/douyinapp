package routers

import (
	"github.com/007team/douyinapp/controller"
	"github.com/007team/douyinapp/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/user/", middlewares.JWTAuthMiddleware(), controller.UserInfo)
	apiRouter.POST("/publish/action/", middlewares.JWTAuthMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", middlewares.JWTAuthMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)

	// 评论操作
	apiRouter.POST("/comment/action/", middlewares.JWTAuthMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middlewares.JWTAuthMiddleware(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
