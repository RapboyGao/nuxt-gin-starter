// [框架]
// 不要修改或删除本文件
// 用于在生产或开发环境代理Vue静态页面

package frontend

import (
	"GinServer/server/utils"
	"fmt"
	"log"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

/**
 * 配置Vue应用服务
 * 根据当前Gin运行模式，选择不同的服务方式：
 * - 生产模式：直接提供静态文件服务
 * - 开发模式：通过反向代理连接到Vue开发服务器
 */
func ServeVue(engine *gin.Engine) {
	// 根路径重定向到应用基础URL
	engine.GET("", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, GetConfig.BaseUrl)
	})

	// 根据Gin运行模式选择服务方式
	if gin.Mode() == gin.ReleaseMode { // 如果是生产模式
		ServeVueProduction(engine) // 生产模式：提供静态文件服务
	} else {
		ServeVueDevelopment(engine) // 开发模式：使用反向代理
	}
}

/**
 * 生产环境Vue服务
 * 直接从打包后的静态文件目录提供服务
 */
func ServeVueProduction(engine *gin.Engine) {
	// 获取Vue静态文件目录路径
	vueDirectory := utils.Dir("vue", ".output", "public")

	// 设置JS文件的MIME类型，确保正确解析
	mime.AddExtensionType(".js", "application/javascript")

	// 为应用基础URL设置静态文件服务
	engine.StaticFS(GetConfig.BaseUrl, http.Dir(vueDirectory.String()))

	// 处理未匹配的路由，返回404页面
	engine.NoRoute(func(ctx *gin.Context) {
		fmt.Println("[Error]: There's no route for", ctx.Request.URL)
		ctx.Header("Content-Type", "text/html")
		ctx.Status(404)
		ctx.File(vueDirectory.Join("404.html").String())
	})
}

/**
 * 开发环境Vue服务
 * 通过反向代理将请求转发到运行中的Nuxt开发服务器
 */
func ServeVueDevelopment(engine *gin.Engine) {
	// 构建目标开发服务器地址
	target := "localhost:" + fmt.Sprint(GetConfig.NuxtPort)

	// 捕获所有匹配基础URL的请求并进行代理
	engine.Any(GetConfig.BaseUrl+"/*filepaths", func(ctx *gin.Context) {
		// 配置反向代理目标URL
		url := &url.URL{}
		url.Scheme = "http" // 转发协议
		url.Host = target   // 目标主机

		// 创建反向代理实例
		proxy := httputil.NewSingleHostReverseProxy(url)

		// 自定义错误处理，在代理失败时返回友好错误信息
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			log.Printf("http: proxy error: %v", err)
			ret := fmt.Sprintf("http proxy error %v", err)
			rw.Write([]byte(ret)) // 将错误信息写入响应
		}

		// 执行代理请求
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})
}
