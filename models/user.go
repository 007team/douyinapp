// Package models 存放数据模型
package models

import "gorm.io/gorm"

type User struct {
	// Id            uint  `gorm:"primaryKey" json:"id,omitempty"`
	gorm.Model
	Name          string	`gorm:"type:varchar(100);not null" json:"name,omitempty"`
	FollowCount   int64 	`gorm:"type:int;not null;DEFAULT:0" json:"follow_count,omitempty"`
	FollowerCount int64 	`gorm:"type:int;not null;DEFAULT:0" json:"follower_count,omitempty"`
	Password      string	`gorm:"type:varchar(100);not null" json:"password"`
	IsFollow      bool  	`json:"is_follow,omitempty"`
	Salt          string	`gorm:"type:varchar(100);not null" json:"salt"`
}




//type User struct {
//	Id            int64  `json:"id,omitempty" gorm:"primaryKey" `
//	Name          string `json:"name,omitempty"`
//	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:follow_count"`
//	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count"`
//	Password      string `gorm:"column:password"`
//	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:is_follow"`
//	Salt          string
//}