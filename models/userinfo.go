package models

type UserInfo struct{
	id 				int 	`json:"id"`
	Name 			string	`json:"name"`
	FollowCount 	int		`json:"follow_count"`
	FollowerCount	int		`json:"follower_count"`
	IsFollow 		bool	`json:"is_follow"`
}
