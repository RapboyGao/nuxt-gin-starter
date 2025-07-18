const { ensureDirSync, ensureFileSync } = require("fs-extra");
const { join } = require("path");

const cwd = process.cwd();

module.exports = function build() {
  ensureDirSync(join(cwd, "vue/.output"));
  ensureFileSync(join(cwd, "tmp/production.exe"));
  require("concurrently")([
    "go build -o ./tmp/production.exe .",
    "nuxt generate",
  ]);
};
