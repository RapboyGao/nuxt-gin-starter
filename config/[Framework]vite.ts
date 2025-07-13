// [框架] Nuxt Config，不要删除或修改
// 导入 Nuxt 配置类型，用于类型检查，确保配置对象符合 Nuxt 的配置规范
import type { NuxtConfig } from "nuxt/config";
// 导入 Vite 的 legacy 插件，该插件用于为旧版浏览器提供兼容性支持
import legacy from "@vitejs/plugin-legacy";

/**
 * 导出 Vite 配置类型，它是 Nuxt 配置中 vite 部分的类型
 * 这样可以方便在其他地方使用和检查 Vite 配置的类型
 */
export type ViteConfig = NuxtConfig["vite"];

/**
 * 导出默认的 Vite 配置对象
 * 该配置对象包含了一些 Vite 构建工具的配置信息
 */
export default {
  vite: {
    // 配置 Vite 插件
    // 使用 legacy 插件为旧版浏览器提供兼容性支持
    // 这里指定目标浏览器为 "defaults"（默认浏览器列表）和 "not IE 11"（排除 IE 11）
    plugins: [legacy({ targets: ["defaults", "not IE 11"] })],
    // 配置 esbuild 编译器的选项
    // 设置 jsxFactory 为 "h"，这是 Vue 3 中用于创建虚拟 DOM 的函数
    // 设置 jsxFragment 为 "Fragment"，用于处理 JSX 中的片段
    esbuild: { jsxFactory: "h", jsxFragment: "Fragment" },
  },
} as ViteConfig;
