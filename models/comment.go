package models

import "time"

type Comment struct {
	Id        int64  `json:"id,omitempty"              gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId    int64  `json:"-"                         gorm:"type:bigint(20) not null"`
	Author    User   `json:"author"                    gorm:"foreignKey:UserId"`
	Content   string `json:"content,omitempty"         gorm:"type:mediumtext NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Comment) tableName() string {
	return "comments"
}
