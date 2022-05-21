package models

import "gorm.io/gorm"

type Comment struct {
	User       User   `gorm:"foreignkey:UserId" json:"user"`
	//Id         int64  `gorm:"primaryKey" json:"id,omitempty"`
	gorm.Model
	UserId	   int64  `gorm:"type:int;not null" json:"uid,omitempty"`
	Content    string `gorm:"type:varchar(500)" json:"content,omitempty"`
	CreateDate string `gorm:"type:varchar(100)" json:"create_date,omitempty"`
}



//type Comment struct {
//
//	User       User   `json:"user"`
//	Content    string `json:"content,omitempty"`
//	CreateDate string `json:"create_date,omitempty"`
//}