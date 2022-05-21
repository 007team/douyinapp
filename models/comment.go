package models

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	UserId     int64  `json:"-"`
	User       User   `json:"user" gorm:"foreignKey:UserId"`
	Content    string `json:"content,omitempty" gorm:"foreignKey:userId"`
	CreateDate string `json:"create_date,omitempty"`
}
