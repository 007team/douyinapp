package models

import (
	"time"
)

type Comment struct {
	Id        int64  `json:"id,omitempty"              gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId    int64  `json:"-"                         gorm:"type:bigint(20) not null"`
	Author    User   `json:"author"                    gorm:"foreignKey:UserId"`
	Content   string `json:"content,omitempty"         gorm:"type:mediumtext NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}




//type Comment struct {
//
//	User       User   `json:"user"`
//	Content    string `json:"content,omitempty"`
//	CreateDate string `json:"create_date,omitempty"`
//}