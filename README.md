# [Nuxt Gin Starter 🚀](https://github.com/RapboyGao/nuxt-gin-starter.git)

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/RapboyGao/nuxt-gin-starter.svg?style=social)](https://github.com/yourusername/nuxt-gin-starter/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/RapboyGao/nuxt-gin-starter.svg?style=social)](https://github.com/yourusername/nuxt-gin-starter/network)

[🇨🇳简体中文](./README.zh-CN.md)

Explore the following to gain in-depth knowledge:

- [Nuxt 4 documentation](https://nuxt.com/docs/getting-started/introduction)
- [Gin](https://gin-gonic.com)
- [GORM](https://gorm.io)
- [Vue](https://vuejs.org)
- [OpenAPI](https://www.openapis.org)
- [OpenAPI Generator](https://openapi-generator.tech)

Programming Languages:

- [Typescript](https://www.typescriptlang.org)
- [Go](https://go.dev)

## Recommended IDE - [VS Code](https://code.visualstudio.com)

## Environment Setup ⚙️

### 1. Powershell (Win10+) 💻

Powershell is a powerful command-line shell and scripting language on Windows. You can install it through the following methods:

- **Official Website**: [Learn about installation on Windows](https://learn.microsoft.com/zh-cn/powershell/scripting/install/installing-powershell-on-windows)
- **Github Releases**: [Get the latest release on GitHub](https://github.com/PowerShell/PowerShell/releases)
- **Mirror**: [Download from the mirror](https://sourceforge.net/projects/powershell.mirror/files/)
- **Official Installation Script**:

```sh
winget install --id Microsoft.Powershell --source winget
```

### 2. Scoop (Win10+) 📦

Scoop is a command-line installer for Windows.

- **Official Website**: [Scoop official site](https://scoop.sh/)
- **Installation Script**:
  ```powershell
  Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Required to run a remote script for the first time
  irm get.scoop.sh | iex
  ```

### 3. HomeBrew (MacOS) 🍎

HomeBrew is a popular package manager for macOS. Install it using the following command:

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### 4. Open Api Generator 🔄

#### Windows 🖥️

Install the Open Api Generator on Windows via Scoop:

```powershell
scoop bucket add java
scoop install openjdk
scoop install openapi-generator-cli@7.1.0 # If not installing via npm
scoop install mingw
```

#### macOS 🍎

##### Java Runtime

```sh
#!/bin/bash

# Check if Homebrew is installed
if ! command -v brew &> /dev/null; then
    echo "Homebrew not found, installing..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# Update Homebrew
brew update

# Install OpenJDK
brew install openjdk

# Link Java (may require administrator privileges)
sudo ln -sfn /opt/homebrew/opt/openjdk/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk.jdk

# Verify the installation
java -version
```

### 5. Go Language 🐹

Go is a programming language used in this project.

- **Official Website**: [Download Go](https://go.dev/dl/)
- **Mirror**: [Download from the mirror](https://studygolang.com/dl)

**Install Common Libraries (Requires Powershell)**:

```powershell
$env:GOPRIVATE = "10.10.110.90:8081" # If a local proxy is needed
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

### 6. pnpm (Requires Scoop) 📦

pnpm is a fast, disk-space efficient package manager.

- **Official Website**: [pnpm installation guide](https://www.pnpm.cn/installation)

### 7. Nodejs 🌐

Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine.

- **Official Website**: [Download Node.js](https://nodejs.org)
- **Mirror**: [Download from the mirror](https://registry.npmmirror.com/binary.html?path=node/v18.13.0/)

### 8. Air ♻️

Air is a live reload utility for Go applications. Install it with:

```sh
go install github.com/cosmtrek/air@latest
```

## Set the `GOPROXY` Environment Variable on Mac 🔧

### Step 1: Open the Terminal 🖥️

You can open the terminal by pressing `Command + Space` to open the Spotlight Search, then typing "Terminal" and hitting Enter.

### Step 2: Edit the Shell Configuration File 📝

Choose the appropriate configuration file based on your shell:

#### If you are using Zsh (the default shell on macOS)

```bash
nano ~/.zshrc
```

#### If you are using Bash

```bash
nano ~/.bashrc
```

### Step 3: Add the Environment Variable ➕

At the end of the opened file, add the following line:

```bash
export GOPROXY="https://goproxy.io,direct"
```

### Step 4: Save and Close the File 💾

- Press `Control + X`.
- Then press `Y` to confirm the save.
- Finally, press `Enter` to exit the editor.

### Step 5: Apply the Configuration 🔄

```bash
source ~/.zshrc  # If you are using Zsh
source ~/.bashrc  # If you are using Bash
```

### Step 6: Verify the Settings ✅

```bash
go env GOPROXY
```

If the setting is successful, the terminal will output: `https://goproxy.io,direct`

### Supplementary Notes 📌

- **Set Multiple Proxies Simultaneously**: You can add multiple proxy addresses in order of priority, separated by commas. For example:
  ```bash
  export GOPROXY="https://goproxy.io,https://goproxy.cn,direct"
  ```
- **Set `GOPRIVATE`**: If you have private modules, you also need to set `GOPRIVATE` to skip the proxy. For example:
  ```bash
  export GOPRIVATE="github.com/your-company/*"
  ```

After completing the above configuration, Go will automatically use the specified proxy server every time you start the terminal.

## File Structure

```plaintext
nuxt-gin-starter/
├── .air.toml                     # Air tool configuration (Go hot reloading)
├── .gitignore                    # Git version control ignore rules
├── .npmrc                        # npm/pnpm configuration
├── .openapi-generator-ignore     # OpenAPI generator ignore rules
├── Dockerfile                    # Containerization deployment configuration
├── LICENSE                       # Open source license (MIT)
├── README.md                     # English project description
├── README.zh-CN.md               # Chinese project description
├── ecosystem.config.js           # PM2 process management configuration
├── go.mod                        # Go module dependency management
├── main.go                       # Go server entry point
├── nuxt.config.ts                # Nuxt.js core configuration (consistent with /config)
├── openapi.yaml                  # OpenAPI specification document (editable)
├── openapitools.json             # OpenAPI generator configuration
├── package.json                  # Node.js project configuration
├── server.config.json            # Server configuration (ports, etc.)
├── tsconfig.json                 # TypeScript compilation configuration
│
├── vue/                          # Nuxt.js frontend code (editable)
│   ├── composables/              # Vue global reusable code
│   │   ├── api/                  # Content generated by OpenAPI generator
│   └── pages/                    # Page components
│
├── server/                       # Gin backend code
│   ├── frontend/                 # Frontend service related
│   │   └── Framework.GetConfig.go # Configuration file loading
│   ├── model/                    # Database models (editable)
│   │   ├── Example.Product.go    # Example product model
│   │   └── Framework.DB.go       # Database initialization
│   ├── routes/                   # API route definitions (editable)
│   └── utils/                    # Utility functions
│       ├── Framework.Directory.go # Directory operation tools
│       ├── Framework.Excelize.go  # Excel processing tools
│       ├── Framework.MapStructure.go # Data structure conversion
│       └── Framework.Percentage.go # Percentage calculation tools
│
├── config/                       # Nuxt project configuration
│   ├── [Framework]misc.ts        # Nuxt miscellaneous configuration
│   ├── [Framework]nitro.ts       # Nitro engine configuration
│   ├── [Framework]vite.ts        # Vite build configuration
│   └── index.ts                  # Configuration entry
│
└── .vscode/                      # VSCode development configuration
```

## Create your own project 🛠️

### 1. API 📄

Modify [openapi.yaml](openapi.yaml) then run `api:generate` script in [package.json](package.json). 🚀

- API generated contents are located at
  - [server/api](server/api/api_default.go)
  - [vue/composables/api](vue/composables/api/index.ts)

### 2. Server Logics 💻

Write your own server logics in [server/routes](server/routes/Framework.Handlers.go) about how the server responses. 📡

### 3. Models 📐

#### 1. Define your own models in [server/model](server/model/Example.Product.go). 📝

#### 2. Register the models in [Framework.DB.go](server/model/Framework.DB.go). 📚

```go
db.AutoMigrate(&Product{}) // and your own models
```

### 4. Frontend 🌈

Create your own frontend pages in [vue](vue/pages/index.vue). 🎨

### 5. Develop 🚧

Run `dev` script in [package.json](package.json). 🏃‍♂️
