package response

// JSONResult 统一封装响应
type JSONResult struct {
	Code int         `json:"code" `
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
