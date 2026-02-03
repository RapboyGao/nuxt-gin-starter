<template>
  <section class="status-card">
    <div class="status-header">
      <span class="status-dot online" />
      <h3>Server Time</h3>
    </div>
    <div v-if="date" class="status-body">
      <p class="status-text">Live data from Gin server:</p>
      <p class="status-value">{{ date }}</p>
    </div>
    <div v-else class="status-body">
      <p class="status-text muted">Waiting for server response...</p>
    </div>
  </section>
</template>

<script setup lang="ts">
import { testPost } from '@/composables/auto-generated-api';
import type { TestResponseBody } from '@/composables/auto-generated-api';

const response = ref<TestResponseBody | undefined>();

const date = computed(() => {
  const time = response.value?.time;
  if (time) {
    return new Date(time);
  }
  return undefined;
});

const loadTime = async () => {
  response.value = await testPost();
};

onMounted(() => {
  loadTime();
  const timer = setInterval(loadTime, 1000);
  onBeforeUnmount(() => {
    clearInterval(timer);
  });
});
</script>

<style scoped lang="scss">
.status-card {
  padding: 14px 16px;
  border-radius: 14px;
  background: #111214;
  border: 1px solid rgba(255, 255, 255, 0.08);
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.2);
  color: #e8e8e8;
}

.status-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 10px;
}

.status-header h3 {
  margin: 0;
  font-size: 16px;
}

.status-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #4ade80;
  box-shadow: 0 0 8px rgba(74, 222, 128, 0.6);
}

.status-body {
  display: grid;
  gap: 4px;
}

.status-text {
  margin: 0;
  font-size: 13px;
  color: #9aa3af;
}

.status-text.muted {
  color: #6b7280;
}

.status-value {
  margin: 0;
  font-size: 14px;
  font-weight: 600;
  color: #dbeafe;
}
</style>
