package api

import (
	"github.com/RapboyGao/nuxtGin/endpoint"
)

// AllEndpoints
// Register all HTTP endpoints consumed by runtime.DefaultAPIServerConfig.
// 注册所有 HTTP 端点，供 runtime.DefaultAPIServerConfig 统一装载。
var AllEndpoints = []endpoint.EndpointLike{
	ProductCreateEndpoint,
	ProductGetEndpoint,
	ProductUpdateEndpoint,
	ProductDeleteEndpoint,
	ProductListEndpoint,
}

// AllWebSocketEndpoints
// Register all websocket endpoints consumed by runtime.DefaultAPIServerConfig.
// 注册所有 WebSocket 端点，供 runtime.DefaultAPIServerConfig 统一装载。
var AllWebSocketEndpoints = []endpoint.WebSocketEndpointLike{
	ProductCRUDWebSocketEndpoint,
}
