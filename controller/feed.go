package controller

import (
	"fmt"
	"github.com/007team/douyinapp/dao/redis"
	"time"

	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Feed same demo video list for every request

func Feed(c *gin.Context) {
	// 可选参数，限制返回视频的最新投稿时间戳，精确到秒。条件：update_time
	//last_time := c.Query("latest_time")
	// 用户登录态，有则去关注列表获取最新视频
	userIdStr,ok := c.Get("user_id") // 获取用户id
	fmt.Println(userIdStr)
	if ok{
		// 登录态
		//userId := userIdStr.(int64)














	}else{
		// 无登录
		var videolist []models.Video
		videolist,err := logic.GetVideo()
		if err!=nil{
			c.JSON(http.StatusOK,Response{
				StatusCode:2,
				StatusMsg: "视频获取错误",

			})
			return
		}
		// 信息脱敏
		length := len(videolist)

		for i:=0;i<length;i++{
			videolist[i].Author.Salt =""
			videolist[i].Author.Password =""
			videolist[i].FavoriteCount = redis.FavoriteCountById(videolist[i].Id)
		}

		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "视频获取成功",

			},
			VideoList: videolist,
			NextTime:  time.Now().Unix(),
		})




	}




	//var videolist []models.Video
	//videolist,err := logic.GetVideo()
	//if err!=nil{
	//	c.JSON(http.StatusOK,Response{
	//		StatusCode:2,
	//		StatusMsg: "视频获取错误",
	//
	//	})
	//	return
	//}
	//// 信息脱敏
	//length := len(videolist)
	//
	//for i:=0;i<length;i++{
	//	videolist[i].Author.Salt=""
	//	videolist[i].Author.Password=""
	//}
	//
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//		StatusMsg:  "视频获取成功",
	//
	//	},
	//	VideoList: videolist,
	//	NextTime:  time.Now().Unix(),
	//})

}
