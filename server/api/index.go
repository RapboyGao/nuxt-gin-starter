package api

import (
	"github.com/RapboyGao/nuxtGin/endpoint"
)

// AllEndpoints
// Register all HTTP endpoints consumed by runtime.DefaultAPIServerConfig.
// They are mounted during server build and exported through the unified TS flow.
// 注册所有 HTTP 端点，供 runtime.DefaultAPIServerConfig 统一装载。
// 这些端点会在服务构建时挂载，并进入统一的 TS 导出流程。
var AllEndpoints = []endpoint.EndpointLike{
	ProductCreateEndpoint,
	ProductGetEndpoint,
	ProductUpdateEndpoint,
	ProductDeleteEndpoint,
	ProductListEndpoint,
}

// AllWebSocketEndpoints
// Register all websocket endpoints consumed by runtime.DefaultAPIServerConfig.
// They participate in route registration and generated client export.
// 注册所有 WebSocket 端点，供 runtime.DefaultAPIServerConfig 统一装载。
// 这些端点会参与路由注册以及生成客户端导出。
var AllWebSocketEndpoints = []endpoint.WebSocketEndpointLike{
	ProductCRUDWebSocketEndpoint,
}
