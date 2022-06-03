package controller

import (
	"log"

	"github.com/007team/douyinapp/logic"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// 获取参数
	userIdInterface, ok := c.Get("user_id")
	if !ok {
		FavoriteActionError("FavoriteAction:c.Get('user_id') failed", c, CodeInvalidParam)
		return
	}
	videoIdInterface, ok := c.Get("video_id")
	if !ok {
		FavoriteActionError("FavoriteAction:c.Get('video_id') failed", c, CodeInvalidParam)
		return
	}
	actionTypeInterface, ok := c.Get("action_type")
	if !ok {
		FavoriteActionError("FavoriteAction:c.Get('action_type') failed", c, CodeInvalidParam)
		return
	}

	//参数解析转换
	userId, ok := userIdInterface.(int64)
	if !ok {
		FavoriteActionError("FavoriteAction:userIdInterface -> userId failed", c, CodeInvalidParam)
		return
	}
	videoId, ok := videoIdInterface.(int64)
	if !ok {
		FavoriteActionError("FavoriteAction:videoIdInterface -> videoId failed", c, CodeInvalidParam)
		return
	}
	actionType, ok := actionTypeInterface.(int64)
	if !ok {
		FavoriteActionError("FavoriteAction:actionTypeInterface -> actionType failed", c, CodeInvalidParam)
		return
	}

	if err := logic.AddFavorite(userId, videoId, actionType); err != nil {
		FavoriteActionResponseFunc(c, 0, CodeSuccess)
	} else {
		FavoriteActionResponseFunc(c, 1, CodeUserNotExist)
	}

}

func FavoriteActionError(logStr string, c *gin.Context, code ResCode) {
	log.Fatalln(logStr)
	FavoriteActionResponseFunc(c, 1, code)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userIdInterface, ok := c.Get("user_id")
	if !ok {
		FavoriteActionError("FavoriteAction:c.Get('user_id') failed", c, CodeInvalidParam)
		return
	}

	userId, ok := userIdInterface.(int64)
	if !ok {
		FavoriteActionError("FavoriteAction:userIdInterface -> userId failed", c, CodeInvalidParam)
		return
	}
	if videoList, err := logic.GetFavoriteVideoList(userId); err != nil {
		VideoListResponseFunc(c, 0, CodeSuccess, videoList)
	} else {
		VideoListResponseFunc(c, 1, CodeServerBusy, videoList)
	}
}
