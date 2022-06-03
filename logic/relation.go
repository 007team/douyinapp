package logic

import (
	"github.com/007team/douyinapp/dao/mysql"
	"github.com/007team/douyinapp/dao/redis"
	"github.com/007team/douyinapp/models"
)

func FollowAction(userId, toUserId int64) (err error) {
	// 在我的关注列表里添加对方用户
	err = redis.FollowAction(userId, toUserId)
	// 在对方的粉丝列表里添加我
	err = redis.FollowerActionToUser(toUserId, userId)

	return err
}

func UnfollowAction(userId, toUserId int64) (err error) {
	// 在我的关注列表中删除对方用户
	err = redis.UnFollowAction(userId, toUserId)
	// 在对方的粉丝列表中删除我的id
	err = redis.UnFollowerActionToUser(userId, toUserId)

	return err
}

func FollowList(userId int64) (users []models.User, err error) {
	// 返回用户的关注列表
	es, err := redis.FollowList(userId)
	if err != nil {
		return nil, nil
	}
	// mysql查询用户
	if len(es) != 0 {
		users, err = mysql.FollowList(es)
	}

	return users, err
}

func FollowerList(userId int64) (users []models.User, err error) {
	// 获取用户的粉丝列表
	es, err := redis.FollowerList(userId)
	if err != nil {
		return nil, nil
	}
	if len(es) != 0 {
		users, err = mysql.FollowList(es)
	}

	return users, err
}
