package mysql

import (
	"github.com/007team/douyinapp/models"
)

/**
作用: 获取该用户所拥有的所有视频
传入: 用户名
返回: Video数组
*/
func GetVideoArr(user_id int64) (VideoArr []models.Video) {
	db.Preload("Author").Find(&VideoArr, "user_id = ?", user_id)

	return
}
