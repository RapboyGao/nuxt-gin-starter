package main

import (
	"GinServer/server/api"

	"github.com/RapboyGao/nuxtGin"
	"github.com/RapboyGao/nuxtGin/runtime"
)

func main() {
	cfg := runtime.DefaultAPIServerConfig(api.AllEndpoints, api.AllWebSocketEndpoints)
	if err := nuxtGin.RunServerFromConfig(cfg); err != nil {
		panic(err)
	}
}
