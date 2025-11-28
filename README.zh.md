[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/astkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/astkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/astkratos)](https://pkg.go.dev/github.com/orzkratos/astkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/astkratos/main.svg)](https://coveralls.io/github/orzkratos/astkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/astkratos.svg)](https://github.com/orzkratos/astkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/astkratos)](https://goreportcard.com/report/github.com/orzkratos/astkratos)

# astkratos

基于 Go AST 的高级 Kratos 项目代码结构分析引擎。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🔍 **gRPC 检测引擎**: 智能识别和提取 gRPC 客户端、服务器和服务  
📊 **结构体分析**: 全面的 Go 结构体解析，包含详细的 AST 信息和源码映射  
📁 **智能文件扫描**: 基于模式的文件系统导航，支持可定制的后缀匹配  
📦 **模块智能分析**: 高级 go.mod 解析，包含依赖分析和工具链版本解析  
🎯 **代码生成支持**: 专为自动化代码生成和项目结构分析工作流程设计

## 架构

```
┌─────────────────────────────────────────────────────────────────┐
│  gRPC Detection Engine                                          │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │  WalkFiles("./api", "_grpc.pb.go")                        │  │
│  │  ┌─────────────────────────────────────────────────────┐  │  │
│  │  │  Pattern Matching:                                  │  │  │
│  │  │    - "type *Client interface {"  → Client           │  │  │
│  │  │    - "type *Server interface {"  → Server           │  │  │
│  │  │    - "type Unimplemented*Server" → Stub             │  │  │
│  │  └─────────────────────────────────────────────────────┘  │  │
│  │  Returns: []*GrpcTypeDefinition { Name, Package, Path }   │  │
│  └───────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

## 安装

```bash
go get github.com/orzkratos/astkratos
```

## 使用方法

### 项目分析

全面的项目分析 - 合并模块信息和 gRPC 组件检测：

```go
package main

import (
	"os"
	"path/filepath"

	"github.com/orzkratos/astkratos"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	tempDIR := rese.V1(os.MkdirTemp("", "go-kratos-test-*"))
	defer func() {
		must.Done(os.RemoveAll(tempDIR))
	}()

	output := rese.A1(osexec.NewExecConfig().WithPath(tempDIR).Exec("kratos", "new", "demo1kratos"))
	zaplog.SUG.Debugln(string(output))

	projectRoot := osmustexist.ROOT(filepath.Join(tempDIR, "demo1kratos"))
	report := astkratos.AnalyzeProject(projectRoot)
	zaplog.SUG.Debugln(neatjsons.S(report))
}
```

⬆️ **源码:** [Source](internal/demos/demo1x/main.go)

### gRPC 检测

扫描 `_grpc.pb.go` 文件，提取 gRPC 类型定义（客户端、服务器、服务）：

```go
package main

import (
	"os"
	"path/filepath"

	"github.com/orzkratos/astkratos"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	tempDIR := rese.V1(os.MkdirTemp("", "go-kratos-test-*"))
	defer func() {
		must.Done(os.RemoveAll(tempDIR))
	}()

	output := rese.A1(osexec.NewExecConfig().WithPath(tempDIR).Exec("kratos", "new", "demo2kratos"))
	zaplog.SUG.Debugln(string(output))

	projectRoot := osmustexist.ROOT(filepath.Join(tempDIR, "demo2kratos"))
	apiPath := osmustexist.ROOT(filepath.Join(projectRoot, "api"))

	clients := astkratos.ListGrpcClients(apiPath)
	zaplog.SUG.Debugln("gRPC Clients:")
	zaplog.SUG.Debugln(neatjsons.S(clients))

	servers := astkratos.ListGrpcServers(apiPath)
	zaplog.SUG.Debugln("gRPC Servers:")
	zaplog.SUG.Debugln(neatjsons.S(servers))

	services := astkratos.ListGrpcServices(apiPath)
	zaplog.SUG.Debugln("gRPC Services:")
	zaplog.SUG.Debugln(neatjsons.S(services))
}
```

⬆️ **源码:** [Source](internal/demos/demo2x/main.go)

### 模块信息

提取 go.mod 元数据，包括依赖和工具链版本：

```go
package main

import (
	"os"
	"path/filepath"

	"github.com/orzkratos/astkratos"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	tempDIR := rese.V1(os.MkdirTemp("", "go-kratos-test-*"))
	defer func() {
		must.Done(os.RemoveAll(tempDIR))
	}()

	output := rese.A1(osexec.NewExecConfig().WithPath(tempDIR).Exec("kratos", "new", "demo3kratos"))
	zaplog.SUG.Debugln(string(output))

	projectRoot := osmustexist.ROOT(filepath.Join(tempDIR, "demo3kratos"))

	moduleInfo := rese.P1(astkratos.GetModuleInfo(projectRoot))
	zaplog.SUG.Debugln("Module Info:")
	zaplog.SUG.Debugln(neatjsons.S(moduleInfo))

	zaplog.SUG.Debugln("Toolchain Version:", moduleInfo.GetToolchainVersion())

	apiPath := osmustexist.ROOT(filepath.Join(projectRoot, "api"))
	zaplog.SUG.Debugln("Has gRPC Clients:", astkratos.HasGrpcClients(apiPath))
	zaplog.SUG.Debugln("Has gRPC Servers:", astkratos.HasGrpcServers(apiPath))
	zaplog.SUG.Debugln("gRPC Service Count:", astkratos.CountGrpcServices(apiPath))
}
```

⬆️ **源码:** [Source](internal/demos/demo3x/main.go)

## API 参考

### 核心类型

- **`GrpcTypeDefinition`**: 表示包含包和名称信息的 gRPC 类型定义
- **`StructDefinition`**: 完整的结构体分析，包含 AST 类型、源码和代码片段
- **`ModuleInfo`**: 全面的 Go 模块元数据，包括依赖和工具链信息
- **`ProjectReport`**: 包含聚合结果的全面项目分析报告

### 主要函数

- **`ListGrpcClients(root string)`**: 从项目中提取所有 gRPC 客户端接口
- **`ListGrpcServers(root string)`**: 检测 gRPC 服务器接口
- **`ListGrpcServices(root string)`**: 检测可用的 gRPC 服务
- **`ListGrpcUnimplementedServers(root string)`**: 查找未实现的服务器结构
- **`GetStructsMap(path string)`**: 解析和分析特定文件中的 Go 结构体
- **`GetModuleInfo(projectPath string)`**: 提取全面的模块和依赖信息

### 便利函数

- **`HasGrpcClients(root string)`**: 检查是否存在 gRPC 客户端
- **`HasGrpcServers(root string)`**: 检查是否存在 gRPC 服务器
- **`CountGrpcServices(root string)`**: 获取 gRPC 服务的数量
- **`AnalyzeProject(projectRoot string)`**: 包含聚合结果的全面项目分析

### 调试函数

- **`SetDebugMode(enable bool)`**: 启用或禁用调试输出，用于开发和故障排查
- **`IsDebugMode()`**: 返回当前调试模式状态

## 示例

### 调试模式

**启用调试输出查看详细分析过程：**
```go
astkratos.SetDebugMode(true)
report := astkratos.AnalyzeProject(".")
astkratos.SetDebugMode(false)
```

**检查当前调试状态：**
```go
if astkratos.IsDebugMode() {
    fmt.Println("调试模式已启用")
}
```

### 结构体分析

**解析 Go 结构体并获取 AST 信息：**
```go
structs := astkratos.GetStructsMap("internal/biz/account.go")
for name, def := range structs {
    fmt.Printf("结构体: %s\n", name)
    fmt.Printf("代码: %s\n", def.StructCode)
}
```

**访问 AST 类型和源码：**
```go
structs := astkratos.GetStructsMap("account.go")
accountDef := structs["Account"]
fmt.Printf("字段数量: %d\n", len(accountDef.Type.Fields.List))
```

### 未实现存根检测

**查找 proto 生成的未实现存根：**
```go
unimplemented := astkratos.ListGrpcUnimplementedServers("./api")
for _, stub := range unimplemented {
    fmt.Printf("存根: %s 在 %s\n", stub.Name, stub.Package)
}
```

### 模式匹配文件扫描

**按后缀模式扫描文件：**
```go
pattern := utils.NewSuffixPattern([]string{"_grpc.pb.go"})
utils.WalkFiles("./api", pattern, func(path string, info os.FileInfo) error {
    fmt.Printf("发现: %s\n", path)
    return nil
})
```

## 使用场景

**🛠 代码生成工具**: 基于 proto 定义生成服务实现
**📈 项目分析**: 分析项目架构并生成文档
**🔧 重构工具**: 理解代码依赖关系并协助重构
**🚀 CI/CD 集成**: 在构建管道中验证项目结构
**📋 架构文档**: 生成项目结构图

## 演示项目

完整的可运行演示（包含 Kratos 项目设置）：[internal/demos](internal/demos)

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/astkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/astkratos)