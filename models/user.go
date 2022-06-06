// Package models 存放数据模型
package models

import "time"

type User struct {
	Id            int64  `json:"id,omitempty"             gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;"`
	Name          string `json:"name,omitempty"           gorm:"uniqueIndex:idx_name; type:varchar(64) UNIQUE collate utf8mb4_general_ci not null" `
	FollowCount   int64  `json:"follow_count,omitempty"   gorm:"column:follow_count; type:INT NOT NULL DEFAULT 0 "`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count; type:INT NOT NULL DEFAULT 0"`
	Password      string `json:"-"                        gorm:"column:password; type:varchar(200) NOT NULL"`
	IsFollow      bool   `json:"is_follow,omitempty"      gorm:"-"`
	Salt          string `json:"-"                        gorm:"column:salt;  type:varchar(255) NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (User) tableName() string {
	return "users"
}
