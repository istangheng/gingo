统一封装请求响应：

```golang

// 状态码
const (
	OK                  = 0    // 请求成功。一般用于GET与POST请求
	Created             = -201 // 已创建。成功请求并创建了新的资源
	NoContent           = -204 // 无内容。服务器成功处理，但未返回内容。在未更新网页的情况下，可确保浏览器继续显示当前文档
	BadRequest          = -400 // 请求参数错误
	Unauthorized        = -401 // 请求要求用户的身份认证
	Forbidden           = -403 // 没有操作权限
	NotFound            = -404 // 服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置"您所请求的资源无法找到"的个性页面
	InternalServerError = -500 // 服务器内部错误，无法完成请求
)
```

```golang
// 响应具体信息
// MsgFlags 响应信息展示
var MsgFlags = map[int]string{
	OK:                       "请求成功",
	Created:                  "成功请求并创建了新的资源",
	NoContent:                "无内容，服务器成功处理，但未返回内容",
	BadRequest:               "请求参数错误",
	Unauthorized:             "请求要求用户的身份认证",
	Forbidden:                "没有操作权限",
	NotFound:                 "未找到资源",
	InternalServerError:      "服务器内部错误",
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
```

```golang
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Resp(c *gin.Context, errCode int, msg string, data interface{}) {
	if msg == "" {
		msg = ecode.GetMsg(errCode)
	}
	respBody := Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, respBody)
}

// OK 返回正确
func OK(c *gin.Context, data interface{}) {
	errCode := ecode.OK
	msg := ""
	Resp(c, errCode, msg, data)
}

// Fail
func Fail(c *gin.Context, errCode int, data interface{}) {
	msg := ecode.GetMsg(errCode)
	Resp(c, errCode, msg, data)
}

// Error 返回错误
func Error(c *gin.Context, errCode int, msg string) {
	Resp(c, errCode, msg, nil)
}

// BadRequest 请求参数错误
func BadRequest(c *gin.Context, msg string) {
	errCode := ecode.BadRequest
	Resp(c, errCode, msg, nil)
}

// InternalServerError 服务器内部错误
func InternalServerError(c *gin.Context, msg string) {
	errCode := ecode.InternalServerError
	Resp(c, errCode, msg, nil)
}

// Forbidden 请求参数错误
func Forbidden(c *gin.Context, msg string) {
	errCode := ecode.Forbidden
	Resp(c, errCode, msg, nil)
}

// NotFound 未找到资源
func NotFound(c *gin.Context, msg string) {
	Resp(c, ecode.NotFound, msg, nil)
}
```

```golang
// 具体调用
response.Error(ctx, ecode.Forbidden, ecode.GetMsg(ecode.Forbidden))
```
