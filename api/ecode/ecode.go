package ecode

const (
	OK                  = 0   // 请求成功。一般用于GET与POST请求
	Created             = 201 // 已创建。成功请求并创建了新的资源
	NoContent           = 204 // 无内容。服务器成功处理，但未返回内容。在未更新网页的情况下，可确保浏览器继续显示当前文档
	BadRequest          = 400 // 请求参数错误
	Unauthorized        = 401 // 请求要求用户的身份认证
	Forbidden           = 403 // 没有操作权限
	NotFound            = 404 // 服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置"您所请求的资源无法找到"的个性页面
	InternalServerError = 500 // 服务器内部错误，无法完成请求
)
