<template>
  <div>
    <div v-if="date" class="bg-black">
      If you're seeing this, you are receiving data from Gin server:
      {{ date }}
    </div>
    <NuxtWelcome />
  </div>
</template>

<script setup lang="ts">
import type { TestResponse } from '@/composables/api';
const response = ref<TestResponse | undefined>();

const date = computed(() => {
  const time = response.value?.time;
  if (time) {
    return new Date(time);
  } else {
    return undefined;
  }
});

useSeoMeta({
  title: 'Hello World',
  applicationName: 'Hello World',
});

setInterval(async () => {
  response.value = (await myApi.testPost()).data;
}, 1000);
</script>

<style lang="scss">
.bg-black {
  background-color: black;
  color: #cccccc;
}
</style>
