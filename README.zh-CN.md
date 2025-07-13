# [Nuxt 3 Gin starter 🚀](https://github.com/RapboyGao/nuxt3-gin-starter.git)

[![许可证：MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub 星标](https://img.shields.io/github/stars/RapboyGao/nuxt3-gin-starter.svg?style=social)](https://github.com/RapboyGao/nuxt3-gin-starter/stargazers)
[![GitHub 分叉](https://img.shields.io/github/forks/RapboyGao/nuxt3-gin-starter.svg?style=social)](https://github.com/RapboyGao/nuxt3-gin-starter/network)

想要深入了解相关知识，可以查阅 [Nuxt 3 文档](https://nuxt.com/docs/getting-started/introduction) 和 [Gin 框架文档](https://gin-gonic.com)。

## 环境配置 ⚙️

### 1. PowerShell（Win10+）💻

PowerShell 是 Windows 系统上功能强大的命令行 shell 和脚本语言。你可以通过以下方式安装：

- **官方网站**：[Windows 系统安装指南](https://learn.microsoft.com/zh-cn/powershell/scripting/install/installing-powershell-on-windows)
- **GitHub 发布页**：[获取最新版本](https://github.com/PowerShell/PowerShell/releases)
- **镜像地址**：[从镜像下载](https://sourceforge.net/projects/powershell.mirror/files/)
- **官方安装脚本**：
  ```sh
  winget install --id Microsoft.Powershell --source winget
  ```

### 2. Scoop（Win10+）📦

Scoop 是 Windows 系统的命令行安装工具。

- **官方网站**：[Scoop 官网](https://scoop.sh/)
- **安装脚本**：
  ```powershell
  Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # 可选：首次运行远程脚本时需要
  irm get.scoop.sh | iex
  ```

### 3. HomeBrew（MacOS）🍎

HomeBrew 是 macOS 系统上流行的包管理器。使用以下命令安装：

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### 4. Open Api Generator 🔄

#### Windows 系统 🖥️

通过 Scoop 在 Windows 上安装 Open Api Generator：

```powershell
scoop bucket add java
scoop install openjdk
scoop install openapi-generator-cli@7.1.0 # 若不通过 npm 安装
scoop install mingw
```

#### macOS 系统 🍎

##### Java 运行环境

```sh
#!/bin/bash

# 检查是否安装 HomeBrew
if ! command -v brew &> /dev/null; then
    echo "未找到 HomeBrew，正在安装..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# 更新 HomeBrew
brew update

# 安装 OpenJDK
brew install openjdk

# 链接 Java（可能需要管理员权限）
sudo ln -sfn /opt/homebrew/opt/openjdk/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk.jdk

# 验证安装
java -version
```

### 5. Go 语言 🐹

Go 是本项目使用的编程语言。

- **官方网站**：[下载 Go](https://go.dev/dl/)
- **镜像地址**：[从镜像下载](https://studygolang.com/dl)

**安装常用库（需使用 PowerShell）**：

```powershell
$env:GOPRIVATE = "10.10.110.90:8081" # 若需要本地代理
$env:GOPROXY = "https://goproxy.io,direct"
go get github.com/arduino/go-paths-helper
go get github.com/gin-contrib/cors
go get github.com/gin-gonic/gin
go get github.com/golang-module/carbon
go get github.com/mitchellh/mapstructure
go get github.com/xuri/excelize/v2
go get github.com/samber/lo
go get gorm.io/driver/sqlite
go get gorm.io/gorm
```

### 6. pnpm（需先安装 Scoop）📦

pnpm 是一款快速、节省磁盘空间的包管理器。

- **官方网站**：[pnpm 安装指南](https://www.pnpm.cn/installation)

### 7. Node.js 🌐

Node.js 是基于 Chrome V8 引擎的 JavaScript 运行时环境。

- **官方网站**：[下载 Node.js](https://nodejs.org)
- **镜像地址**：[从镜像下载](https://registry.npmmirror.com/binary.html?path=node/v18.13.0/)

### 8. Air ♻️

Air 是 Go 应用的热重载工具。使用以下命令安装：

```sh
go install github.com/cosmtrek/air@latest
```

## 在 Mac 上设置 `GOPROXY` 环境变量 🔧

### 步骤 1：打开终端 🖥️

按下 `Command + Space` 打开 Spotlight 搜索，输入“终端”并按回车即可打开终端。

### 步骤 2：编辑 Shell 配置文件 📝

根据你使用的 Shell 选择对应的配置文件：

#### 如果你使用 Zsh（macOS 默认 Shell）

```bash
nano ~/.zshrc
```

#### 如果你使用 Bash

```bash
nano ~/.bashrc
```

### 步骤 3：添加环境变量 ➕

在打开的文件末尾添加以下内容：

```bash
export GOPROXY="https://goproxy.io,direct"
```

### 步骤 4：保存并关闭文件 💾

- 按下 `Control + X`
- 然后按 `Y` 确认保存
- 最后按 `Enter` 退出编辑器

### 步骤 5：应用配置 🔄

```bash
source ~/.zshrc  # 如果你使用 Zsh
source ~/.bashrc  # 如果你使用 Bash
```

### 步骤 6：验证设置 ✅

```bash
go env GOPROXY
```

如果设置成功，终端会输出：`https://goproxy.io,direct`

### 补充说明 📌

- **同时设置多个代理**：你可以按优先级添加多个代理地址，用逗号分隔。例如：
  ```bash
  export GOPROXY="https://goproxy.io,https://goproxy.cn,direct"
  ```
- **设置 `GOPRIVATE`**：如果你有私有模块，还需要设置 `GOPRIVATE` 来跳过代理。例如：
  ```bash
  export GOPRIVATE="github.com/你的公司/*"
  ```

完成上述配置后，每次启动终端时，Go 都会自动使用指定的代理服务器。

## 创建你自己的项目 🛠️

### 1. API 📄

修改 [gen/openapi.yaml](gen/openapi.yaml)，然后运行 [package.json](package.json) 中的 `api:generate` 脚本。🚀

### 2. 服务器逻辑 💻

在 [server/routes](server/routes/Framework.Handlers.go) 中编写你自己的服务器逻辑，定义服务器如何响应请求。📡

### 3. 模型 📐

#### 1. 在 [server/model](server/model/Framework.Example.Product.go) 中定义自己的模型。📝

#### 2. 在 [Framework.DB.go](server/model/Framework.DB.go) 中注册模型。📚

```go
db.AutoMigrate(&Product{}) // 以及你的其他模型
```

### 4. 前端 🌈

在 [vue](vue/pages/index.vue) 中创建你自己的前端页面。🎨

### 5. 开发 🚧

运行 [package.json](package.json) 中的 `dev` 脚本。🏃‍♂️
