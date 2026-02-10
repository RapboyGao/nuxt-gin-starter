<template>
  <nav class="nav-top" aria-label="Main navigation">
    <NuxtLink
      v-for="item in items"
      :key="item.to"
      :to="item.to"
      class="nav-link"
      :class="{ active: route.path === item.to }"
    >
      {{ item.title }}
    </NuxtLink>
  </nav>
</template>

<script setup lang="ts">
const route = useRoute();

const items: ReadonlyArray<{ to: string; title: string }> = [
  {
    to: '/',
    title: 'Overview',
  },
  {
    to: '/playground',
    title: 'Playground',
  },
];
</script>

<style scoped lang="scss">
.nav-top {
  width: 100%;
  height: 100px;
  margin: 0;
  padding: 0 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  border-radius: 0;
  background: rgba(17, 23, 32, 0.92);
  border-bottom: 1px solid rgba(255, 255, 255, 0.12);
  backdrop-filter: blur(10px);
}

.nav-link {
  padding: 8px 14px;
  border-radius: 999px;
  color: #d5e1f3;
  text-decoration: none;
  border: 1px solid transparent;
  font-size: 13px;
  font-weight: 600;
  transition: transform 180ms ease, border-color 180ms ease, background 180ms ease;
}

.nav-link:hover {
  border-color: rgba(126, 194, 255, 0.45);
  background: rgba(126, 194, 255, 0.1);
  transform: translateY(-1px);
}

.nav-link.active {
  color: #f8fcff;
  border-color: rgba(159, 236, 204, 0.75);
  background: linear-gradient(135deg, rgba(129, 219, 255, 0.18), rgba(159, 236, 204, 0.14));
}

@media (max-width: 480px) {
  .nav-top {
    justify-content: center;
    gap: 8px;
    padding: 0 8px;
  }

  .nav-link {
    min-width: 120px;
    padding: 8px 10px;
  }
}
</style>
