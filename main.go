package main

import (
	"GinServer/server/api"
	"os"
	"strconv"
	"strings"

	"github.com/RapboyGao/nuxtGin"
	"github.com/RapboyGao/nuxtGin/runtime"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := runtime.DefaultAPIServerConfig(api.AllEndpoints, api.AllWebSocketEndpoints)
	// Starter keeps `server.config.json` as the runtime fallback source so the
	// packaged app and local dev use the same server ports and base URL.
	// Starter 保留 `server.config.json` 作为运行时回退来源，确保打包产物和本地开发共用同一份端口与 base URL。
	cfg.ExportTSOnRun = shouldExportTSOnRun(cfg)
	if err := nuxtGin.RunServerFromConfig(cfg); err != nil {
		panic(err)
	}
}

func shouldExportTSOnRun(cfg runtime.APIServerConfig) bool {
	// Allow explicit environment override so production runtimes can disable
	// TS export while local development keeps the default behavior.
	// 支持通过环境变量显式覆盖，便于生产环境关闭 TS 导出而本地开发保持默认行为。
	if raw := strings.TrimSpace(os.Getenv("NUXT_GIN_EXPORT_TS_ON_RUN")); raw != "" {
		enabled, err := strconv.ParseBool(raw)
		if err == nil {
			return enabled
		}
	}
	return runtime.ResolveGinMode(cfg.GinMode) != gin.ReleaseMode
}
