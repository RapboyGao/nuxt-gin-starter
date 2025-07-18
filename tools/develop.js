const os = require("os");
const { readJSONSync } = require("fs-extra");
const { join } = require("path");
const cwd = process.cwd();
const serverConfig = readJSONSync(join(cwd, "server.config.json"));

// 如果是macOS
if (os.platform() === "darwin") {
  // 如果是macOS，使用open命令打开浏览器
  require("concurrently")([`nuxt dev --port=${serverConfig.nuxtPort}`, "~/go/bin/air"]);
} else {
  require("concurrently")([`nuxt dev --port=${serverConfig.nuxtPort}`, "air"]);
}
