package mysql

import (
	"github.com/007team/douyinapp/models"
	"log"
)

func FollowList(es []string) (users []models.User, err error) {
	err = db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowList db.Find failed", err)
		return nil, err
	}
	return users, nil
}

func FollowerList(es []string) (users []models.User, err error) {
	err = db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowerList db.Find failed", err)
		return nil, err
	}
	return users, nil
}
