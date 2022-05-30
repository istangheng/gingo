package api

import "gingo/api/ecode"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Resp(errCode int, msg string, data interface{}) Response {
	if msg == "" {
		msg = ecode.GetMsg(errCode)
	}
	respBody := Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	}
	return respBody
}

// OK 返回正确
func OK(data interface{}) Response {
	errCode := ecode.OK
	msg := ""
	return Resp(errCode, msg, data)
}

// Fail
func Fail(errCode int, data interface{}) Response {
	msg := ecode.GetMsg(errCode)
	return Resp(errCode, msg, data)
}

// Error 返回错误
func Error(errCode int, msg string) Response {
	return Resp(errCode, msg, nil)
}

// NotFound 未找到资源
func NotFound(msg string) Response {
	return Resp(ecode.NotFound, msg, nil)
}
