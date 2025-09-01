# astkratos

基于 Go AST 的高级 Kratos 项目代码结构分析引擎。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🔍 **gRPC 发现引擎**: 智能识别和提取 gRPC 客户端、服务器和服务  
📊 **结构体分析**: 全面的 Go 结构体解析，包含详细的 AST 信息和源码映射  
📁 **智能文件遍历**: 基于模式的文件系统导航，支持可定制的后缀匹配  
📦 **模块智能分析**: 高级 go.mod 解析，包含依赖分析和工具链版本解析  
🎯 **代码生成就绪**: 专为自动化代码生成和项目结构分析工作流程设计

## 安装

```bash
go install github.com/orzkratos/astkratos@latest
```

## 使用方法

### 基本 gRPC 分析

```go
package main

import (
    "fmt"
    "github.com/orzkratos/astkratos"
)

func main() {
    // 列出项目中的所有 gRPC 客户端
    clients := astkratos.ListGrpcClients("./api")
    for _, client := range clients {
        fmt.Printf("客户端: %s.%s\n", client.Package, client.Name)
    }
    
    // 列出所有 gRPC 服务器
    servers := astkratos.ListGrpcServers("./api") 
    for _, server := range servers {
        fmt.Printf("服务器: %s.%s\n", server.Package, server.Name)
    }
    
    // 提取 gRPC 服务
    services := astkratos.ListGrpcServices("./api")
    for _, service := range services {
        fmt.Printf("服务: %s.%s\n", service.Package, service.Name)
    }
}
```

### 结构体分析

```go
// 分析 Go 文件中的所有结构体
structs := astkratos.ListStructsMap("internal/biz/user.go")
for name, def := range structs {
    fmt.Printf("结构体: %s\n", name)
    fmt.Printf("源码: %s\n", def.StructCode)
}
```

### 模块信息

```go
// 获取全面的模块信息
moduleInfo, err := astkratos.GetModuleInfo(".")
if err == nil {
    fmt.Printf("模块: %s\n", moduleInfo.Module.Path)
    fmt.Printf("Go 版本: %s\n", moduleInfo.GetToolchainVersion())
    
    // 列出依赖
    for _, req := range moduleInfo.Require {
        status := "直接依赖"
        if req.Indirect {
            status = "间接依赖"
        }
        fmt.Printf("依赖: %s@%s (%s)\n", req.Path, req.Version, status)
    }
}
```

### 自定义文件遍历

```go
// 为特定文件类型创建自定义后缀匹配器
matcher := astkratos.NewSuffixMatcher([]string{".proto", "_grpc.pb.go"})

// 使用自定义处理逻辑遍历文件
err := astkratos.WalkFiles("./api", matcher, func(path string, info os.FileInfo) error {
    fmt.Printf("处理中: %s\n", path)
    // 您的自定义文件处理逻辑
    return nil
})
```

## API 参考

### 核心类型

- **`GrpcTypeDefinition`**: 表示包含包和名称信息的 gRPC 类型定义
- **`StructDefinition`**: 完整的结构体分析，包含 AST 类型、源码和代码片段  
- **`ModuleInfo`**: 全面的 Go 模块元数据，包括依赖和工具链信息
- **`SuffixMatcher`**: 基于模式匹配的智能文件过滤功能

### 主要函数

- **`ListGrpcClients(root string)`**: 从项目中提取所有 gRPC 客户端接口
- **`ListGrpcServers(root string)`**: 识别所有 gRPC 服务器接口  
- **`ListGrpcServices(root string)`**: 发现可用的 gRPC 服务
- **`ListGrpcUnimplementedServers(root string)`**: 查找未实现的服务器结构
- **`ListStructsMap(path string)`**: 解析和分析特定文件中的 Go 结构体
- **`GetModuleInfo(projectPath string)`**: 提取全面的模块和依赖信息

## 使用场景

**🛠 代码生成工具**: 为 Kratos 项目自动生成服务实现  
**📈 项目分析**: 分析项目架构并生成文档或指标  
**🔧 重构工具**: 理解代码依赖关系并协助重构决策  
**🚀 CI/CD 集成**: 验证项目结构并确保构建管道的一致性  
**📋 架构文档**: 生成项目结构图和 API 文档

## 演示项目

有关完整的使用示例和高级模式，请参阅：[astkratos-demos](https://github.com/orzkratos/astkratos-demos)

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/astkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/astkratos)