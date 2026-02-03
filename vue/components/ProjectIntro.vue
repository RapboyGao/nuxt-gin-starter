<template>
  <section class="intro-card">
    <header class="intro-header">
      <div class="intro-badge">Nuxt + Gin</div>
      <h2>Nuxt Gin Starter</h2>
      <p class="intro-subtitle">
        A pragmatic starter for building full-stack apps with a Nuxt frontend and a Gin backend.
      </p>
    </header>

    <div class="intro-body">
      <div class="intro-grid">
        <div class="intro-item">
          <h3>Unified Dev Workflow</h3>
          <p>Run frontend and backend together with a single `nuxt-gin` workflow.</p>
        </div>
        <div class="intro-item">
          <h3>Endpoint-First API</h3>
          <p>Define typed endpoints in Go and consume them cleanly from the Nuxt app.</p>
        </div>
        <div class="intro-item">
          <h3>Production Ready</h3>
          <p>Built-in CORS, static serving, and a sensible project layout for growth.</p>
        </div>
        <div class="intro-item">
          <h3>Extensible</h3>
          <p>Drop in models, routes, and UI quickly without fighting the scaffolding.</p>
        </div>
      </div>

      <div class="intro-repo">
        <span>Repository</span>
        <a class="intro-link" href="https://github.com/RapboyGao/nuxt-gin-starter.git" target="_blank" rel="noreferrer">
          https://github.com/RapboyGao/nuxt-gin-starter.git
        </a>
      </div>
    </div>
  </section>
</template>

<style scoped lang="scss">
.intro-card {
  margin-top: 20px;
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

.intro-body {
  margin-top: 16px;
  display: grid;
  gap: 16px;
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
