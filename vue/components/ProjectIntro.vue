<template>
  <section
    class="intro-card"
    :style="glowStyle"
    @mousemove="handleMouseMove"
    @mouseleave="handleMouseLeave"
  >
    <div class="intro-glow intro-glow-a" aria-hidden="true"></div>
    <div class="intro-glow intro-glow-b" aria-hidden="true"></div>

    <header class="intro-header">
      <div class="intro-badge">{{ badge }}</div>
      <h2>{{ title }}</h2>
      <p class="intro-subtitle">{{ subtitleEn }}</p>
      <p class="intro-subtitle-cn">{{ subtitleZh }}</p>
    </header>

    <div class="intro-body">
      <div class="intro-overview">
        <p>{{ overviewEn }}</p>
        <p class="intro-cn">{{ overviewZh }}</p>
      </div>

      <div class="intro-grid">
        <div
          v-for="(feature, idx) in features"
          :key="feature.key"
          class="intro-item"
          :style="{ '--delay': `${idx * 70}ms` }"
        >
          <h3>{{ feature.title }}</h3>
          <p>{{ feature.descriptionEn }}</p>
          <p class="intro-cn">{{ feature.descriptionZh }}</p>
        </div>
      </div>

      <div class="intro-repo">
        <span>Repository / 仓库地址</span>
        <a class="intro-link" :href="repoUrl" target="_blank" rel="noreferrer">
          {{ repoUrl }}
        </a>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
type IntroFeature = {
  key: string;
  title: string;
  descriptionEn: string;
  descriptionZh: string;
};

const badge = 'Nuxt + Gin';
const title = 'Nuxt Gin Starter';
const subtitleEn = 'A pragmatic starter for building full-stack apps with a Nuxt frontend and a Gin backend.';
const subtitleZh = '一个用于构建 Nuxt 前端 + Gin 后端的实用全栈项目模板。';
const overviewEn =
  'This template focuses on a clean, predictable workflow: the Gin server handles APIs and static delivery, while Nuxt provides the UI and SSR. The structure is intentionally minimal so you can scale it without fighting conventions.';
const overviewZh =
  '本模板强调清晰、可预测的开发体验：Gin 负责 API 与静态资源服务，Nuxt 负责 UI 与 SSR。结构保持精简，便于你在业务增长时稳定扩展。';

const repoUrl = 'https://github.com/RapboyGao/nuxt-gin-starter.git';

const features: ReadonlyArray<IntroFeature> = [
  {
    key: 'workflow',
    title: 'Unified Dev Workflow / 统一开发流程',
    descriptionEn: 'Run frontend and backend together with a single `nuxt-gin` workflow.',
    descriptionZh: '前后端通过一条 `nuxt-gin` 流程统一启动与调试。',
  },
  {
    key: 'endpoint',
    title: 'Endpoint-First API / 以 Endpoint 为中心',
    descriptionEn: 'Define typed endpoints in Go and consume them cleanly from the Nuxt app.',
    descriptionZh: '在 Go 中定义 Endpoint，可自动生成 TypeScript Axios 等前端客户端，保持类型一致并降低调用成本。',
  },
  {
    key: 'production',
    title: 'Production Ready / 面向生产',
    descriptionEn: 'Built-in CORS, static serving, and a sensible project layout for growth.',
    descriptionZh: '内置 CORS 与静态资源服务，结构清晰便于扩展。',
  },
  {
    key: 'extensible',
    title: 'Extensible / 易扩展',
    descriptionEn: 'Drop in models, routes, and UI quickly without fighting the scaffolding.',
    descriptionZh: '快速添加模型、路由与界面，不受脚手架束缚。',
  },
  {
    key: 'separation',
    title: 'Clear Separation / 清晰分层',
    descriptionEn: 'Backend logic stays in Go, frontend logic stays in Vue, and they meet at typed APIs.',
    descriptionZh: '后端逻辑在 Go，前端逻辑在 Vue，通过类型化 API 连接。',
  },
  {
    key: 'iteration',
    title: 'Fast Iteration / 快速迭代',
    descriptionEn: 'Update endpoints and UI in minutes, with minimal boilerplate and a stable base path.',
    descriptionZh: '最少样板代码，稳定的 API 基础路径，改动更快落地。',
  },
];

