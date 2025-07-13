// [框架] Nuxt Config，不要删除或修改
import type { NuxtConfig } from "nuxt/config";
import legacy from "@vitejs/plugin-legacy";

export type ViteConfig = NuxtConfig["vite"];

export default {
  vite: {
    plugins: [ legacy({ targets: ["defaults", "not IE 11"] })],
    esbuild: { jsxFactory: "h", jsxFragment: "Fragment" },
  },
} as ViteConfig;
