# Nuxt Gin Starter 🚀

[![GitHub stars](https://img.shields.io/github/stars/RapboyGao/nuxt-gin-starter?style=flat-square)](https://github.com/RapboyGao/nuxt-gin-starter/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/RapboyGao/nuxt-gin-starter?style=flat-square)](https://github.com/RapboyGao/nuxt-gin-starter/network)
[![GitHub issues](https://img.shields.io/github/issues/RapboyGao/nuxt-gin-starter?style=flat-square)](https://github.com/RapboyGao/nuxt-gin-starter/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/RapboyGao/nuxt-gin-starter?style=flat-square)](https://github.com/RapboyGao/nuxt-gin-starter/pulls)
[![License](https://img.shields.io/badge/license-MIT-0b5fff?style=flat-square)](./LICENSE)
[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white&style=flat-square)](https://go.dev)
[![Nuxt](https://img.shields.io/badge/Nuxt-4.x-00DC82?logo=nuxt&logoColor=white&style=flat-square)](https://nuxt.com)
[![Vue](https://img.shields.io/badge/Vue-3.x-42b883?logo=vuedotjs&logoColor=white&style=flat-square)](https://vuejs.org)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178c6?logo=typescript&logoColor=white&style=flat-square)](https://www.typescriptlang.org)
[![Gin](https://img.shields.io/badge/Gin-1.11-008ecf?style=flat-square)](https://gin-gonic.com)
[![GORM](https://img.shields.io/badge/GORM-1.31-2c3e50?style=flat-square)](https://gorm.io)
[![Powered by nuxtGin](https://img.shields.io/badge/powered%20by-nuxtGin-111827?style=flat-square)](https://pkg.go.dev/github.com/RapboyGao/nuxtGin)

A production-oriented full-stack starter built on **Nuxt + Gin**, with typed **HTTP Endpoint** and typed **WebSocketEndpoint** workflow.

---

## English 🇺🇸

### ✨ Why This Starter

- Endpoint-first backend design in Go
- Generated TypeScript clients for HTTP and WebSocket
- One source of truth for API contracts
- Fast local iteration with one command
- Clear frontend/backend boundaries

### 🧩 Quick Start

#### Requirements

- Go `1.24+`
- Node.js `20+`
- pnpm `9+`

#### Install & Run

```bash
pnpm install
pnpm dev
```

#### Build

```bash
pnpm build
```

#### Common Scripts

```bash
pnpm dev        # run Nuxt + Gin in development
pnpm build      # production build
pnpm cleanup    # clean generated artifacts
pnpm update:dep # update toolchain dependencies
```

> This project no longer uses Air.

### 🗂️ Project Layout

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

### 🔧 Endpoint (HTTP) Usage

#### 1) Define models in Go

```go
type ProductCreateRequest struct {
    Name  string  `json:"name" tsdoc:"Product name"`
    Price float64 `json:"price" tsdoc:"Product price"`
    Code  string  `json:"code" tsdoc:"Product code"`
}
```

#### 2) Define endpoint

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
    "CreateProduct",
    endpoint.HTTPMethodPost,
    "/products",
    func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
        return ProductModelResponse{}, nil
    },
)
```

#### 3) Register API group

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

#### 4) Build routes + export TS

```go
if _, err := HTTPAPI.Build(engine, "vue/composables/auto-generated-api.ts"); err != nil {
    return err
}
```

#### 5) Frontend call (auto-import request helpers)

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

### 🔌 WebSocketEndpoint Usage

#### 1) Define WS models

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

#### 2) Define WebSocket endpoint

```go
var ChatWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
    ws := endpoint.NewWebSocketEndpoint()
    ws.Name = "ChatDemo"
    ws.Path = "/chat-demo"
    ws.Description = "WebSocket demo with typed message handlers"
    ws.ServerMessageType = reflect.TypeOf(wsServerEnvelope{})

    // Enables generated onXxxType methods
    ws.MessageTypes = []string{"chat", "error", "pong", "system", "whoami"}

    return ws
}()
```

#### 3) Register WS API group

```go
var WSAPI = endpoint.WebSocketAPI{
    BasePath:  "/ws-go/v1",
    GroupPath: "/ws-go/v1",
    Endpoints: []endpoint.WebSocketEndpointLike{ChatWebSocketEndpoint},
}
```

#### 4) Build routes + export TS

```go
if _, err := WSAPI.Build(engine, "vue/composables/auto-generated-ws.ts"); err != nil {
    return err
}
```

#### 5) Frontend call

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

### 🧠 Backend Bootstrap Pattern

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

## 中文 🇨🇳

### ✨ 为什么选这个模板

- 后端采用 Endpoint-first 设计（Go）
- 自动生成 HTTP/WS TypeScript 客户端
- 协议以 Go 模型为唯一真相源
- 一条命令即可本地联调
- 前后端边界清晰，便于协作

### 🧩 快速开始

#### 环境要求

- Go `1.24+`
- Node.js `20+`
- pnpm `9+`

#### 安装并启动

```bash
pnpm install
pnpm dev
```

#### 构建

```bash
pnpm build
```

#### 常用脚本

```bash
pnpm dev        # 启动 Nuxt + Gin 开发环境
pnpm build      # 生产构建
pnpm cleanup    # 清理生成物
pnpm update:dep # 更新工具链依赖
```

> 本项目已不再使用 Air。

### 🗂️ 项目结构

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

### 🔧 Endpoint（HTTP）用法

#### 1）在 Go 中定义模型

```go
type ProductCreateRequest struct {
    Name  string  `json:"name" tsdoc:"商品名"`
    Price float64 `json:"price" tsdoc:"价格"`
    Code  string  `json:"code" tsdoc:"商品编码"`
}
```

#### 2）定义 Endpoint

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
    "CreateProduct",
    endpoint.HTTPMethodPost,
    "/products",
    func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
        return ProductModelResponse{}, nil
    },
)
```

#### 3）注册 API 组

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

#### 4）注册路由并导出 TS

```go
if _, err := HTTPAPI.Build(engine, "vue/composables/auto-generated-api.ts"); err != nil {
    return err
}
```

#### 5）前端调用（自动导入 request 函数）

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

### 🔌 WebSocketEndpoint 用法

#### 1）定义 WS 模型

```go
type wsChatPayload struct {
    User    string `json:"user" tsdoc:"发送者"`
    Content string `json:"content" tsdoc:"消息文本"`
}

type wsServerEnvelope struct {
    Type    string `json:"type" tsdoc:"事件类型"`
    Client  string `json:"client" tsdoc:"客户端ID"`
    Message string `json:"message" tsdoc:"事件消息"`
    At      int64  `json:"at" tsdoc:"时间戳(毫秒)"`
}
```

#### 2）定义 WebSocketEndpoint

```go
var ChatWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
    ws := endpoint.NewWebSocketEndpoint()
    ws.Name = "ChatDemo"
    ws.Path = "/chat-demo"
    ws.Description = "WebSocket 示例"
    ws.ServerMessageType = reflect.TypeOf(wsServerEnvelope{})

    // 用于生成 onXxxType 方法
    ws.MessageTypes = []string{"chat", "error", "pong", "system", "whoami"}

    return ws
}()
```

#### 3）注册 WS API 组

```go
var WSAPI = endpoint.WebSocketAPI{
    BasePath:  "/ws-go/v1",
    GroupPath: "/ws-go/v1",
    Endpoints: []endpoint.WebSocketEndpointLike{ChatWebSocketEndpoint},
}
```

#### 4）注册路由并导出 TS

```go
if _, err := WSAPI.Build(engine, "vue/composables/auto-generated-ws.ts"); err != nil {
    return err
}
```

#### 5）前端调用

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

### 🧠 后端启动推荐模式

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

## 🌐 Ecosystem

- nuxtGin docs: <https://pkg.go.dev/github.com/RapboyGao/nuxtGin>
- Nuxt docs: <https://nuxt.com/docs>
- Gin docs: <https://gin-gonic.com>
- GORM docs: <https://gorm.io>

## 📄 License

MIT
