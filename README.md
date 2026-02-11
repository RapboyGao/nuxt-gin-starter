# Nuxt Gin Starter 🚀

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white&style=flat-square)](https://go.dev)
[![Nuxt](https://img.shields.io/badge/Nuxt-4.x-00DC82?logo=nuxt&logoColor=white&style=flat-square)](https://nuxt.com)
[![Vue](https://img.shields.io/badge/Vue-3.x-42b883?logo=vuedotjs&logoColor=white&style=flat-square)](https://vuejs.org)
[![nuxtGin](https://img.shields.io/badge/nuxtGin-v0.2.16-111827?style=flat-square)](https://pkg.go.dev/github.com/RapboyGao/nuxtGin)
[![nuxt-gin-tools](https://img.shields.io/badge/nuxt--gin--tools-v0.2.16-0b5fff?style=flat-square)](https://www.npmjs.com/package/nuxt-gin-tools)
[![License](https://img.shields.io/badge/license-MIT-0b5fff?style=flat-square)](./LICENSE)

Typed full-stack starter with Nuxt + Gin + GORM, using endpoint-first API design.

➡️ [Jump to 中文](#中文)

## English

### Features

- Endpoint-first HTTP API in Go
- Typed WebSocket endpoint with generated TS client
- Unified schema generation to `auto-generated-types.ts`
- Runtime bootstrap via `RunServerFromConfig`
- HTTP and WebSocket Product CRUD demos side-by-side

### Versions (Current)

- `github.com/RapboyGao/nuxtGin`: `v0.2.16`
- `nuxt-gin-tools`: `0.2.16`
- Nuxt: `4.3.x`

### Quick Start

```bash
pnpm install
pnpm dev
```

Build:

```bash
pnpm build
```

### Project Structure

```text
.
├── main.go
├── server.config.json
├── server/
│   ├── api/
│   │   ├── index.go
│   │   ├── Product.go
│   │   ├── ProductCRUD.go
│   │   └── WebSocketDemo.go
│   └── model/
└── vue/
    ├── components/
    └── composables/
        ├── auto-generated-api.ts
        ├── auto-generated-ws.ts
        └── auto-generated-types.ts
```

### Server Bootstrap (v0.2.16)

`main.go` now uses `RunServerFromConfig`:

```go
func main() {
	cfg := runtime.DefaultAPIServerConfig(api.AllEndpoints, api.AllWebSocketEndpoints)
	if err := nuxtGin.RunServerFromConfig(cfg); err != nil {
		panic(err)
	}
}
```

### HTTP Endpoint Usage

Define request/response structs in Go, then register endpoint:

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
	"CreateProduct",
	endpoint.HTTPMethodPost,
	"/products",
	func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
		// validate + db create
		return ProductModelResponse{}, nil
	},
)
```

Frontend call (auto-import function names):

```ts
const res = await requestListProductsGet({
  query: { Page: 1, PageSize: 50 },
})
```

### WebSocket Endpoint Usage (Current Pattern)

In this repo, WebSocket message envelope is unified to `endpoint.WebSocketMessage`:

```go
ws.ClientMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
ws.ServerMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
```

Server sends typed business payload through `payload`:

```go
func wrapProductWSMessage(eventType string, payload wsProductOverview) endpoint.WebSocketMessage {
	body, _ := json.Marshal(payload)
	return endpoint.WebSocketMessage{
		Type:    eventType,
		Payload: body,
	}
}
```

Client send example:

```ts
ws.sendTypedMessage({
  type: 'list',
  payload: { Page: 1, PageSize: 0 },
})
```

Client receive example:

```ts
ws.onMessage((message) => {
  if (message.type === 'list') {
    // parse message.payload then render items
  }
})
```

### Notes

- This project does not use Air.
- `ProductCRUD.go` HTTP mutations also trigger WebSocket `sync`.
- On first connect, frontend requests list via payload (`type: "list"`), instead of relying on an auto list push.

### Ecosystem

- [`nuxtGin`](https://github.com/RapboyGao/nuxtGin)
- [`nuxt-gin-tools` GitHub](https://github.com/RapboyGao/nuxt-gin-tools.git)
- [`nuxt-gin-tools` NPM](https://www.npmjs.com/package/nuxt-gin-tools)

---

## 中文

### 功能概览

- 基于 Go 的 Endpoint-first HTTP API 设计
- 带类型的 WebSocket 端点与自动生成 TS 客户端
- 统一类型输出到 `auto-generated-types.ts`
- 使用 `RunServerFromConfig` 启动服务
- 前端同时展示 HTTP 与 WebSocket 两套 Product CRUD Demo

### 当前版本

- `github.com/RapboyGao/nuxtGin`: `v0.2.16`
- `nuxt-gin-tools`: `0.2.16`
- Nuxt: `4.3.x`

### 快速开始

```bash
pnpm install
pnpm dev
```

构建：

```bash
pnpm build
```

### 目录结构

```text
.
├── main.go
├── server.config.json
├── server/
│   ├── api/
│   │   ├── index.go
│   │   ├── Product.go
│   │   ├── ProductCRUD.go
│   │   └── WebSocketDemo.go
│   └── model/
└── vue/
    ├── components/
    └── composables/
        ├── auto-generated-api.ts
        ├── auto-generated-ws.ts
        └── auto-generated-types.ts
```

### 服务启动（v0.2.16）

`main.go` 采用 `RunServerFromConfig`：

```go
func main() {
	cfg := runtime.DefaultAPIServerConfig(api.AllEndpoints, api.AllWebSocketEndpoints)
	if err := nuxtGin.RunServerFromConfig(cfg); err != nil {
		panic(err)
	}
}
```

### HTTP Endpoint 用法

先定义 Go 请求/响应结构，再注册端点：

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
	"CreateProduct",
	endpoint.HTTPMethodPost,
	"/products",
	func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
		// 校验 + 入库
		return ProductModelResponse{}, nil
	},
)
```

前端调用（使用自动导入函数）：

```ts
const res = await requestListProductsGet({
  query: { Page: 1, PageSize: 50 },
})
```

### WebSocket Endpoint 用法（当前实现）

本项目统一使用 `endpoint.WebSocketMessage` 作为 WS 信封：

```go
ws.ClientMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
ws.ServerMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
```

服务端把业务结构编码到 `payload`：

```go
func wrapProductWSMessage(eventType string, payload wsProductOverview) endpoint.WebSocketMessage {
	body, _ := json.Marshal(payload)
	return endpoint.WebSocketMessage{
		Type:    eventType,
		Payload: body,
	}
}
```

前端发送：

```ts
ws.sendTypedMessage({
  type: 'list',
  payload: { Page: 1, PageSize: 0 },
})
```

前端接收：

```ts
ws.onMessage((message) => {
  if (message.type === 'list') {
    // 解析 payload 后渲染 items
  }
})
```

### 说明

- 本项目不再使用 Air。
- `ProductCRUD.go` 的 HTTP 写操作会联动触发 WebSocket `sync`。
- 首次连接后，前端通过 `type: "list"` + payload 主动拉取全量列表。

### 生态

- [`nuxtGin`](https://github.com/RapboyGao/nuxtGin)
- [`nuxt-gin-tools` GitHub](https://github.com/RapboyGao/nuxt-gin-tools.git)
- [`nuxt-gin-tools` NPM](https://www.npmjs.com/package/nuxt-gin-tools)
