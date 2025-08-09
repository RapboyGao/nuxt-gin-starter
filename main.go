package main

import (
	// 引入服务器相关的包，用于创建服务器路由
	"GinServer/server"
	// 引入格式化输出包，用于格式化字符串
)

// main 函数是Go程序的入口点，程序从这里开始执行
func main() {
	// 调用server包中的RunServer函数，启动Gin服务器
	server.RunServer()
}
