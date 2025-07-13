const serverConfig = require("./server.config.json");

module.exports = {
  apps: [
    {
      name: "Nuxt-Gin",
      port: serverConfig.ginPort,
      error_file: "logs/error.log",
      out_file: "logs/logs.log",
      combine_logs: true,
      script: "./tmp/production.exe",
      env: {
        NODE_ENV: "production",
        GIN_MODE: "release",
      },
    },
  ],
};
