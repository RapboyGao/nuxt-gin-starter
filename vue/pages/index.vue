<template>
  <div class="page">
    <ProjectIntro />
    <div class="grid">
      <ServerTime />
      <ProductPreview />
    </div>
  </div>
</template>

<script setup lang="ts">
useSeoMeta({
  title: 'Hello World',
  applicationName: 'Hello World',
});
</script>

<style lang="scss">
.page {
  display: grid;
  background: #0b0c0e;
  min-height: 100vh;
  min-width: 100vw;
}

.grid {
  display: grid;
  gap: 16px;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  padding-top: 24px;
}
</style>
