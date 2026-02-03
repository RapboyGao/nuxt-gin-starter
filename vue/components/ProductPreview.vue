<template>
  <section class="status-card">
    <div class="status-header">
      <span class="status-dot pulse" />
      <h3>Product Preview</h3>
    </div>
    <div v-if="product" class="status-body">
      <p class="status-text">Product from getProductAPI</p>
      <div class="product-row">
        <span class="product-name">{{ product.name }}</span>
        <span class="product-price">${{ product.price }}</span>
      </div>
    </div>
    <div v-else-if="error" class="status-body">
      <p class="status-text error">{{ error }}</p>
    </div>
    <div v-else class="status-body">
      <p class="status-text muted">Loading product...</p>
    </div>
  </section>
</template>

<script setup lang="ts">
import { getProduct } from '@/composables/auto-generated-api';
import type { ProductResponseBody } from '@/composables/auto-generated-api';

const product = ref<ProductResponseBody | undefined>();
const error = ref('');

const loadProduct = async () => {
  try {
    error.value = '';
    product.value = await getProduct({
      path: { ID: 'p_1001' },
      query: { WithStock: true },
    });
  } catch (err) {
    console.error(err);
    error.value = 'load product failed';
  }
};

onMounted(() => {
  loadProduct();
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
  background: #60a5fa;
  box-shadow: 0 0 8px rgba(96, 165, 250, 0.6);
}

.status-dot.pulse {
  animation: pulse 1.6s ease-in-out infinite;
}

.status-body {
  display: grid;
  gap: 6px;
}

.status-text {
  margin: 0;
  font-size: 13px;
  color: #9aa3af;
}

.status-text.muted {
  color: #6b7280;
}

.status-text.error {
  color: #fca5a5;
}

.product-row {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 12px;
}

.product-name {
  font-weight: 600;
  color: #e5e7eb;
}

.product-price {
  font-weight: 600;
  color: #93c5fd;
}

@keyframes pulse {
  0% {
    box-shadow: 0 0 0 rgba(96, 165, 250, 0.6);
  }
  70% {
    box-shadow: 0 0 10px rgba(96, 165, 250, 0);
  }
  100% {
    box-shadow: 0 0 0 rgba(96, 165, 250, 0);
  }
}
</style>
