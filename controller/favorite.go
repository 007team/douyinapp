package controller

import (
	"github.com/007team/douyinapp/logic"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// FavoriteAction 用户对视频点赞或取消赞
func FavoriteAction(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		log.Println("c.Get(\"user_id\") failed")
		FavoriteActionResponseFunc(c, 1, CodeInvalidParam)
	}
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt videoId failed", err)
		FavoriteActionResponseFunc(c, 1, CodeInvalidParam)
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt actionType failed", err)
		FavoriteActionResponseFunc(c, 1, CodeInvalidParam)
	}

	// 业务逻辑
	if actionType == 1 {
		err := logic.FavoriteAction(userId.(int64), videoId)
		if err != nil {
			log.Println("logic.FavoriteAction failed", err)
			FavoriteActionResponseFunc(c, 1, CodeServerBusy)
		}
	} else if actionType == 2 {
		err := logic.UnFavoriteAction(userId.(int64), videoId)
		if err != nil {
			log.Println("logic.UnFavoriteAction failed", err)
			FavoriteActionResponseFunc(c, 1, CodeServerBusy)
		}
	} else {
		log.Println("FavoriteAction actionType InvalidParam")
		FavoriteActionResponseFunc(c, 1, CodeInvalidParam)
	}

	// 操作成功！
	FavoriteActionResponseFunc(c, 0, CodeSuccess)

}

// FavoriteList 点赞视频列表
func FavoriteList(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		log.Println("c.Get(user_id) failed")
		FavoriteListResponseFunc(c, 1, CodeInvalidParam, nil)
	}

	// 业务逻辑
	videos, err := logic.FavoriteList(userId.(int64))
	if err != nil {
		log.Println("logic.FavoriteList failed", err)
		FavoriteListResponseFunc(c, 1, CodeServerBusy, nil)
	}

	FavoriteListResponseFunc(c, 0, CodeSuccess, videos)
}
