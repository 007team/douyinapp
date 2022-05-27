package models

import (
	"time"
)

//type Video struct {
//	Author        User  `gorm:"foreignkey:UserId"`
//	//Id            int64  `gorm:"primaryKey" json:"id,omitempty" `
//	gorm.Model
//	UserId		  int64	 `gorm:"type:int;not null" json:"user_id"`
//	PlayUrl       string `gorm:"type:varchar(255)" json:"play_url" json:"play_url,omitempty"`
//	CoverUrl      string `gorm:"type:varchar(255)" json:"cover_url,omitempty"`
//	FavoriteCount int64  `gorm:"type:int;not null" json:"favorite_count,omitempty"`
//	CommentCount  int64  `gorm:"type:int;not null" json:"comment_count,omitempty"`
//	IsFavorite    bool   `gorm:"not null" json:"is_favorite,omitempty"`
//	Title		  string `gorm:"type:varchar(100)" json:"title"`
//}




var LastVideoId int64 = 0

type Video struct {
	Id            int64  `json:"id,omitempty"              gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId        int64  `json:"-"                         gorm:"type:bigint(20)  NOT NULL"`
	Author        User   `json:"author"                    gorm:"foreignKey:UserId"`
	PlayUrl       string `json:"play_url,omitempty"        gorm:"type:varchar(255) NOT NULL"`
	CoverUrl      string `json:"cover_url,omitempty"       gorm:"type:varchar(255) NOT NULL"`
	FavoriteCount int64  `json:"favorite_count,omitempty"  gorm:"type:int  NOT NULL DEFAULT 0"`
	CommentCount  int64  `json:"comment_count,omitempty"   gorm:"type:int  NOT NULL DEFAULT 0"`
	IsFavorite    bool   `json:"is_favorite,omitempty"     gorm:"type:tinyint(1) not null default 0"`

	Title         string `json:"title,omitempty"           gorm:"type:varchar(255)  collate utf8mb4_general_ci NOT NULL DEFAULT ''  "`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Video) tableName() string {
	return "videos"

}

