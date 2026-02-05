<template>
  <div class="page">
    <ProjectIntro />
    <ProductCrud />
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

</style>
