# Nuxt Gin Starter

[![GitHub stars](https://img.shields.io/github/stars/RapboyGao/nuxt-gin-starter?style=for-the-badge)](https://github.com/RapboyGao/nuxt-gin-starter/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/RapboyGao/nuxt-gin-starter?style=for-the-badge)](https://github.com/RapboyGao/nuxt-gin-starter/network)
[![GitHub issues](https://img.shields.io/github/issues/RapboyGao/nuxt-gin-starter?style=for-the-badge)](https://github.com/RapboyGao/nuxt-gin-starter/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/RapboyGao/nuxt-gin-starter?style=for-the-badge)](https://github.com/RapboyGao/nuxt-gin-starter/pulls)
[![License](https://img.shields.io/badge/license-MIT-0b5fff?style=for-the-badge)](./LICENSE)
[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white&style=for-the-badge)](https://go.dev)
[![Nuxt](https://img.shields.io/badge/Nuxt-4.x-00DC82?logo=nuxt&logoColor=white&style=for-the-badge)](https://nuxt.com)
[![Vue](https://img.shields.io/badge/Vue-3.x-42b883?logo=vuedotjs&logoColor=white&style=for-the-badge)](https://vuejs.org)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178c6?logo=typescript&logoColor=white&style=for-the-badge)](https://www.typescriptlang.org)
[![Gin](https://img.shields.io/badge/Gin-1.11-008ecf?style=for-the-badge)](https://gin-gonic.com)
[![GORM](https://img.shields.io/badge/GORM-1.31-2c3e50?style=for-the-badge)](https://gorm.io)
[![Powered by nuxtGin](https://img.shields.io/badge/powered%20by-nuxtGin-111827?style=for-the-badge)](https://pkg.go.dev/github.com/RapboyGao/nuxtGin)

A production-oriented full-stack starter built on `Nuxt + Gin`, with **typed HTTP Endpoint** and **typed WebSocketEndpoint** workflow.

一个面向生产的全栈模板，基于 `Nuxt + Gin`，核心是 **类型化 HTTP Endpoint** 与 **类型化 WebSocketEndpoint** 的前后端协作流。

---

## Why This Starter / 为什么选它

- Endpoint-first backend design (`Go`) and generated client-first frontend usage (`TypeScript`).
- One source of truth for API contracts: models in Go.
- Built-in HTTP + WebSocket code generation, runtime validators, and typed helpers.
- Fast local iteration with one command.
- Clear directory boundaries for team collaboration.

- 后端以 Endpoint 为中心定义接口，前端以生成客户端为中心消费接口。
- Go 模型即协议来源，避免前后端协议漂移。
- 内置 HTTP + WebSocket 代码生成、运行时校验与类型化辅助方法。
- 一条命令启动本地开发。
- 目录职责清晰，适合多人协作。

---

## Quick Start / 快速开始

### Requirements / 环境要求

- `Go 1.24+`
- `Node.js 20+`
- `pnpm 9+`

### Install & Run / 安装并启动

```bash
pnpm install
pnpm dev
```

### Build / 构建

```bash
pnpm build
```

### Common Scripts / 常用脚本

```bash
pnpm dev        # run Nuxt + Gin in development
pnpm build      # production build
pnpm cleanup    # clean generated artifacts
pnpm update:dep # update toolchain dependencies
```

> Note: this project no longer uses Air.
>
> 说明：本项目已不再使用 Air。

---

## Project Layout / 项目结构

```text
nuxt-gin-starter/
├── main.go
├── nuxt.config.ts
├── server.config.json
├── server/
│   ├── api/
│   │   ├── index.go
│   │   ├── Product.go
│   │   ├── ProductCRUD.go
│   │   └── WebSocketDemo.go
│   └── model/
│       ├── DB.go
│       └── Example.Product.go
└── vue/
    ├── pages/
    ├── components/
    └── composables/
        ├── auto-generated-api.ts
        └── auto-generated-ws.ts
```

---

## Endpoint (HTTP) Deep Guide / Endpoint（HTTP）完整指南

### 1) Define Models in Go / 在 Go 中定义模型

```go
type ProductCreateRequest struct {
    Name  string  `json:"name" tsdoc:"Product name / 商品名"`
    Price float64 `json:"price" tsdoc:"Product price / 价格"`
    Code  string  `json:"code" tsdoc:"Product code / 编码"`
}

type ProductModelResponse struct {
    ID        uint    `json:"id" tsdoc:"Product id"`
    Name      string  `json:"name" tsdoc:"Product name"`
    Price     float64 `json:"price" tsdoc:"Product price"`
    Code      string  `json:"code" tsdoc:"Product code"`
    CreatedAt int64   `json:"createdAt" tsdoc:"Created timestamp(ms)"`
    UpdatedAt int64   `json:"updatedAt" tsdoc:"Updated timestamp(ms)"`
}
```

### 2) Pick Endpoint Constructor / 选择 Endpoint 构造器

- `endpoint.NewEndpointNoParams`: only request body, no path/query/header/cookie params.
- `endpoint.NewEndpointNoBody`: no body, supports path/query/header/cookie params.
- `endpoint.NewEndpoint`: fully custom generic endpoint.

- `endpoint.NewEndpointNoParams`：仅 body。
- `endpoint.NewEndpointNoBody`：无 body，可带 path/query/header/cookie。
- `endpoint.NewEndpoint`：完整泛型版本。

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
    "CreateProduct",
    endpoint.HTTPMethodPost,
    "/products",
    func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
        // ...
        return ProductModelResponse{}, nil
    },
)
```

### 3) Register API Group with ServerAPI / 用 ServerAPI 注册 API 组

```go
var HTTPAPI = endpoint.ServerAPI{
    BasePath:  "/api-go/v1",
    GroupPath: "/api-go/v1",
    Endpoints: []endpoint.EndpointLike{
        ProductCreateEndpoint,
        ProductGetEndpoint,
        ProductUpdateEndpoint,
        ProductDeleteEndpoint,
        ProductListEndpoint,
    },
}
```

### 4) Build Gin + Export TS in one call / 一次调用同时注册路由并导出 TS

```go
if _, err := HTTPAPI.Build(engine, "vue/composables/auto-generated-api.ts"); err != nil {
    return err
}
```

### 5) Frontend Usage (Auto-import request helpers) / 前端使用（自动导入 request 函数）

```ts
const list = await requestListProductsGet({
  query: { Page: 1, PageSize: 20 },
});

await requestCreateProductPost({
  name: 'Nova Lamp',
  price: 129.99,
  code: 'SKU-001',
});
```

### 6) Key TS Features Generated / 生成后的关键能力

- Per-endpoint class (`ListProductsGet`, etc.)
- Request helper functions (`requestListProductsGet`, etc.)
- Runtime validators (`validateProductListResponse`, etc.)
- `tsdoc` comments on fields

- 每个接口有类封装（如 `ListProductsGet`）
- 同时生成函数式调用（如 `requestListProductsGet`）
- 生成运行时校验函数（如 `validateProductListResponse`）
- 字段 `tsdoc` 自动转注释

---

## WebSocketEndpoint Deep Guide / WebSocketEndpoint 完整指南

### 1) Define WS Models / 定义 WS 模型

```go
type wsChatPayload struct {
    User    string `json:"user" tsdoc:"Sender"`
    Content string `json:"content" tsdoc:"Message text"`
}

type wsServerEnvelope struct {
    Type    string `json:"type" tsdoc:"Event type"`
    Client  string `json:"client" tsdoc:"Client id"`
    Message string `json:"message" tsdoc:"Event message"`
    At      int64  `json:"at" tsdoc:"Timestamp(ms)"`
}
```

### 2) Build a typed WebSocketEndpoint / 构建类型化 WebSocketEndpoint

```go
var ChatWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
    ws := endpoint.NewWebSocketEndpoint()
    ws.Name = "ChatDemo"
    ws.Path = "/chat-demo"
    ws.Description = "WebSocket demo with typed message handlers"
    ws.ServerMessageType = reflect.TypeOf(wsServerEnvelope{})

    // IMPORTANT:
    // This controls generated union type and onXxxType helpers in TS.
    ws.MessageTypes = []string{"chat", "error", "pong", "system", "whoami"}

    endpoint.RegisterWebSocketTypedHandler(ws, "chat", func(payload wsChatPayload, ctx *endpoint.WebSocketContext) (any, error) {
        // ...
        return wsServerEnvelope{}, nil
    })

    return ws
}()
```

### 3) Register WebSocketAPI / 注册 WebSocketAPI

```go
var WSAPI = endpoint.WebSocketAPI{
    BasePath:  "/ws-go/v1",
    GroupPath: "/ws-go/v1",
    Endpoints: []endpoint.WebSocketEndpointLike{ChatWebSocketEndpoint},
}
```

### 4) Build + Export TS / 注册并导出 TS

```go
if _, err := WSAPI.Build(engine, "vue/composables/auto-generated-ws.ts"); err != nil {
    return err
}
```

### 5) Frontend Usage / 前端使用

`ChatDemo` class and typed methods are generated from backend metadata.

`ChatDemo` 类与类型化方法由后端元数据自动生成。

```ts
const ws = new ChatDemo({
  serialize: (v) => v,
  deserialize: (v) => ensureWsServerEnvelope(v),
});

const offChat = ws.onChatType((message) => {
  console.log(message.message);
});

ws.send({ type: 'chat', payload: { user: 'demo', content: 'hello' } });
```

When `MessageTypes` is declared, generated TS includes concrete handlers like:

当你声明 `MessageTypes` 后，生成的 TS 会包含具体订阅方法，例如：

- `onChatType`
- `onErrorType`
- `onPongType`
- `onSystemType`
- `onWhoamiType`

---

## Full Backend Bootstrap Pattern / 后端启动推荐写法

```go
func main() {
    nuxtGin.ConfigureGinMode()
    nuxtGin.LogServer()

    engine := gin.New()
    engine.Use(gin.Logger(), gin.Recovery())

    nuxtGin.ServeVue(engine)

    if err := api.BuildAllAPIs(engine); err != nil {
        log.Fatal(err)
    }

    if err := engine.Run(":" + fmt.Sprint(nuxtGin.GetConfig.GinPort)); err != nil {
        log.Fatal(err)
    }
}
```

---

## Frontend + Backend Contract Rules / 前后端协作规则

1. Add/change Go models first.
2. Keep `json` + `tsdoc` tags on every public field.
3. Regenerate clients via normal run/build workflow.
4. Use generated request/ws helpers in Vue.
5. Avoid hand-maintaining duplicated TS types.

1. 先改 Go 模型。
2. 对外字段统一写 `json` + `tsdoc`。
3. 通过常规运行/构建触发客户端更新。
4. Vue 优先使用生成的 request/ws helper。
5. 不手写重复 TS 协议。

---

## FAQ

### Q: Why do I see class/function mismatch errors in Vue?
A: Prefer generated `requestXxx` helper functions for HTTP and explicit class construction (`new ChatDemo(...)`) for WS.

### Q: How to get `onXxxType` methods in WS client?
A: Set `ws.MessageTypes = []string{...}` in your `WebSocketEndpoint`.

### Q: Do I need Air?
A: No.

---

## Ecosystem

- nuxtGin docs: <https://pkg.go.dev/github.com/RapboyGao/nuxtGin>
- Nuxt docs: <https://nuxt.com/docs>
- Gin docs: <https://gin-gonic.com>
- GORM docs: <https://gorm.io>

---

## License

MIT