const glowStyle = ref<Record<string, string>>({
  '--glow-a-x': '88%',
  '--glow-a-y': '6%',
  '--glow-b-x': '8%',
  '--glow-b-y': '92%',
});

const handleMouseMove = (event: MouseEvent) => {
  const target = event.currentTarget as HTMLElement | null;
  if (!target) return;
  const rect = target.getBoundingClientRect();
  if (!rect.width || !rect.height) return;

  const x = ((event.clientX - rect.left) / rect.width) * 100;
  const y = ((event.clientY - rect.top) / rect.height) * 100;

  glowStyle.value = {
    '--glow-a-x': `${x}%`,
    '--glow-a-y': `${y}%`,
    '--glow-b-x': `${100 - x}%`,
    '--glow-b-y': `${100 - y}%`,
  };
};

const handleMouseLeave = () => {
  glowStyle.value = {
    '--glow-a-x': '88%',
    '--glow-a-y': '6%',
    '--glow-b-x': '8%',
    '--glow-b-y': '92%',
  };
};
</script>

<style scoped lang="scss">
.intro-card {
  --fg-main: #f1f4f9;
  --fg-soft: #c8d0dc;
  --fg-muted: #96a4b7;
  --panel: rgba(255, 255, 255, 0.06);
  --panel-strong: rgba(255, 255, 255, 0.1);
  --edge: rgba(255, 255, 255, 0.13);
  margin: 4px 12px 12px;
  padding: 24px;
  border-radius: 24px;
  position: relative;
  overflow: hidden;
  isolation: isolate;
  color: var(--fg-main);
  background:
    radial-gradient(120% 90% at 90% -10%, rgba(140, 214, 255, 0.22), transparent 65%),
    radial-gradient(85% 80% at -15% 110%, rgba(255, 163, 107, 0.18), transparent 70%),
    linear-gradient(135deg, #101418 0%, #111722 44%, #141826 100%);
  border: 1px solid var(--edge);
  box-shadow:
    0 28px 50px rgba(0, 0, 0, 0.38),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  animation: card-in 560ms ease-out both;
}

.intro-glow {
  position: absolute;
  border-radius: 999px;
  filter: blur(36px);
  z-index: -1;
  pointer-events: none;
  opacity: 0.6;
}

.intro-glow-a {
  width: 220px;
  height: 220px;
  left: calc(var(--glow-a-x) - 110px);
  top: calc(var(--glow-a-y) - 110px);
  background: rgba(145, 227, 255, 0.35);
  transition: left 120ms ease-out, top 120ms ease-out;
}

.intro-glow-b {
  width: 200px;
  height: 200px;
  left: calc(var(--glow-b-x) - 100px);
  top: calc(var(--glow-b-y) - 100px);
  background: rgba(255, 171, 122, 0.3);
  transition: left 120ms ease-out, top 120ms ease-out;
}

.intro-header h2 {
  margin: 10px 0 8px;
  font-size: clamp(26px, 4vw, 34px);
  letter-spacing: 0.3px;
  line-height: 1.16;
}

.intro-badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  border-radius: 999px;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: #09131b;
  background: linear-gradient(135deg, #99dcff 0%, #8de5c6 100%);
  box-shadow: 0 10px 22px rgba(116, 206, 240, 0.38);
}

.intro-badge::before {
  content: '';
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #0c2b3b;
  animation: pulse 1.7s ease-in-out infinite;
}

.intro-subtitle {
  margin: 0;
  color: var(--fg-soft);
  line-height: 1.7;
  max-width: 75ch;
}

.intro-subtitle-cn {
  margin: 6px 0 0;
  color: var(--fg-muted);
  font-size: 13px;
  line-height: 1.7;
}

.intro-body {
  margin-top: 20px;
  display: grid;
  gap: 18px;
}

.intro-overview {
  padding: 16px 18px;
  border-radius: 16px;
  background: linear-gradient(140deg, rgba(255, 255, 255, 0.07), rgba(255, 255, 255, 0.03));
  border: 1px solid var(--panel-strong);
  backdrop-filter: blur(8px);
  animation: rise-in 500ms ease-out both;
  animation-delay: 120ms;
}

.intro-overview p {
  margin: 0;
  color: var(--fg-soft);
  line-height: 1.6;
}

.intro-overview .intro-cn {
  margin-top: 8px;
  color: var(--fg-muted);
  font-size: 12px;
}

.intro-grid {
  display: grid;
  gap: 14px;
  grid-template-columns: repeat(auto-fit, minmax(230px, 1fr));
}

.intro-item {
  padding: 14px 15px;
  border-radius: 14px;
  background: var(--panel);
  border: 1px solid rgba(255, 255, 255, 0.09);
  transition:
    transform 220ms ease,
    border-color 220ms ease,
    background-color 220ms ease,
    box-shadow 220ms ease;
  animation: rise-in 500ms ease-out both;
  animation-delay: calc(200ms + var(--delay, 0ms));
}

.intro-item:hover {
  transform: translateY(-4px);
  border-color: rgba(154, 225, 255, 0.5);
  background: rgba(255, 255, 255, 0.095);
  box-shadow: 0 14px 24px rgba(0, 0, 0, 0.25);
}

.intro-item h3 {
  margin: 0 0 8px;
  font-size: 14px;
  color: #eff5fc;
  line-height: 1.45;
}

.intro-item p {
  margin: 0;
  color: var(--fg-soft);
  line-height: 1.5;
  font-size: 14px;
}

.intro-item .intro-cn {
  margin-top: 6px;
  color: var(--fg-muted);
  font-size: 12px;
}

.intro-repo {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  align-items: center;
  padding: 12px 14px;
  border-radius: 12px;
  background: linear-gradient(120deg, rgba(144, 221, 255, 0.14), rgba(157, 240, 201, 0.08));
  border: 1px solid rgba(142, 216, 249, 0.36);
  animation: rise-in 500ms ease-out both;
  animation-delay: 260ms;
}

.intro-repo span {
  color: #bccee0;
  font-size: 12px;
  letter-spacing: 0.02em;
}

.intro-link {
  color: #caf2ff;
  text-decoration: none;
  position: relative;
  transition: color 180ms ease;
}

.intro-link::after {
  content: '';
  position: absolute;
  left: 0;
  bottom: -2px;
  width: 100%;
  height: 1px;
  background: currentColor;
  opacity: 0;
  transform: scaleX(0);
  transform-origin: left center;
  transition: transform 220ms ease, opacity 220ms ease;
}

.intro-link:hover {
  color: #ffffff;
}

.intro-link:hover::after {
  opacity: 1;
  transform: scaleX(1);
}

@media (max-width: 640px) {
  .intro-card {
    margin: 10px;
    padding: 16px;
    border-radius: 18px;
  }

  .intro-grid {
    grid-template-columns: 1fr;
  }
}

@keyframes card-in {
  from {
    opacity: 0;
    transform: translateY(12px) scale(0.992);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes rise-in {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes pulse {
  0%,
  100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.35;
    transform: scale(1.35);
  }
}

@media (prefers-reduced-motion: reduce) {
  .intro-card,
  .intro-overview,
  .intro-item,
  .intro-repo,
  .intro-badge::before {
    animation: none;
  }

  .intro-item,
  .intro-link::after {
    transition: none;
  }
}
</style>
