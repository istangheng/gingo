package service

import "gingo/initialize"

// Ping 测试
func Ping() string {
	result, _ := initialize.SetPing()
	return result
}
