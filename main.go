package main

import (
	"GinServer/server/api"

	"github.com/RapboyGao/nuxtGin"
	// 引入格式化输出包，用于格式化字符串
)

// main 函数是Go程序的入口点，程序从这里开始执行
func main() {
	// 调用server包中的RunServer函数，启动Gin服务器
	nuxtGin.RunServer(api.AllEndpoints)
}