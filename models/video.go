package models

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	UserId        int64  `json:"-"`
	Author        User   `json:"author" gorm:"foreignKey:UserId"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"gorm:""`
	Title         string `json:"title"`
}
