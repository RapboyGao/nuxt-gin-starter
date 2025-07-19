// [框架] Nuxt Config，不要删除或修改
// 导入 Nuxt 配置类型，用于类型检查，确保配置对象符合 Nuxt 的配置规范
import type { NuxtConfig } from "nuxt/config";
// 导入服务器配置文件，该文件包含了服务器的一些基础配置信息，如端口号、基础 URL 等
import serverConfig from "../server.config.json";
/**
 * 导出默认的 Nuxt 配置对象
 * 该配置对象包含了一些 Nuxt 项目的基本配置信息
 */
export default {
  buildDir: "vue/.nuxt", // 设置构建目录为 "vue/.nuxt"，表示 Nuxt 项目的构建输出将存放在该目录下
  srcDir: "vue", // 设置源代码目录为 "vue"，表示 Nuxt 项目的源代码将存放在该目录下
  // 设置服务器端代码的目录为 "vue/server"，表示服务器端的代码将存放在该目录下
  // serverDir: "vue/server",
  // 禁用服务器端渲染（SSR），即页面将在客户端进行渲染
  ssr: false,
  /**
   * 配置应用的基础 URL
   * 从 server.config.json 文件中获取 baseUrl 作为应用的基础 URL
   * 这个基础 URL 将用于构建应用的路由和请求地址
   */
  app: { baseURL: serverConfig.baseUrl },
  /**
   * 配置开发服务器的端口号
   * 从 server.config.json 文件中获取 nuxtPort 作为开发服务器的端口号
   * 启动开发服务器时，将使用该端口号进行监听
   */
  devServer: {
    port: serverConfig.nuxtPort,
    cors: {
      allowCredentials: true,
      maxAge: "1728000",
      origin: "*", // 允许所有源进行跨域请求
      allowHeaders: ["X-Requested-With", "Content-Type"],
      allowMethods: ["PUT", "POST", "GET", "DELETE", "OPTIONS"],
    },
  },
  /**
   * 配置实验性功能
   * 禁用 payloadExtraction 功能，该功能可能用于提取页面的有效负载数据
   * 这里禁用它可能是为了避免某些兼容性问题或特定的项目需求
   */
  experimental: {
    payloadExtraction: false,
  },
} as NuxtConfig;
