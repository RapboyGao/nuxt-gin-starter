/**
 * 跨域资源共享 (CORS) 中间件
 * 为所有响应添加 CORS 头，允许前端跨域访问 API
 */
import type { H3Event } from "h3";
import { defineEventHandler } from "h3";

export default defineEventHandler((event: H3Event) => {
  // 从事件中提取 Node.js HTTP 请求和响应对象
  const { res, req } = event.node;

  // 定义响应头类型
  type Headers = { [key: string]: string | number | readonly string[] };

  // 定义 CORS 相关的响应头
  const headers: Headers = {
    // 允许携带凭证（如 cookies、HTTP 认证）
    "Access-Control-Allow-Credentials": "true",

    // 预检请求的结果可以缓存的时间（秒）
    "Access-Control-Max-Age": 1728000,

    // 允许的源，使用请求头中的 origin 字段或默认允许所有源
    "Access-Control-Allow-Origin": req.headers.origin || "*",

    // 允许的请求头
    "Access-Control-Allow-Headers": "X-Requested-With,Content-Type",

    // 允许的 HTTP 方法
    "Access-Control-Allow-Methods": "PUT,POST,GET,DELETE,OPTIONS",
  };

  // 遍历并设置所有 CORS 响应头
  for (const key in headers) {
    res.setHeader(key, headers[key]);
  }
});
