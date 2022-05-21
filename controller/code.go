package controller

// Code msg
//
// 如需自定义Code，请提前告知组长

type ResCode int64

const (
	CodeInvalidParam ResCode = 1000 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeSuccess
)

// 存储 Code 及状态描述
var codeMsgMap = map[ResCode]string{
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙，请稍后再试",
	CodeNeedLogin:       "请登录",
	CodeSuccess:         "",
}

// Msg     根据 Code 返回对应的状态描述
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
