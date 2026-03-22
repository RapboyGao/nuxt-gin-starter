package main

import (
	"GinServer/server/api"

	"github.com/RapboyGao/nuxtGin"
	"github.com/RapboyGao/nuxtGin/runtime"
)

func main() {
	cfg := runtime.DefaultAPIServerConfig(api.AllEndpoints, api.AllWebSocketEndpoints)
	// Starter keeps `server.config.json` as the runtime fallback source so the
	// packaged app and local dev use the same server ports and base URL.
	// Starter 保留 `server.config.json` 作为运行时回退来源，确保打包产物和本地开发共用同一份端口与 base URL。
	cfg.ExportTSOnRun = true
	if err := nuxtGin.RunServerFromConfig(cfg); err != nil {
		panic(err)
	}
}
