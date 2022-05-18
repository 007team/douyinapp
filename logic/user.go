package logic

import (
	"log"

	"github.com/007team/douyinapp/dao/mysql"

	"github.com/007team/douyinapp/models"
)

func UserInfo(user *models.User) (err error) {
	if err = mysql.UserInfo(user); err != nil {
		log.Fatalln("mysql.UserInfo failed", err)
		return err
	}
	return nil
}
