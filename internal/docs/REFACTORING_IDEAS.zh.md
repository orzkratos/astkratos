# AST 重构优化思路

本文档记录了 `astkratos` 项目中关于代码分析逻辑的未来优化方向。

## 背景

当前，项目中用于发现 gRPC 服务（Clients, Servers, Services）的函数（如 `ListGrpcClients`, `ListGrpcServers`）依赖于简单的“字符串逐行匹配”来解析 `_grpc.pb.go` 文件。

例如，通过检查一行代码是否以 `type` 开头、以 `Client interface {` 结尾来识别一个 gRPC 客户端。

这种方法的缺点是：
- **脆弱性**: 对代码格式（如空格、换行）、注释、甚至是 `protoc-gen-go-grpc` 代码生成器的版本更新非常敏感。任何微小的格式变化都可能导致匹配失败。
- **低效性**: `AnalyzeProject` 函数会为不同类型的定义（Client, Server 等）多次重复地读取和遍历同一个文件，造成不必要的 I/O 和计算开销。
- **不一致性**: 项目的核心目标是成为一个“基于 Go AST 的高级分析引擎”，并且已经在 `GetStructsMap` 中使用了 AST。但在最关键的 gRPC 服务发现上却使用字符串匹配，这与项目定位不符。

## 重构方案：全面采用 AST 分析

### 总体结构示意图

```
  ┌─────────────────────────────────────────────────┐
  │  analyze(path)                                  │
  │  ┌───────────────────────────────────────────┐  │
  │  │  go/parser.ParseFile() → AST              │  │
  │  │  遍历 AST 节点：                           │  │
  │  │    - InterfaceType + "Client" → Client    │  │
  │  │    - InterfaceType + "Server" → Server    │  │
  │  │    - StructType + "Unimplemented" → Stub  │  │
  │  └───────────────────────────────────────────┘  │
  │  返回: { Clients, Servers, Stubs }              │
  └─────────────────────────────────────────────────┘
```

为了使服务发现逻辑更健壮、高效，并与项目目标保持一致，建议将其重构为完全基于 Go 语言官方的 `go/ast`（抽象语法树）包进行分析。

### 核心步骤

1.  **建立统一的 AST 分析核心**
    - 创建一个新的内部函数，例如 `analyzeGrpcPbGoFile(path string)`。
    - 该函数接收 `_grpc.pb.go` 文件路径，调用 `go/parser` 将文件解析为一棵完整的 AST。
    - 遍历这棵 AST，**一次性**解析出文件中定义的所有 gRPC Client、Server 和 Unimplemented Server。

2.  **通过 AST 节点精准识别 gRPC 定义**
    - 在 AST 中，所有代码结构都是明确的节点对象。我们将通过节点的类型和属性来精准定位目标，而不是模糊的字符串匹配。
    - **gRPC Client**: 查找一个`类型定义节点 (TypeSpec)`，其名称以后缀 `Client` 结尾，且类型为`接口 (InterfaceType)`。
    - **gRPC Server**: 查找一个`类型定义节点`，其名称以后缀 `Server` 结尾，类型为`接口`，并排除 `Unsafe...` 的情况。
    - **Unimplemented Server**: 查找一个`类型定义节点`，其名称以 `Unimplemented` 开头、`Server` 结尾，且类型为`结构体 (StructType)`。

3.  **重构上层函数，整合逻辑**
    - 重构 `AnalyzeProject` 函数，使其在文件遍历时，对每个 `_grpc.pb.go` 文件只调用一次新的 `analyzeGrpcPbGoFile` 函数。
    - 将解析出的所有定义（Clients, Servers 等）缓存起来。
    - `ListGrpcClients`, `ListGrpcServers` 等现有函数不再执行任何文件 I/O，而是直接从缓存的结果中筛选并返回数据。

### 预期收益

- **健壮性**: 不再受代码格式、注释或代码生成器版本更新的影响。
- **高效率**: 每个文件只解析一次，显著减少冗余操作，提升 `AnalyzeProject` 的性能。
- **一致性**: 使项目完全符合其“基于 AST 的分析引擎”的定位，技术栈统一。
- **可维护性**: 逻辑更集中、更清晰，易于未来扩展和维护。

此方案将作为未来版本的重要优化方向。
