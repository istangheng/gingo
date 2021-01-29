package request

// RegisterRequest 注册请求表单
type RegisterRequest struct {
	Name      string `example:"admin"`
	Role      string `example:"Normal"`
	Telephone string `example:"12345678910"`
	Password  string `example:"123456"`
}

// LoginRequest 登录请求表单
type LoginRequest struct {
	Telephone string `example:"12345678910"`
	Password  string `example:"123456"`
}
