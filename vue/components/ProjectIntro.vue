<template>
  <section class="intro-card">
    <header class="intro-header">
      <div class="intro-badge">Nuxt + Gin</div>
      <h2>Nuxt Gin Starter</h2>
      <p class="intro-subtitle">
        A pragmatic starter for building full-stack apps with a Nuxt frontend
        and a Gin backend.
      </p>
      <p class="intro-subtitle-cn">
        一个用于构建 Nuxt 前端 + Gin 后端的实用全栈项目模板。
      </p>
    </header>

    <div class="intro-body">
      <div class="intro-overview">
        <p>
          This template focuses on a clean, predictable workflow: the Gin server
          handles APIs and static delivery, while Nuxt provides the UI and SSR.
          The structure is intentionally minimal so you can scale it without
          fighting conventions.
        </p>
        <p class="intro-cn">
          本模板强调清晰、可预测的开发体验：Gin 负责 API 与静态资源服务，Nuxt
          负责 UI 与 SSR。 结构保持精简，便于你在业务增长时稳定扩展。
        </p>
      </div>

      <div class="intro-grid">
        <div class="intro-item">
          <h3>Unified Dev Workflow / 统一开发流程</h3>
          <p>
            Run frontend and backend together with a single `nuxt-gin` workflow.
          </p>
          <p class="intro-cn">前后端通过一条 `nuxt-gin` 流程统一启动与调试。</p>
        </div>
        <div class="intro-item">
          <h3>Endpoint-First API / 以 Endpoint 为中心</h3>
          <p>
            Define typed endpoints in Go and consume them cleanly from the Nuxt
            app.
          </p>
          <p class="intro-cn">
            在 Go 中定义 Endpoint，可自动生成 TypeScript Axios
            等前端客户端，保持类型一致并降低调用成本。
          </p>
        </div>
        <div class="intro-item">
          <h3>Production Ready / 面向生产</h3>
          <p>
            Built-in CORS, static serving, and a sensible project layout for
            growth.
          </p>
          <p class="intro-cn">内置 CORS 与静态资源服务，结构清晰便于扩展。</p>
        </div>
        <div class="intro-item">
          <h3>Extensible / 易扩展</h3>
          <p>
            Drop in models, routes, and UI quickly without fighting the
            scaffolding.
          </p>
          <p class="intro-cn">快速添加模型、路由与界面，不受脚手架束缚。</p>
        </div>
        <div class="intro-item">
          <h3>Clear Separation / 清晰分层</h3>
          <p>
            Backend logic stays in Go, frontend logic stays in Vue, and they
            meet at typed APIs.
          </p>
          <p class="intro-cn">
            后端逻辑在 Go，前端逻辑在 Vue，通过类型化 API 连接。
          </p>
        </div>
        <div class="intro-item">
          <h3>Fast Iteration / 快速迭代</h3>
          <p>
            Update endpoints and UI in minutes, with minimal boilerplate and a
            stable base path.
          </p>
          <p class="intro-cn">
            最少样板代码，稳定的 API 基础路径，改动更快落地。
          </p>
        </div>
      </div>

      <div class="intro-repo">
        <span>Repository / 仓库地址</span>
        <a
          class="intro-link"
          href="https://github.com/RapboyGao/nuxt-gin-starter.git"
          target="_blank"
          rel="noreferrer"
        >
          https://github.com/RapboyGao/nuxt-gin-starter.git
        </a>
      </div>
    </div>
  </section>
</template>

<style scoped lang="scss">
.intro-card {
  margin: 10px 10px 10px 10px;
  padding: 20px;
  border-radius: 16px;
  background: linear-gradient(180deg, #141517 0%, #0f1113 100%);
  border: 1px solid rgba(255, 255, 255, 0.08);
  color: #e8e8e8;
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.25);
}

.intro-header h2 {
  margin: 8px 0 6px;
  font-size: 24px;
  letter-spacing: 0.2px;
}

.intro-badge {
  display: inline-flex;
  align-items: center;
  padding: 4px 10px;
  border-radius: 999px;
  font-size: 12px;
  font-weight: 600;
  color: #0b0e11;
  background: #7cc0ff;
}

.intro-subtitle {
  margin: 0;
  color: #b8c0cc;
}

.intro-subtitle-cn {
  margin: 4px 0 0;
  color: #95a2b3;
  font-size: 13px;
}

.intro-body {
  margin-top: 16px;
  display: grid;
  gap: 16px;
}

.intro-overview {
  padding: 12px 14px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.intro-overview p {
  margin: 0;
  color: #aab4c2;
  line-height: 1.6;
}

.intro-overview .intro-cn {
  margin-top: 8px;
  color: #8f9bad;
  font-size: 12px;
}

.intro-grid {
  display: grid;
  gap: 12px;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
}

.intro-item {
  padding: 12px 14px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.intro-item h3 {
  margin: 0 0 6px;
  font-size: 14px;
  color: #dfe7f0;
}

.intro-item p {
  margin: 0;
  color: #9ea7b3;
  line-height: 1.5;
}

.intro-item .intro-cn {
  margin-top: 6px;
  color: #8a96a5;
  font-size: 12px;
}

.intro-repo {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
  padding: 10px 12px;
  border-radius: 10px;
  background: rgba(124, 192, 255, 0.08);
  border: 1px solid rgba(124, 192, 255, 0.2);
}

.intro-repo span {
  color: #9db3c8;
  font-size: 12px;
}

.intro-link {
  color: #7cc0ff;
  text-decoration: none;
}

.intro-link:hover {
  text-decoration: underline;
}
</style>
