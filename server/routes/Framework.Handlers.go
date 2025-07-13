package routes

import (
	"GinServer/server/api" // 项目API定义和响应模型

	"github.com/gin-gonic/gin"        // Gin Web框架
	"github.com/golang-module/carbon" // 时间处理库，提供更友好的API
)

/**
 * 默认API处理函数集合
 * 定义了系统内置的基础API接口实现
 */
var defaultApis = api.DefaultAPI{
	// 测试POST接口
	// 返回当前时间戳（毫秒级），用于验证服务可用性
	TestPost: func(c *gin.Context) {
		c.JSON(200, api.TestResponse{
			// 使用carbon库获取当前时间并转换为毫秒级时间戳
			Time: carbon.Now().TimestampMilli(),
		})
	},
}
