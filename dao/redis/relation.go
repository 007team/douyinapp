package redis

import (
	"log"
)

// FollowAction 向我的关注列表里添加对方的id
func FollowAction(userId, toUserId int64) (err error) {
	userKey := getKeyUserFollowSet(userId) // 我的关注set
	err = rdb.SAdd(userKey, toUserId).Err()
	if err != nil {
		log.Println("FollowAction failed", err)
		return err
	}
	return nil
}

// FollowerActionToUser 向对方的粉丝列表里添加我的id
func FollowerActionToUser(toUserId, userId int64) (err error) {
	toUserIdKey := getKeyUserFollowerSet(toUserId) // 取被关注用户的粉丝set
	err = rdb.SAdd(toUserIdKey, userId).Err()
	if err != nil {
		log.Println("FollowerActionToUser failed", err)
		return err
	}
	return nil

}

// UnFollowAction 在我的关注列表里删除对方id
func UnFollowAction(userId, toUserId int64) (err error) {
	userIdKey := getKeyUserFollowerSet(userId) // 我的关注列表
	err = rdb.SRem(userIdKey, toUserId).Err()
	if err != nil {
		log.Println("UnFollowAction failed", err)
		return err
	}
	return nil
}

// UnFollowerActionToUser 在对方的粉丝列表里删除我的id
func UnFollowerActionToUser(userId, toUserId int64) (err error) {
	toUserIdKey := getKeyUserFollowerSet(toUserId) // 对方的粉丝列表
	err = rdb.SRem(toUserIdKey, userId).Err()
	if err != nil {
		log.Println("UnFollowerActionToUser failed", err)
		return err
	}
	return nil
}

func FollowList(userId int64) (es []string, err error) {
	userIdKey := getKeyUserFollowSet(userId) // 用户的关注列表
	es, err = rdb.SMembers(userIdKey).Result()

	if err != nil {
		log.Println("rdb.SMembers failed", err)
		return nil, err
	}
	return es, err
}

func FollowerList(userId int64) (es []string, err error) {
	// 获取用户的关注列表
	userIdKey := getKeyUserFollowerSet(userId)
	es, err = rdb.SMembers(userIdKey).Result()
	if err != nil {
		log.Println("rdb.SMembers failed", err)
		return nil, err
	}
	return es, err
}
