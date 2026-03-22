const fs = require('fs');
const path = require('path');
const serverConfig = require('./server.config.json');

const rootDir = __dirname;
const packagedBinary = path.resolve(rootDir, '.build/production/server/server-production');
const localBuildBinary = path.resolve(rootDir, '.build/.server/production');
const binaryPath = fs.existsSync(packagedBinary) ? packagedBinary : localBuildBinary;

module.exports = {
  apps: [
    {
      name: 'nuxt-gin-starter',
      cwd: rootDir,
      script: binaryPath,
      interpreter: 'none',
      exec_mode: 'fork',
      instances: 1,
      watch: false,
      merge_logs: true,
      error_file: path.resolve(rootDir, 'logs/error.log'),
      out_file: path.resolve(rootDir, 'logs/out.log'),
      max_memory_restart: '512M',
      env: {
        NODE_ENV: 'production',
        PORT: String(serverConfig.ginPort),
        GIN_MODE: 'release',
        NUXT_GIN_MODE: 'release',
        NUXT_GIN_EXPORT_TS_ON_RUN: 'false',
      },
    },
  ],
};
