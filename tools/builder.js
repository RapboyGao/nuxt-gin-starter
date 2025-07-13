const { ensureDirSync, ensureFileSync } = require("fs-extra");
const { join } = require("path");

module.exports = function build() {
  ensureDirSync(join(__dirname, "../vue/.output"));
  ensureFileSync(join(__dirname, "../tmp/production.exe"));
  require("concurrently")([
    "go build -o ./tmp/production.exe .",
    "nuxt generate",
  ]);
};
