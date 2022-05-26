// Package models 存放数据模型
package models

import "time"


//type User struct {
//	// Id            uint  `gorm:"primaryKey" json:"id,omitempty"`
//	gorm.Model
//	Name          string	`gorm:"type:varchar(100);not null" json:"name,omitempty"`
//	FollowCount   int64 	`gorm:"type:int;not null;DEFAULT:0" json:"follow_count,omitempty"`
//	FollowerCount int64 	`gorm:"type:int;not null;DEFAULT:0" json:"follower_count,omitempty"`
//	Password      string	`gorm:"type:varchar(100);not null" json:"password"`
//	IsFollow      bool  	`json:"is_follow,omitempty"`
//	Salt          string	`gorm:"type:varchar(100);not null" json:"salt"`
//}


//func (u *User)AfterFind(tx *gorm.DB)(err error){
//
//	u.Password=""
//	u.Salt=""
//
//	return
//}





type User struct {
	Id            int64  `json:"id,omitempty"             gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;"`
	Name          string `json:"name,omitempty"           gorm:"uniqueIndex:idx_name; type:varchar(64) UNIQUE collate utf8mb4_general_ci not null" `
	FollowCount   int64  `json:"follow_count,omitempty"   gorm:"column:follow_count; type:INT NOT NULL DEFAULT 0 "`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count; type:INT NOT NULL DEFAULT 0"`
	Password      string `json:"-"                        gorm:"column:password; type:varchar(200) NOT NULL"`
	IsFollow      bool   `json:"is_follow,omitempty"      gorm:"column:is_follow; type:tinyint(1) NOT NULL DEFAULT 0"`
	Salt          string `json:"-"                        gorm:"column:salt;  type:varchar(255) NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

}

func (User) tableName() string {
	return "users"
}

