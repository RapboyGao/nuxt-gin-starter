# Nuxt Gin Starter

A full-stack starter template built with `Nuxt 4 + Gin`, using an endpoint-first workflow.

- Frontend: Nuxt + Vue
- Backend: Gin + GORM
- API style: typed `endpoint.Endpoint` / `endpoint.WebSocketEndpoint`
- Client generation: `nuxtGin` auto-generates TypeScript clients

Chinese README: [README.zh-CN.md](./README.zh-CN.md)

## Requirements

- Go 1.24+
- Node.js 20+
- pnpm 9+

## Quick Start

```bash
pnpm install
pnpm dev
```

Default ports are configured in `server.config.json`.

## Project Structure

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

## How Endpoint Works

You define typed HTTP endpoints in Go, then `nuxtGin` handles routing + TS client generation.

### 1. Define request/response models

```go
type ProductCreateRequest struct {
    Name  string  `json:"name" tsdoc:"Product name"`
    Price float64 `json:"price" tsdoc:"Product unit price"`
    Code  string  `json:"code" tsdoc:"Product code"`
}

type ProductModelResponse struct {
    ID    uint    `json:"id" tsdoc:"Product id"`
    Name  string  `json:"name" tsdoc:"Product name"`
    Price float64 `json:"price" tsdoc:"Product price"`
    Code  string  `json:"code" tsdoc:"Product code"`
}
```

### 2. Define endpoint handler

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
    "CreateProduct",
    endpoint.HTTPMethodPost,
    "/products",
    func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
        // business logic
        return ProductModelResponse{}, nil
    },
)
```

### 3. Register all endpoints in one API group

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

### 4. Start server with HTTP + WS

```go
func main() {
    nuxtGin.MustRunServerWithWebSockets(api.HTTPAPI.Endpoints, api.WSAPI.Endpoints)
}
```

## How WebSocketEndpoint Works

`nuxtGin` provides a typed WebSocket endpoint abstraction. You define:

- path
- server message type
- connect/disconnect behavior
- typed message handlers by message `type`

### 1. Define payload and server message models

```go
type wsChatPayload struct {
    User    string `json:"user" tsdoc:"Sender"`
    Content string `json:"content" tsdoc:"Message text"`
}

type wsServerEnvelope struct {
    Type    string `json:"type" tsdoc:"Server event type"`
    Client  string `json:"client" tsdoc:"Client id"`
    Message string `json:"message" tsdoc:"Event message"`
    At      int64  `json:"at" tsdoc:"Timestamp(ms)"`
}
```

### 2. Build WebSocket endpoint

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

### 3. Register WebSocket API group

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

## Frontend Usage

Generated clients are in:

- HTTP: `vue/composables/auto-generated-api.ts`
- WebSocket: `vue/composables/auto-generated-ws.ts`

Example HTTP usage:

```ts
import { ListProductsGet } from '@/composables/auto-generated-api';

const data = await ListProductsGet.request({
  query: { Page: 1, PageSize: 20 },
});
```

Example WS usage:

```ts
import { chatDemo } from '@/composables/auto-generated-ws';

const ws = chatDemo();
ws.onType('chat', (msg) => console.log(msg));
ws.send({ type: 'chat', payload: { user: 'demo', content: 'hello' } });
```

## Common Commands

```bash
pnpm dev        # start Nuxt + Gin in dev mode
pnpm build      # build project
pnpm cleanup    # clean generated/build artifacts
pnpm update:dep # update toolchain dependencies
```

## Notes

- If you add or change endpoint models, regenerate clients via normal project workflow (`pnpm dev` / build pipeline).

## License

MIT
