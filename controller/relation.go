package controller

import (
	"fmt"
	"github.com/007team/douyinapp/logic"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	userId, ok := c.Get("user_id") // 用户id
	if !ok {
		log.Println(`c.Get("user_id") failed`)
		return
	}
	toUserIdStr := c.Query("to_user_id")    // 对方用户id
	actionTypeStr := c.Query("action_type") // 1-关注 2-取消关注

	fmt.Println("userid : ", userId)
	//userId, err := strconv.ParseInt(userIdStr.(string), 10, 64)
	//if err != nil {
	//	log.Println("userIdStr Atoi failed", err)
	//	return
	//}
	fmt.Println(toUserIdStr)
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		log.Println("toUserIdStr Atoi failed", err)
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		log.Println("action_type Atoi failed", err)
		return
	}
	if err != nil {
		ResponseFunc(c, 1, CodeInvalidParam)
	}

	if actionType == 1 {
		// 关注操作
		err := logic.FollowAction(userId.(int64), toUserId)
		if err != nil {
			log.Println("logic.FollowAction failed", err)
			ResponseFunc(c, 1, CodeServerBusy)
		}

	} else if actionType == 2 {
		// 取消关注
		err := logic.UnfollowAction(userId.(int64), toUserId)
		if err != nil {
			log.Println("logic.FollowAction failed", err)
			ResponseFunc(c, 1, CodeServerBusy)
		}
	} else {
		// 参数错误
		ResponseFunc(c, 1, CodeInvalidParam)
	}

	// 操作成功
	ResponseFunc(c, 0, CodeSuccess)

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	// 获取userId
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		// 参数解析错误
		log.Println("FollowList strconv.ParseInt failed", err)
		FollowListResponseFunc(c, 1, CodeInvalidParam, nil)
	}

	// 业务处理
	users, err := logic.FollowList(userId)
	if err != nil {
		log.Println("logic.FollowList failed", err)
		// 数据库操作错误
		FollowListResponseFunc(c, 1, CodeServerBusy, nil)
	}
	fmt.Println(userId, "的关注列表为")
	fmt.Printf("%v\n", users)
	// 成功返回
	FollowListResponseFunc(c, 0, CodeSuccess, users)

}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	// 获取粉丝列表
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt userIdStr failed", err)
		FollowerListResponseFunc(c, 1, CodeInvalidParam, nil)
	}
	users, err := logic.FollowerList(userId)
	if err != nil {
		log.Println("logic.FollowerList failed", err)
		FollowerListResponseFunc(c, 1, CodeServerBusy, nil)
	}
	fmt.Println(userId, "的关注粉丝为")
	fmt.Printf("%v\n", users)
	FollowerListResponseFunc(c, 0, CodeSuccess, users)

}
