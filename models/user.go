// Package models 存放数据模型
package models

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey" `
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:follow_count"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count"`
	Password      string `gorm:"column:password"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:is_follow"`
	Salt          string
}
