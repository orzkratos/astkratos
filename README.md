[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/astkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/astkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/astkratos)](https://pkg.go.dev/github.com/orzkratos/astkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/astkratos/main.svg)](https://coveralls.io/github/orzkratos/astkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/astkratos.svg)](https://github.com/orzkratos/astkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/astkratos)](https://goreportcard.com/report/github.com/orzkratos/astkratos)

# astkratos

Advanced Kratos project code structure analysis engine with Go AST capabilities.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Core Features

🔍 **gRPC Detection Engine**: Intelligent identification and extraction of gRPC clients, servers, and services  
📊 **Struct Analysis**: Comprehensive Go struct parsing with detailed AST information and source code mapping  
📁 **Smart File Walking**: Pattern-based file system navigation with customizable suffix matching  
📦 **Module Intelligence**: Advanced go.mod parsing with dependency analysis and toolchain version resolution  
🎯 **Code Generation Support**: Designed to support automated code generation and project structure analysis

## Architecture

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

## Installation

```bash
go get github.com/orzkratos/astkratos
```

## Usage

### Project Analysis

Comprehensive project analysis - combines module info and gRPC component detection:

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

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### gRPC Detection

Scan `_grpc.pb.go` files and extract gRPC type definitions (clients, servers, services):

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

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Module Information

Extract go.mod metadata including dependencies and toolchain version:

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

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

## API Reference

### Core Types

- **`GrpcTypeDefinition`**: Represents gRPC type definitions with package and name information
- **`StructDefinition`**: Complete struct analysis with AST type, source code, and code snippets
- **`ModuleInfo`**: Comprehensive Go module metadata including dependencies and toolchain info
- **`ProjectReport`**: Comprehensive project analysis with aggregated results

### Main Functions

- **`ListGrpcClients(root string)`**: Extract all gRPC client interfaces from project
- **`ListGrpcServers(root string)`**: Detect gRPC server interfaces
- **`ListGrpcServices(root string)`**: Detect available gRPC services
- **`ListGrpcUnimplementedServers(root string)`**: Find unimplemented server structures
- **`GetStructsMap(path string)`**: Parse and analyze Go structs in specific files
- **`GetModuleInfo(projectPath string)`**: Extract comprehensive module and dependency information

### Convenience Functions

- **`HasGrpcClients(root string)`**: Check if gRPC clients exist
- **`HasGrpcServers(root string)`**: Check if gRPC servers exist
- **`CountGrpcServices(root string)`**: Get the count of gRPC services
- **`AnalyzeProject(projectRoot string)`**: Comprehensive project analysis with aggregated results

### Debug Functions

- **`SetDebugMode(enable bool)`**: Enable or disable debug output for development and troubleshooting
- **`IsDebugMode()`**: Returns current debug mode status

## Examples

### Debug Mode

**Enable debug output to see detailed analysis process:**
```go
astkratos.SetDebugMode(true)
report := astkratos.AnalyzeProject(".")
astkratos.SetDebugMode(false)
```

**Check current debug status:**
```go
if astkratos.IsDebugMode() {
    fmt.Println("Debug mode is enabled")
}
```

### Struct Analysis

**Parse Go structs with AST information:**
```go
structs := astkratos.GetStructsMap("internal/biz/account.go")
for name, def := range structs {
    fmt.Printf("Struct: %s\n", name)
    fmt.Printf("Code: %s\n", def.StructCode)
}
```

**Access AST type and source code:**
```go
structs := astkratos.GetStructsMap("account.go")
accountDef := structs["Account"]
fmt.Printf("Fields count: %d\n", len(accountDef.Type.Fields.List))
```

### Unimplemented Stub Detection

**Find proto-generated unimplemented stubs:**
```go
unimplemented := astkratos.ListGrpcUnimplementedServers("./api")
for _, stub := range unimplemented {
    fmt.Printf("Stub: %s in %s\n", stub.Name, stub.Package)
}
```

### Pattern-based File Walking

**Walk files with suffix pattern matching:**
```go
pattern := utils.NewSuffixPattern([]string{"_grpc.pb.go"})
utils.WalkFiles("./api", pattern, func(path string, info os.FileInfo) error {
    fmt.Printf("Found: %s\n", path)
    return nil
})
```

## Use Cases

**🛠 Code Generation Tools**: Generate service implementations based on proto definitions
**📈 Project Analysis**: Analyze project architecture and generate documentation
**🔧 Refactoring Tools**: Understand code dependencies and assist refactoring
**🚀 CI/CD Integration**: Validate project structure in build pipelines
**📋 Architecture Documentation**: Generate project structure diagrams

## Demo Projects

Complete runnable demos with Kratos project setup: [internal/demos](internal/demos)

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/astkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/astkratos)