# Nuxt Gin Starter 🚀

[![Go](https://img.shields.io/badge/Go-1.24%2B-00ADD8?logo=go&logoColor=white&style=flat-square)](https://go.dev)
[![Nuxt](https://img.shields.io/badge/Nuxt-4.x-00DC82?logo=nuxt&logoColor=white&style=flat-square)](https://nuxt.com)
[![Vue](https://img.shields.io/badge/Vue-3.x-42B883?logo=vuedotjs&logoColor=white&style=flat-square)](https://vuejs.org)
[![TypeScript](https://img.shields.io/badge/TypeScript-enabled-3178C6?logo=typescript&logoColor=white&style=flat-square)](https://www.typescriptlang.org)
[![Gin](https://img.shields.io/badge/Gin-powered-008ECF?style=flat-square)](https://gin-gonic.com)
[![GORM](https://img.shields.io/badge/GORM-ready-5D6D7E?style=flat-square)](https://gorm.io)
[![nuxtGin](https://img.shields.io/badge/nuxtGin-v0.2.20-111827?style=flat-square)](https://pkg.go.dev/github.com/RapboyGao/nuxtGin)
[![nuxt-gin-tools](https://img.shields.io/badge/nuxt--gin--tools-v0.2.22-0B5FFF?style=flat-square)](https://www.npmjs.com/package/nuxt-gin-tools)
[![License](https://img.shields.io/badge/license-MIT-0B5FFF?style=flat-square)](./LICENSE)

Typed full-stack starter for building **Nuxt + Gin** applications with an **endpoint-first API workflow**, **generated TypeScript clients**, and a practical development setup that keeps frontend and backend moving together.

Quick Jump:

- [English](#english)
- [中文](#中文)

---

## English

### ✨ Why This Starter

`nuxt-gin-starter` is designed for teams who want:

- ⚡ a fast local workflow for Nuxt + Go
- 🧠 a typed API contract generated from backend endpoint definitions
- 🔌 typed WebSocket support with generated client helpers
- 🧱 a clean starter structure you can expand without fighting framework glue
- 📦 a path from development to deployment packaging

Instead of treating frontend and backend as two disconnected projects, this starter gives you one project shape, one runtime config, and one API story.

### 🌟 Features

- 🛣️ Endpoint-first HTTP APIs defined in Go
- 🔌 Typed WebSocket endpoint with generated TypeScript client helpers
- 🧾 Shared generated schema output in `auto-generated-types.ts`
- 🧰 Unified CLI workflow via `nuxt-gin-tools`
- 🗃️ GORM + SQLite demo ready out of the box
- 🔄 HTTP mutations broadcast WebSocket `sync` updates
- 🧩 Product CRUD demo for both HTTP and WebSocket flows
- ✅ startup validation for `server.config.json`
- 🧱 service-layer extraction for reusable product business logic

### 📌 Current Stack

- `github.com/RapboyGao/nuxtGin`: `v0.2.20`
- `nuxt-gin-tools`: `0.2.18`
- Nuxt: `4.4.2`
- Vue: `3.x`
- Go: `1.24+`

### 🖼️ Project Shape

```text
.
├── main.go
├── package.json
├── go.mod
├── server.config.json
├── pack.config.ts
├── server/
│   ├── api/
│   │   ├── index.go
│   │   ├── Product.go
│   │   ├── ProductCRUD.go
│   │   └── WebSocketDemo.go
│   ├── model/
│   └── service/
└── vue/
    ├── app.vue
    ├── components/
    ├── composables/
    │   ├── auto-generated-api.ts
    │   ├── auto-generated-ws.ts
    │   └── auto-generated-types.ts
    └── pages/
```

### 🧭 Architecture Overview

#### Backend

- `main.go` boots the app through `runtime.DefaultAPIServerConfig(...)`
- `server/api` contains endpoint definitions and WebSocket route wiring
- `server/service` contains shared business logic for product CRUD
- `server/model` contains database initialization and GORM models

#### Frontend

- `vue/pages` contains the app pages
- `vue/components` contains the demo UI
- `vue/composables` contains generated API/WS/schema files plus local utilities

#### Runtime Flow

- Nuxt serves the UI
- Gin serves APIs and WebSocket endpoints
- `nuxtGin` exports typed TS client files during development/runtime setup
- `nuxt-gin-tools` coordinates local dev, build, and packing

### ⚡ Quick Start

```bash
pnpm install
pnpm dev
```

That gives you:

- Nuxt dev server
- Go watcher / restart flow
- generated API and WebSocket client artifacts

Build:

```bash
pnpm build
```

Cleanup:

```bash
pnpm cleanup
```

Update dependencies:

```bash
pnpm update:dep
```

### ⚙️ Runtime Config

Create or edit `server.config.json`:

```json
{
  "ginPort": 8099,
  "nuxtPort": 3000,
  "baseUrl": "/",
  "killPortBeforeDevelop": true,
  "cleanupBeforeDevelop": false
}
```

Fields:

- `ginPort`: port used by Gin
- `nuxtPort`: port used by Nuxt dev server
- `baseUrl`: frontend base path
- `killPortBeforeDevelop`: free occupied ports before dev starts
- `cleanupBeforeDevelop`: force cleanup/bootstrap before dev

Runtime note:

- ✅ invalid ports or an invalid `baseUrl` fail fast during startup

### 🏁 Server Bootstrap

`main.go` uses the current recommended bootstrap flow:

```go
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
```

### 🧱 HTTP API Pattern

Define request/response structs in Go, then expose endpoints with `endpoint.NewEndpoint...`.

Example:

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
	"CreateProduct",
	endpoint.HTTPMethodPost,
	"/products",
	func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
		// validate + persist + return normalized response
		return ProductModelResponse{}, nil
	},
)
```

Generated frontend call:

```ts
const res = await requestListProductsGet({
  query: { Page: 1, PageSize: 50 },
});
```

### 🔌 WebSocket Pattern

This starter uses `endpoint.WebSocketMessage` as the shared envelope:

```go
ws.ClientMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
ws.ServerMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
```

The endpoint declares both client-side and server-side message types through one `MessageTypes` list:

```go
ws.MessageTypes = []string{
  "create",
  "delete",
  "created",
  "deleted",
  "error",
  "list",
  "sync",
  "system",
  "update",
  "updated",
}
```

Typical setup:

- register server payload types with `RegisterWebSocketServerPayloadType[...]`
- register client message handlers with `RegisterWebSocketTypedHandler(...)`
- use `sendTypedMessage(...)` from the generated client
- subscribe with `onXXXPayload(...)` and `onXXXType(...)`

Client send example:

```ts
ws.sendTypedMessage({
  type: 'list',
  payload: { Page: 1, PageSize: 0 },
});
```

Client receive example:

```ts
const decodeOptions = { decode: parseOverviewPayload };

ws.onListPayload((payload) => {
  items.value = payload.items ?? [];
}, decodeOptions);

ws.onSyncPayload((payload) => {
  items.value = payload.items ?? [];
}, decodeOptions);
```

### 🔄 Shared Product Service Layer

This starter now keeps product CRUD business logic inside:

- `server/service/product_service.go`

That means:

- HTTP and WebSocket paths share the same validation rules
- persistence behavior is consistent across transport layers
- errors like `invalid product id` / `product not found` stay aligned

### 🧪 Demo Behavior

Included demos:

- `ProductCRUDDemo.vue`: HTTP CRUD flow
- `WebSocketDemo.vue`: WebSocket CRUD flow
- `ProjectIntro.vue`: starter overview page

Behavior notes:

- HTTP writes also broadcast WebSocket `sync`
- initial WebSocket connect sends a `system` snapshot
- frontend can also request full list explicitly via `type: "list"`

### 🛠️ Development Workflow

Scripts from `package.json`:

```json
{
  "dev": "nuxt-gin dev",
  "build": "nuxt-gin build",
  "postinstall": "nuxt-gin install",
  "update:dep": "nuxt-gin update",
  "cleanup": "nuxt-gin cleanup",
  "nuxt:dev": "nuxt dev"
}
```

Useful variants:

```bash
# frontend only
pnpm exec nuxt-gin dev --skip-go

# go only
pnpm exec nuxt-gin dev --skip-nuxt

# skip bootstrap checks
pnpm exec nuxt-gin dev --no-cleanup
```

### 📦 Packaging

The starter now uses:

- `pack.config.ts`

Default config:

```ts
import createPackConfig from 'nuxt-gin-tools/src/pack';

export default createPackConfig({});
```

You can expand it with custom zip name, output path, extra files, or package metadata.

When packing:

- the CLI prints a styled banner
- warns on ambiguous config
- fails on invalid config
- prints the final `.7z` path and bundle directory after build

### 🎨 Frontend Notes

- The demo UI already includes a stronger visual direction than a blank starter
- Generated API and WebSocket files live under `vue/composables`
- `useRuntimeConfig().public.ginPort` is exposed in development

### 📝 Notes

- 💨 Air is not used
- 🧩 `nuxtGin` handles typed endpoint + TS export concerns
- 🔒 WebSocket broadcast now uses a serialized write path
- 🗃️ SQLite storage is initialized under `.build/temp/gorm.db`
- ✅ `go test ./...` for the starter is expected to work after dependency bootstrap

### 🌐 Ecosystem

- [`nuxtGin`](https://github.com/RapboyGao/nuxtGin)
- [`nuxt-gin-tools` GitHub](https://github.com/RapboyGao/nuxt-gin-tools.git)
- [`nuxt-gin-tools` NPM](https://www.npmjs.com/package/nuxt-gin-tools)

---

## 中文

### ✨ 为什么用这个 Starter

`nuxt-gin-starter` 面向的是希望把 **Nuxt 前端** 和 **Gin 后端** 放在同一个开发节奏里的团队。

它适合你如果你想要：

- ⚡ 前后端统一的本地开发体验
- 🧠 从 Go endpoint 定义自动生成 TypeScript 客户端
- 🔌 WebSocket 也能保持类型化
- 🧱 一个足够干净、可以长期扩展的项目骨架
- 📦 从开发到打包发布的一条完整路径

### 🌟 主要特性

- 🛣️ 基于 Go 的 Endpoint-first HTTP API
- 🔌 带类型的 WebSocket 端点与客户端辅助
- 🧾 统一输出共享类型到 `auto-generated-types.ts`
- 🧰 通过 `nuxt-gin-tools` 提供统一命令
- 🗃️ 内置 GORM + SQLite 示例
- 🔄 HTTP 写操作会触发 WebSocket `sync`
- 🧩 同时包含 HTTP 与 WebSocket 两套 Product CRUD Demo
- ✅ 启动时校验 `server.config.json`
- 🧱 Product 业务逻辑抽离到 service 层

### 📌 当前技术栈

- `github.com/RapboyGao/nuxtGin`: `v0.2.20`
- `nuxt-gin-tools`: `0.2.18`
- Nuxt: `4.4.2`
- Vue: `3.x`
- Go: `1.24+`

### 🖼️ 目录结构

```text
.
├── main.go
├── package.json
├── go.mod
├── server.config.json
├── pack.config.ts
├── server/
│   ├── api/
│   │   ├── index.go
│   │   ├── Product.go
│   │   ├── ProductCRUD.go
│   │   └── WebSocketDemo.go
│   ├── model/
│   └── service/
└── vue/
    ├── app.vue
    ├── components/
    ├── composables/
    │   ├── auto-generated-api.ts
    │   ├── auto-generated-ws.ts
    │   └── auto-generated-types.ts
    └── pages/
```

### 🧭 架构说明

#### 后端

- `main.go` 通过 `runtime.DefaultAPIServerConfig(...)` 启动
- `server/api` 负责 endpoint 与 WebSocket 路由定义
- `server/service` 负责 Product CRUD 共享业务逻辑
- `server/model` 负责数据库初始化与 GORM 模型

#### 前端

- `vue/pages` 放页面
- `vue/components` 放演示组件
- `vue/composables` 放生成的 API / WS / schema 文件以及本地组合逻辑

#### 运行链路

- Nuxt 负责前端 UI
- Gin 负责 API 与 WebSocket 服务
- `nuxtGin` 负责导出强类型 TS 客户端
- `nuxt-gin-tools` 负责本地开发、构建与打包协同

### ⚡ 快速开始

```bash
pnpm install
pnpm dev
```

这会启动：

- Nuxt 开发服务
- Go watcher / 自动重启流程
- 自动生成的 API / WebSocket 客户端产物

构建：

```bash
pnpm build
```

清理：

```bash
pnpm cleanup
```

更新依赖：

```bash
pnpm update:dep
```

### ⚙️ 运行时配置

编辑 `server.config.json`：

```json
{
  "ginPort": 8099,
  "nuxtPort": 3000,
  "baseUrl": "/",
  "killPortBeforeDevelop": true,
  "cleanupBeforeDevelop": false
}
```

字段说明：

- `ginPort`：Gin 服务端口
- `nuxtPort`：Nuxt 开发端口
- `baseUrl`：前端基础路径
- `killPortBeforeDevelop`：开发前是否释放占用端口
- `cleanupBeforeDevelop`：开发前是否强制 cleanup / bootstrap

运行时说明：

- ✅ 非法端口或非法 `baseUrl` 会在启动时直接报错

### 🏁 服务启动方式

当前推荐写法如下：

```go
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
```

### 🧱 HTTP API 模式

先在 Go 中定义请求/响应结构，再通过 `endpoint.NewEndpoint...` 暴露接口。

示例：

```go
var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
	"CreateProduct",
	endpoint.HTTPMethodPost,
	"/products",
	func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
		// 校验 + 持久化 + 返回规范化结果
		return ProductModelResponse{}, nil
	},
)
```

前端调用示例：

```ts
const res = await requestListProductsGet({
  query: { Page: 1, PageSize: 50 },
});
```

### 🔌 WebSocket 模式

本 Starter 使用 `endpoint.WebSocketMessage` 作为统一信封：

```go
ws.ClientMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
ws.ServerMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
```

端点会统一声明支持的消息类型：

```go
ws.MessageTypes = []string{
  "create",
  "delete",
  "created",
  "deleted",
  "error",
  "list",
  "sync",
  "system",
  "update",
  "updated",
}
```

典型组合方式：

- 用 `RegisterWebSocketServerPayloadType[...]` 注册服务端 payload 类型
- 用 `RegisterWebSocketTypedHandler(...)` 注册客户端消息处理器
- 前端用生成的 `sendTypedMessage(...)`
- 前端订阅 `onXXXPayload(...)` / `onXXXType(...)`

前端发送示例：

```ts
ws.sendTypedMessage({
  type: 'list',
  payload: { Page: 1, PageSize: 0 },
});
```

前端接收示例：

```ts
const decodeOptions = { decode: parseOverviewPayload };

ws.onListPayload((payload) => {
  items.value = payload.items ?? [];
}, decodeOptions);

ws.onSyncPayload((payload) => {
  items.value = payload.items ?? [];
}, decodeOptions);
```

### 🔄 Product Service 层

现在 Product CRUD 的共享业务逻辑已经抽到：

- `server/service/product_service.go`

这样做的好处：

- HTTP 与 WebSocket 复用同一套校验
- 数据持久化逻辑一致
- `invalid product id` / `product not found` 等错误语义保持统一

### 🧪 Demo 说明

内置页面 / 组件：

- `ProductCRUDDemo.vue`：HTTP CRUD
- `WebSocketDemo.vue`：WebSocket CRUD
- `ProjectIntro.vue`：首页介绍

行为说明：

- HTTP 写操作会广播 WebSocket `sync`
- WebSocket 首次连接会先发 `system` 快照
- 前端也可主动发送 `type: "list"` 来拉取全量列表

### 🛠️ 开发流程

`package.json` 中的脚本：

```json
{
  "dev": "nuxt-gin dev",
  "build": "nuxt-gin build",
  "postinstall": "nuxt-gin install",
  "update:dep": "nuxt-gin update",
  "cleanup": "nuxt-gin cleanup",
  "nuxt:dev": "nuxt dev"
}
```

常用变体：

```bash
# 仅前端
pnpm exec nuxt-gin dev --skip-go

# 仅 Go
pnpm exec nuxt-gin dev --skip-nuxt

# 跳过预检查
pnpm exec nuxt-gin dev --no-cleanup
```

### 📦 打包

当前使用：

- `pack.config.ts`

默认配置：

```ts
import createPackConfig from 'nuxt-gin-tools/src/pack';

export default createPackConfig({});
```

后续可以在这里扩展：

- zip 名称
- 输出目录
- 额外复制文件
- package metadata

打包时 CLI 会：

- 输出样式化 banner
- 对有歧义的配置给出 `warn`
- 对非法配置直接 `error`
- 在结束时打印 `.7z` 文件路径和 bundle 目录

### 🎨 前端说明

- 默认 Demo UI 已不是纯空白页
- 自动生成的 API / WebSocket 文件放在 `vue/composables`
- 开发环境会暴露 `useRuntimeConfig().public.ginPort`

### 📝 备注

- 💨 不使用 Air
- 🧩 `nuxtGin` 负责 endpoint 与 TS 导出
- 🔒 WebSocket 广播已使用串行化写入路径
- 🗃️ SQLite 默认存储在 `.build/temp/gorm.db`
- ✅ 完成依赖初始化后，`go test ./...` 应可直接运行

### 🌐 生态

- [`nuxtGin`](https://github.com/RapboyGao/nuxtGin)
- [`nuxt-gin-tools` GitHub](https://github.com/RapboyGao/nuxt-gin-tools.git)
- [`nuxt-gin-tools` NPM](https://www.npmjs.com/package/nuxt-gin-tools)
