package controller

import (
	"net/http"

	"github.com/007team/douyinapp/models"

	"github.com/gin-gonic/gin"
)

// 响应结构体和响应函数
// 例： Response 的响应函数为 ResponseFunc()

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User models.User `json:"user"`
}

type VideoListResponse struct {
	Response
	VideoList []models.Video `json:"video_list"`
}

type UserListResponse struct {
	Response
	UserList []models.User `json:"user_list"`
}

type FeedResponse struct {
	Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []models.Comment `json:"comment_list,omitempty"`
}
type FollowListResponse struct {
	Response
	UserList []models.User `json:"user_list,omitempty"`
}

// StatusMsg 返回状态描述，这个message是直接给用户看的，所以在定义Code时请勿描述业务相关内容
// 例如：当业务处理失败时 (如数据库查询失败时)，请返回CodeServerBusy
// 具体错误码定义请阅读 code.go 文件

func ResponseFunc(c *gin.Context, StatusCode int32, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: StatusCode, // 状态码，0-成功，其他值-失败
		StatusMsg:  code.Msg(),
	})
}

// UserLoginResponseFunc 返回用户登录的响应
func UserLoginResponseFunc(c *gin.Context, StatusCode int32, code ResCode, userid int64, token string) {
	c.JSON(http.StatusOK, &UserLoginResponse{
		Response: Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		UserId: userid,
		Token:  token,
	})
}

func UserResponseFunc(c *gin.Context, StatusCode int32, code ResCode, user models.User) {
	c.JSON(http.StatusOK, &UserResponse{
		Response: Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		User: user,
	})
}

func VideoListResponseFunc(c *gin.Context, StatusCode int32, code ResCode, videos []models.Video) {
	c.JSON(http.StatusOK, &VideoListResponse{
		Response: Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		VideoList: videos,
	})
}

func UserListResponseFunc(c *gin.Context, StatusCode int32, code ResCode, users []models.User) {
	c.JSON(http.StatusOK, &UserListResponse{
		Response: Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		UserList: users,
	})
}

func FeedResponseFunc(c *gin.Context, StatusCode int32, code ResCode, videos []models.Video, NextTime int64) {
	c.JSON(http.StatusOK, &FeedResponse{
		Response: Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		VideoList: videos,
		NextTime:  NextTime,
	})
}

func CommentListResponseFunc(c *gin.Context, StatusCode int32, code ResCode, comments []models.Comment) {
	c.JSON(http.StatusOK, &CommentListResponse{
		Response: Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		CommentList: comments,
	})
}

func FollowListResponseFunc(c *gin.Context, StatusCode int32, code ResCode, users []models.User) {
	c.JSON(http.StatusOK, &FollowListResponse{
		Response{
			StatusCode: StatusCode,
			StatusMsg:  code.Msg(),
		},
		users,
	})
}
