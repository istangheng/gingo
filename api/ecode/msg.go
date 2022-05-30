package ecode

import "errors"

var MsgFlags = map[int]string{
	OK:                  "请求成功",
	Created:             "成功请求并创建了新的资源",
	NoContent:           "无内容，服务器成功处理，但未返回内容",
	BadRequest:          "请求参数错误",
	Unauthorized:        "请求要求用户的身份认证",
	Forbidden:           "没有操作权限",
	NotFound:            "未找到资源",
	InternalServerError: "服务器内部错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[InternalServerError]
}

func GetError(code int) error {
	return errors.New(GetMsg(code))
}
