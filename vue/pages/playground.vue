<template>
  <div class="page">
    <NavigationMenu />
    <ProductCRUDDemo />
    <WebSocketDemo />
  </div>
</template>

<script setup lang="ts">
import ProductCRUDDemo from '@/components/ProductCRUDDemo.vue';

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
