# Nuxt Gin Starter

基于 `Nuxt 4 + Gin` 的全栈模板，采用 endpoint-first 开发方式。

- 前端：Nuxt + Vue
- 后端：Gin + GORM
- API 形态：类型化 `endpoint.Endpoint` / `endpoint.WebSocketEndpoint`
- 客户端：由 `nuxtGin` 自动生成 TypeScript API/WS 客户端

英文文档： [README.md](./README.md)

## 环境要求

- Go 1.24+
- Node.js 20+
- pnpm 9+

## 快速开始

```bash
pnpm install
pnpm dev
```

默认端口配置见 `server.config.json`。

## 项目结构

```text
nuxt-gin-starter/
├── main.go
├── nuxt.config.ts
├── server.config.json
├── server/
│   ├── api/
│   │   ├── index.go
│   │   ├── ProductCRUD.go
│   │   ├── Product.go
│   │   └── WebSocketDemo.go
│   └── model/
│       ├── DB.go
│       └── Example.Product.go
└── vue/
    ├── components/
    ├── pages/
    └── composables/
        ├── auto-generated-api.ts
        └── auto-generated-ws.ts
```

## Endpoint（HTTP）怎么写

核心流程：定义模型 -> 定义 Endpoint -> 注册到 `ServerAPI`。

### 1. 定义请求/响应模型

```go
type ProductCreateRequest struct {
    Name  string  `json:"name" tsdoc:"商品名称"`
    Price float64 `json:"price" tsdoc:"商品单价"`
    Code  string  `json:"code" tsdoc:"商品编码"`
}

type ProductModelResponse struct {
    ID    uint    `json:"id" tsdoc:"商品ID"`
    Name  string  `json:"name" tsdoc:"商品名称"`
    Price float64 `json:"price" tsdoc:"商品单价"`
    Code  string  `json:"code" tsdoc:"商品编码"`
}
```

### 2. 定义 Endpoint

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
    "CreateProduct",
    endpoint.HTTPMethodPost,
    "/products",
    func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
        // 业务逻辑
        return ProductModelResponse{}, nil
    },
)
```

### 3. 统一注册到 HTTPAPI

```go
const httpBasePath = "/api-go/v1"

var HTTPAPI = endpoint.ServerAPI{
    BasePath:  httpBasePath,
    GroupPath: httpBasePath,
    Endpoints: []endpoint.EndpointLike{
        ProductCreateEndpoint,
        ProductGetEndpoint,
        ProductUpdateEndpoint,
        ProductDeleteEndpoint,
        ProductListEndpoint,
    },
}
```

## WebSocketEndpoint 怎么写

核心流程：定义消息模型 -> 定义 WebSocketEndpoint -> 注册到 `WebSocketAPI`。

### 1. 定义消息模型

```go
type wsChatPayload struct {
    User    string `json:"user" tsdoc:"发送者"`
    Content string `json:"content" tsdoc:"消息内容"`
}

type wsServerEnvelope struct {
    Type    string `json:"type" tsdoc:"服务端消息类型"`
    Client  string `json:"client" tsdoc:"客户端ID"`
    Message string `json:"message" tsdoc:"消息文本"`
    At      int64  `json:"at" tsdoc:"时间戳(毫秒)"`
}
```

### 2. 构建 WebSocketEndpoint

```go
var ChatWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
    ws := endpoint.NewWebSocketEndpoint()
    ws.Name = "ChatDemo"
    ws.Path = "/chat-demo"
    ws.ServerMessageType = reflect.TypeOf(wsServerEnvelope{})

    ws.OnConnect = func(ctx *endpoint.WebSocketContext) error {
        return ctx.Publish(wsServerEnvelope{Type: "system", Client: ctx.ID, Message: "connected", At: time.Now().UnixMilli()})
    }

    endpoint.RegisterWebSocketTypedHandler(ws, "chat", func(payload wsChatPayload, ctx *endpoint.WebSocketContext) (any, error) {
        event := wsServerEnvelope{Type: "chat", Client: ctx.ID, Message: payload.User + ": " + payload.Content, At: time.Now().UnixMilli()}
        return event, ctx.Publish(event)
    })

    return ws
}()
```

### 3. 注册到 WSAPI

```go
const wsBasePath = "/ws-go/v1"

var WSAPI = endpoint.WebSocketAPI{
    BasePath:  wsBasePath,
    GroupPath: wsBasePath,
    Endpoints: []endpoint.WebSocketEndpointLike{
        ChatWebSocketEndpoint,
    },
}
```

## 启动方式

在 `main.go` 中同时挂载 HTTP + WS：

```go
func main() {
    nuxtGin.MustRunServerWithWebSockets(api.HTTPAPI.Endpoints, api.WSAPI.Endpoints)
}
```

## 前端调用示例

生成文件位置：

- HTTP：`vue/composables/auto-generated-api.ts`
- WebSocket：`vue/composables/auto-generated-ws.ts`

HTTP 调用：

```ts
import { ListProductsGet } from '@/composables/auto-generated-api';

const data = await ListProductsGet.request({
  query: { Page: 1, PageSize: 20 },
});
```

WS 调用：

```ts
import { chatDemo } from '@/composables/auto-generated-ws';

const ws = chatDemo();
ws.onType('chat', (msg) => console.log(msg));
ws.send({ type: 'chat', payload: { user: 'demo', content: 'hello' } });
```

## 常用命令

```bash
pnpm dev        # 启动 Nuxt + Gin 开发环境
pnpm build      # 构建
pnpm cleanup    # 清理生成物
pnpm update:dep # 更新工具链依赖
```

## 说明

- 当你修改 Endpoint 模型后，请通过常规开发/构建流程触发客户端更新。

## License

MIT
