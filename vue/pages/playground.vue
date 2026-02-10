<template>
  <div class="page">
    <NavigationMenu />
    <ProductCrud />
    <WebSocketDemo />
  </div>
</template>

<script setup lang="ts">
useSeoMeta({
  title: 'Playground - Nuxt Gin Starter',
  applicationName: 'Nuxt Gin Starter',
});
</script>

<style scoped lang="scss">
.page {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  min-height: 100vh;
  background: #0b0c0e;
}
</style>
