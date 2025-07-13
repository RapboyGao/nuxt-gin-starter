const Fs = require("fs-extra");
const Zip = require("7zip-min");

/**
 * @param  {...string} names 路径名称
 * @returns 相对于Workspace位置的路径名称
 */
function p(...names) {
  return require("path").join(__dirname, "../", ...names);
}

/** 打包文件夹相对于Workspace位置 */
const DIST_PACKAGE = "dist/package";

/** 需要打包的文件相对位置列表 */
const FILE_LIST = [
  "vue/.output",
  "tmp/production.exe",
  "start.bat",
  "start.ps1",
  "start-pm2.ps1",
  "ecosystem.config.js",
  "server.config.json",
];

/** package.json 文件夹相对于Workspace位置 */
const PACKAGE_JSON = DIST_PACKAGE + "/package.json";
/** 打包package.7z位置 */
const _7Z_PATH = DIST_PACKAGE + "/../package.7z";

/** package.json 内容 */
const PACKAGE_JSON_CONTENT = {
  private: true,
  scripts: {
    "start": "./tmp/production.exe",
    "pm2:start": "pm2 kill && pm2 start && pm2 monit",
    "pm2:kill": "pm2 kill",
  },
  devDependencies: {
    pm2: "latest",
  },
};

/**
 * 把FILE_LIST中的文件拷贝到DIST_PACKAGE的对应位置
 */
function copyFiles() {
  for (const path of FILE_LIST) {
    Fs.copySync(p(path), p(DIST_PACKAGE, path));
  }
}
/**
 * 打包文件为7z格式
 */
async function pack() {
  Fs.removeSync(p(DIST_PACKAGE));
  Fs.removeSync(p(_7Z_PATH));
  copyFiles();
  Fs.outputJSONSync(p(PACKAGE_JSON), PACKAGE_JSON_CONTENT, { spaces: 2 });
  Zip.pack(p(DIST_PACKAGE), p(_7Z_PATH), () => {
    Fs.removeSync(p(DIST_PACKAGE));
  });
}

/**
 * 先generate，让后打包
 */
async function buildAndPack() {
  await require("./builder")();
  await pack();
}

buildAndPack();
