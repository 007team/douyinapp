package redis

import (
	"github.com/007team/douyinapp/models"
	"log"
)

func UserFollowCount(userId int64) (FollowCount int64, err error) {
	KeyUserFollowSet := getKeyUserFollowSet(userId)
	FollowCount, err = rdb.SCard(KeyUserFollowSet).Result()
	if err != nil {
		log.Println("rdb.SCard(KeyUserFollowSet).Result() failed", err)
		return 0, err
	}
	return FollowCount, err
}

func UserFollowerCount(userId int64) (FollowerCount int64, err error) {
	KeyUserFollowerSet := getKeyUserFollowerSet(userId)
	FollowerCount, err = rdb.SCard(KeyUserFollowerSet).Result()
	if err != nil {
		log.Println("rdb.SCard(KeyUserFollowerSet).Result() failed", err)
		return 0, err
	}
	return FollowerCount, err
}

// IsFollowUser 我是否关注了这个用户
func IsFollowUser(user *models.User, myUserId int64) (ok bool, err error) {
	KeyMyFollowSet := getKeyUserFollowSet(myUserId)
	ok, err = rdb.SIsMember(KeyMyFollowSet, user.Id).Result()
	if err != nil {
		log.Println("rdb.SIsMember(KeyMyFollowSet,user.Id).Result() failed", err)
		return false, err
	}
	return ok, err
}
