package main

import (
	// 引入前端配置相关的包，用于配置Gin模式和获取配置信息
	"GinServer/server/frontend"
	// 引入路由相关的包，用于创建服务器路由
	"GinServer/server/routes"
	// 引入工具包，用于记录服务器日志等操作
	"GinServer/server/utils"
	// 引入格式化输出包，用于格式化字符串
	"fmt"
	// 引入日志包，用于记录程序运行中的日志信息
	"log"
)

// main 函数是Go程序的入口点，程序从这里开始执行
func main() {
	// 调用frontend包中的ConfigureGinMode函数，根据项目目录下是否存在node_modules目录来配置Gin框架的运行模式
	// 如果存在node_modules目录，使用调试模式，输出详细日志；如果不存在，使用生产模式，禁用详细日志以提高性能
	frontend.ConfigureGinMode()

	// 调用frontend包中的GetConfig函数，获取前端配置信息，该配置信息包含Gin服务器端口、Nuxt应用端口和应用基础URL等
	var sharedConfig = frontend.GetConfig

	// 将Gin服务器端口转换为字符串，并添加冒号前缀，用于后续启动服务器时指定端口
	port := ":" + fmt.Sprint(sharedConfig.GinPort)

	// 调用utils包中的LogServer函数，记录服务器启动日志，第一个参数为是否使用颜色输出日志，这里传入false表示不使用颜色输出
	// 第二个参数为Gin服务器端口，用于在日志中显示服务器启动的端口信息
	utils.LogServer(false, sharedConfig.GinPort)

	// 调用routes包中的CreateServerRouters函数，创建适合本程序的Gin引擎实例
	// 该函数会配置默认中间件（日志、恢复）、CORS跨域支持、Vue前端服务，并注册API路由
	router := routes.CreateServerRouters()

	// 调用router.Run方法启动Gin服务器，监听指定端口
	// 如果启动过程中出现错误，使用log.Fatal函数记录错误信息并终止程序
	log.Fatal(router.Run(port))
}
