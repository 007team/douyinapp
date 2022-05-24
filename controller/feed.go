package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})
	FeedResponseFunc(c, 0, CodeSuccess, DemoVideos, time.Now().Unix())
}
