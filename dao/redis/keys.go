package redis

import "strconv"

const (
	Prefix                      = "douyinapp:"
	KeyAllVideoZSet             = "video:all"              // 所有视频  用来记录每个视频的点赞数量
	KeyAllVideoCommentCountZSet = "video:commentCount"     // 所有视频  用来记录所有视频的评论数
	KeyUserFavoriteSet          = "video:userFavorite:"    // 一个用户所有的点赞视频    key：value     用户id：视频id的集合
	KeyUserFollowSet            = "relation:userFollow:"   // 一个用户的关注列表	一个用户一个
	KeyUserFollowerSet          = "relation:userFollower:" // 一个用户的粉丝列表	一个用户一个
)

// getKeyAllVideoZSet 返回所有视频ZSet的key
func getKeyAllVideoZSet() string {
	return Prefix + KeyAllVideoZSet
}

func getKeyAllVideoCommentCountZSet() string {
	return Prefix + KeyAllVideoCommentCountZSet
}

// getKeyUserFavoriteSet 用户的所有点赞视频Set的key
func getKeyUserFavoriteSet(userId int64) string {
	return Prefix + KeyUserFavoriteSet + strconv.Itoa(int(userId))
}

// getKeyUserFollowSet 用户关注列表Set的key
func getKeyUserFollowSet(userId int64) string {
	return Prefix + KeyUserFollowSet + strconv.Itoa(int(userId))
}

// getKeyUserFollowerSet 用户粉丝列表Set的key
func getKeyUserFollowerSet(userId int64) string {
	return Prefix + KeyUserFollowerSet + strconv.Itoa(int(userId))
}
