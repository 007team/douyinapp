package models

import "time"

type Comment struct {
	Id        int64  `json:"id,omitempty" gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId    int64  `json:"-" gorm:"type:bigint(20) not null"`
	VideoId	  int64  `json:"video_id"  gorm:"type:bigint(20)"`
	Author    User   `json:"author" gorm:"foreignKey:UserId"`
	Content   string `json:"content,omitempty" gorm:"type:mediumtext collate utf8mb4_general_ci NOT NULL"`
	CreatedAt time.Time `json:"create_date"`
	UpdatedAt time.Time	`json:"update_date"`
}

func (Comment) tableName() string {
	return "comments"

}