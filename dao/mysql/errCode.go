package mysql

import "errors"

var (
	ErrorUserExist           = errors.New("用户已存在")
	ErrorInvalidUserPassword = errors.New("用户密码错误")
)
