import { createDefaultConfig } from 'nuxt-gin-tools/src/nuxt-config';
import { defineNuxtConfig } from 'nuxt/config';
import SERVER_CONFIG from './server.config.json';
const config = createDefaultConfig({
  serverConfig: SERVER_CONFIG,
});

export default defineNuxtConfig({
  ...config,
});
