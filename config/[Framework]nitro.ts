// [框架] Nuxt Config，不要删除或修改
// 导入 Nuxt 配置类型，用于类型检查，确保配置对象符合 Nuxt 的配置规范
import type { NuxtConfig } from "nuxt/config";
// 导入服务器配置文件，该文件包含了服务器的一些基础配置信息，如端口号、基础 URL 等
import serverConfig from "../server.config.json";
// 导入 API 的基础路径，该路径用于构建 API 请求的 URL
import { BASE_PATH } from "../vue/composables/api/base";

/**
 * 导出 Nitro 配置类型，它是 Nuxt 配置中 nitro 部分的类型
 * 这样可以方便在其他地方使用和检查 Nitro 配置的类型
 */
export type NitroConfig = NuxtConfig["nitro"];

/**
 * 处理基础路径，去除协议和域名部分，只保留路径部分
 * 例如，将 "https://example.com/api-go" 转换为 "/api-go"
 */
const thisBasePath = BASE_PATH.replace(/^https?:[/]{2}[^/]+/, "");

/**
 * 定义代理目标 URL
 * 拼接本地服务器的地址和处理后的基础路径，用于开发环境的代理
 * 例如，当 Gin 服务器端口为 8099 时，目标 URL 可能为 "http://localhost:8099/api-go"
 */
const target = `http://localhost:${serverConfig.ginPort}${thisBasePath}`;

/**
 * 导出默认的 Nitro 配置对象
 * 该配置对象包含了开发环境下的代理配置，用于将特定路径的请求代理到目标服务器
 */
export default {
  devProxy: {
    // 定义代理规则，将匹配 thisBasePath 的请求代理到目标服务器
    [thisBasePath]: {
      // 目标服务器的 URL
      target: target,
      // 是否改变请求的源，设置为 true 可以避免跨域问题
      changeOrigin: true,
    },
  },
} as NitroConfig;
