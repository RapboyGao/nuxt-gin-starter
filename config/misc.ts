// [框架] Nuxt Config，不要删除或修改
import type { NuxtConfig } from "nuxt/config";
import serverConfig from "../server.config.json";

export default {
  rootDir: "vue",
  ssr: false,
  app: { baseURL: serverConfig.baseUrl },
  devServer: { port: serverConfig.nuxtPort },
  experimental: {
    payloadExtraction: false,
  },
} as NuxtConfig;
