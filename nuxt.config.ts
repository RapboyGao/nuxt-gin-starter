import { createDefaultConfig } from 'nuxt-gin-tools/src/nuxt-config';
import { defineNuxtConfig } from 'nuxt/config';
import SERVER_CONFIG from './server.config.json';
import { BASE_PATH } from './vue/composables/api/base';

const config = createDefaultConfig({
  apiBasePath: BASE_PATH,
  serverConfig: SERVER_CONFIG,
});

export default defineNuxtConfig({
  ...config,
});
