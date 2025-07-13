package frontend

import (
	"fmt"

	"github.com/arduino/go-paths-helper" // 文件路径操作工具
	"github.com/gin-gonic/gin"           // Gin Web框架
)

/**
 * 配置Gin框架运行模式
 * 根据项目目录下是否存在node_modules目录决定运行模式：
 * - 存在：开发模式（默认模式，输出详细日志）
 * - 不存在：生产模式（禁用详细日志，提高性能）
 */
func ConfigureGinMode() {
	// 创建指向node_modules目录的路径对象
	path := paths.New("node_modules")

	// 将路径转换为绝对路径（基于当前工作目录）
	path.ToAbs()

	// 判断node_modules目录是否存在
	if path.IsDir() {
		// 开发环境：存在node_modules目录，使用调试模式
		fmt.Println("/node_modules found: using gin.DebugMode")
	} else {
		// 生产环境：不存在node_modules目录，使用生产模式
		fmt.Println("/node_modules not found: using gin.ReleaseMode")
		gin.SetMode(gin.ReleaseMode)
	}
}
