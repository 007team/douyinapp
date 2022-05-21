package controller

import (
	"github.com/007team/douyinapp/logic"
	"github.com/007team/douyinapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var videolist []models.Video

	videolist,err := logic.GetVideo()
	if err!=nil{
		c.JSON(http.StatusOK,Response{
			StatusCode:2,
			StatusMsg: "视频获取错误",
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videolist,
		NextTime:  time.Now().Unix(),
	})


}
