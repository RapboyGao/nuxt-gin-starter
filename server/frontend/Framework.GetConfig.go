// [框架]
// 不要修改或删除本文件

package frontend

import (
	"github.com/arduino/go-paths-helper"   // 文件路径操作工具
	jsoniter "github.com/json-iterator/go" // 高性能JSON处理库
)

/**
 * 前端配置结构
 * 包含与前端服务相关的配置参数
 */
type Config struct {
	GinPort  int    `json:"ginPort"`  // Gin服务器端口
	NuxtPort int    `json:"nuxtPort"` // Nuxt应用端口
	BaseUrl  string `json:"baseUrl"`  // 应用基础URL
}

// 使用高性能JSON解析器实例
var json = jsoniter.ConfigFastest

/**
 * 从配置文件加载配置
 * 读取server.config.json文件并解析到Config结构体
 */
func (config *Config) Acquire() {
	// 创建配置文件路径
	jsonPath := paths.New("server.config.json")

	// 读取文件内容
	bytes, _ := jsonPath.ReadFile()

	// 解析JSON数据到结构体
	json.Unmarshal(bytes, config)
}

// 全局配置实例，程序启动时初始化
var GetConfig Config = func() Config {
	config := Config{} // 创建配置实例
	config.Acquire()   // 从文件加载配置
	return config      // 返回初始化后的配置
}()
