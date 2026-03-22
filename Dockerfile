# syntax=docker/dockerfile:1

# Build Nuxt static assets / 构建 Nuxt 静态资源
FROM node:22-bookworm-slim AS frontend-builder

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY nuxt.config.ts ./
COPY tsconfig.json ./
COPY server.config.json ./
COPY vue ./vue

RUN npx nuxt generate

# Build Go server / 构建 Go 服务
FROM golang:1.25-bookworm AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY server ./server
COPY server.config.json ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/production .

# Runtime image / 运行时镜像
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates tzdata \
  && rm -rf /var/lib/apt/lists/*

COPY --from=backend-builder /app/production /app/production
COPY --from=frontend-builder /app/vue/.output/public /app/vue/.output/public
COPY server.config.json /app/server.config.json

ENV GIN_MODE=release
ENV NUXT_GIN_MODE=release
ENV NUXT_GIN_EXPORT_TS_ON_RUN=false

EXPOSE 8099

CMD ["./production"]
