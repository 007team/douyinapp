package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request

func Feed(c *gin.Context) {
	mode, ok := c.Get("mode")
	if !ok {
		log.Println("c.Get(\"mode\") failed")
		return
	}

	if mode.(string) == "NotLoggedIn" {
		// 未登录状态
		var videolist []models.Video
		var err error
		videolist, err = logic.NotLoggedInGetVideo()
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  "视频获取错误",
			})
			return
		}
		// 去敏
		length := len(videolist)
		for i := 0; i < length; i++ {
			videolist[i].Author.Salt = ""
			videolist[i].Author.Password = ""
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "视频获取成功",
			},
			VideoList: videolist,
			NextTime:  time.Now().Unix(),
		})

	} else if mode.(string) == "LoggedIn" {
		// 已登录状态
		//var videolist []models.Video
		userId, ok := c.Get("user_id")
		if !ok {
			log.Println("c.Get(\"user_id\") failed")
			return
		}
		videolist, err := logic.LoggedInGetVideo(userId.(int64))
		if err != nil {
			FeedResponseFunc(c, 1, CodeServerBusy, nil, time.Now().Unix())
			return
		}
		length := len(videolist)
		for i := 0; i < length; i++ {
			videolist[i].Author.Salt = ""
			videolist[i].Author.Password = ""
		}
		FeedResponseFunc(c, 0, CodeSuccess, videolist, time.Now().Unix())
	}

}
