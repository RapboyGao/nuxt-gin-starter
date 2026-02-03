<template>
  <div>
    <div v-if="date" class="bg-black">
      If you're seeing this, you are receiving data from Gin server:
      {{ date }}
    </div>
    <div v-if="product" class="bg-black">
      Product from getProductAPI: {{ product.name }} - ${{ product.price }}
    </div>
    <div v-if="productError" class="bg-black">
      {{ productError }}
    </div>
    <NuxtWelcome />
  </div>
</template>

<script setup lang="ts">
import { getProduct, testPost } from '@/composables/auto-generated-api';
import type { ProductResponseBody, TestResponseBody } from '@/composables/auto-generated-api';

const response = ref<TestResponseBody | undefined>();
const product = ref<ProductResponseBody | undefined>();
const productError = ref('');

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

const loadProduct = async () => {
  try {
    productError.value = '';
    product.value = await getProduct({
      path: { ID: 'p_1001' },
      query: { WithStock: true },
    });
  } catch (error) {
    console.error(error);
    productError.value = 'load product failed';
  }
};

onMounted(() => {
  loadProduct();
});

setInterval(async () => {
  response.value = await testPost();
  console.log(response.value);
}, 1000);
</script>

<style lang="scss">
.bg-black {
  background-color: black;
  color: #cccccc;
}
</style>
