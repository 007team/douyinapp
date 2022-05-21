package models

import "gorm.io/gorm"

type Video struct {
	Author        User  `gorm:"foreignkey:UserId" json:"author"`
	//Id            int64  `gorm:"primaryKey" json:"id,omitempty" `
	gorm.Model
	UserId		  int64	 `gorm:"type:int;not null" json:"uid,omitempty"`
	PlayUrl       string `gorm:"type:varchar(255)" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"type:varchar(255)" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"type:int;not null" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"type:int;not null" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"not null" json:"is_favorite,omitempty"`
	Title		  string `gorm:"type:varchar(100)" json:"title"`
}


//type Video struct {
//	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
//	Author        User   `json:"author"`
//	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
//	CoverUrl      string `json:"cover_url,omitempty"`
//	FavoriteCount int64  `json:"favorite_count,omitempty"`
//	CommentCount  int64  `json:"comment_count,omitempty"`
//	IsFavorite    bool   `json:"is_favorite,omitempty"`
//}