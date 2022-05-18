package mysql

import (
	"log"

	"github.com/007team/douyinapp/models"
)

func UserInfo(user *models.User) (err error) {
	// 根据 userid 查询数据库
	if err = db.First(user, user.Id).Error; err != nil {
		log.Fatalln("mysql.UserInfo 查询错误", err)
		return err
	}
	return nil
}
