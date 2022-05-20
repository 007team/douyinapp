package mysql

import (
	"log"

	"github.com/007team/douyinapp/models"
)

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

/**
作用: 获取该用户所拥有的所有视频
传入: 用户名
返回: Video数组
*/
func GetVideoArr(user_id int64) (VideoArr []models.Video) {
	err := db.Preload("Author").Find(&VideoArr, "user_id = ?", user_id)
	if err != nil {
		log.Fatalln("用户名不存在", err)
		return nil
	}

	return VideoArr
}
