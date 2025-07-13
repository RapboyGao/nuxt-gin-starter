// [框架] Nuxt Config，不要删除或修改
import type { NuxtConfig } from "nuxt/config";
import serverConfig from "../server.config.json";
import { BASE_PATH } from "../vue/composables/api/base";

export type NitroConfig = NuxtConfig["nitro"];

const thisBasePath = BASE_PATH.replace(/^https?:[/]{2}[^/]+/, "");
const target = `http://localhost:${serverConfig.ginPort}${thisBasePath}`;

export default {
  devProxy: {
    [thisBasePath]: {
      target: target,
      changeOrigin: true,
    },
  },
} as NitroConfig;
