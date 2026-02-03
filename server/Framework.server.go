// [框架]
// 不要修改或删除本文件
// 本文件用于导出创建适合本程序的gin.Engine

package server

import (
	"GinServer/server/api"
	"GinServer/server/routes"
	v2 "GinServer/server/v2"
	"fmt"
	"log"
	"net/http"

	"github.com/RapboyGao/nuxtGin"
	"github.com/RapboyGao/nuxtGin/endpoint"
	"github.com/RapboyGao/nuxtGin/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 默认的API处理函数集合
// 将默认API处理函数映射到ApiHandleFunctions结构体
var handlers = api.ApiHandleFunctions{DefaultAPI: routes.MyApi}

/**
 * 创建适合本程序的Gin引擎实例
 * 配置了：
 * - 默认中间件（日志、恢复）
 * - CORS跨域支持
 * - Vue前端服务
 * - API路由注册
 */
func CreateServer() *gin.Engine {
	// 创建默认Gin引擎，包含日志和恢复中间件
	engine := gin.Default()

	// 启用CORS支持，使用默认配置
	engine.Use(cors.Default())

	// 配置Vue前端服务（开发环境代理/生产环境静态文件）
	nuxtGin.ServeVue(engine)

	// 注册API路由
	for _, route := range api.GetRoutes(handlers) {
		// 根据路由定义的HTTP方法注册不同类型的路由
		switch route.Method {
		case http.MethodGet:
			engine.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			engine.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			engine.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			engine.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			engine.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	if _, err := endpoint.ApplyEndpoints(engine, v2.AllEndpoints); err != nil {
		panic(err)
	}

	return engine
}

/**
 * 运行Gin服务器
 * 配置了：
 * - Gin服务器端口
 * - 日志记录
 * - 路由注册
 */
func RunServer() {
	// 调用frontend包中的ConfigureGinMode函数，根据项目目录下是否存在node_modules目录来配置Gin框架的运行模式
	// 如果存在node_modules目录，使用调试模式，输出详细日志；如果不存在，使用生产模式，禁用详细日志以提高性能
	nuxtGin.ConfigureGinMode()

	// 调用frontend包中的GetConfig函数，获取前端配置信息，该配置信息包含Gin服务器端口、Nuxt应用端口和应用基础URL等
	var sharedConfig = nuxtGin.GetConfig

	// 将Gin服务器端口转换为字符串，并添加冒号前缀，用于后续启动服务器时指定端口
	port := ":" + fmt.Sprint(sharedConfig.GinPort)

	// 调用utils包中的LogServer函数，记录服务器启动日志，第一个参数为是否使用颜色输出日志，这里传入false表示不使用颜色输出
	// 第二个参数为Gin服务器端口，用于在日志中显示服务器启动的端口信息
	utils.LogServer(false, sharedConfig.GinPort)

	// 调用routes包中的CreateServerRouters函数，创建适合本程序的Gin引擎实例
	// 该函数会配置默认中间件（日志、恢复）、CORS跨域支持、Vue前端服务，并注册API路由
	router := CreateServer()

	// 调用router.Run方法启动Gin服务器，监听指定端口
	// 如果启动过程中出现错误，使用log.Fatal函数记录错误信息并终止程序
	log.Fatal(router.Run(port))
}
