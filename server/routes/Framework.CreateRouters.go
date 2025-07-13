// [框架]
// 不要修改或删除本文件
// 本文件用于导出创建适合本程序的gin.Engine

package routes

import (
	"GinServer/server/api"
	"GinServer/server/frontend"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 默认的API处理函数集合
// 将默认API处理函数映射到ApiHandleFunctions结构体
var handlers = api.ApiHandleFunctions{DefaultAPI: defaultApis}

/**
 * 创建适合本程序的Gin引擎实例
 * 配置了：
 * - 默认中间件（日志、恢复）
 * - CORS跨域支持
 * - Vue前端服务
 * - API路由注册
 */
func CreateServerRouters() *gin.Engine {
	// 创建默认Gin引擎，包含日志和恢复中间件
	engine := gin.Default()

	// 启用CORS支持，使用默认配置
	engine.Use(cors.Default())

	// 配置Vue前端服务（开发环境代理/生产环境静态文件）
	frontend.ServeVue(engine)

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

	return engine
}
