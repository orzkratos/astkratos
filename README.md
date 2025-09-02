# astkratos

Advanced Kratos project code structure analysis engine with Go AST capabilities.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Key Features

🔍 **gRPC Discovery Engine**: Intelligent identification and extraction of gRPC clients, servers, and services  
📊 **Struct Analysis**: Comprehensive Go struct parsing with detailed AST information and source code mapping  
📁 **Smart File Traversal**: Pattern-based file system navigation with customizable suffix matching  
📦 **Module Intelligence**: Advanced go.mod parsing with dependency analysis and toolchain version resolution  
🎯 **Code Generation Ready**: Designed for automated code generation and project structure analysis workflows

## Install

```bash
go install github.com/orzkratos/astkratos@latest
```

## Usage

### Basic gRPC Analysis

```go
package main

import (
    "fmt"
    "github.com/orzkratos/astkratos"
)

func main() {
    // List all gRPC clients in project
    clients := astkratos.ListGrpcClients("./api")
    for _, client := range clients {
        fmt.Printf("Client: %s.%s\n", client.Package, client.Name)
    }
    
    // List all gRPC servers
    servers := astkratos.ListGrpcServers("./api") 
    for _, server := range servers {
        fmt.Printf("Server: %s.%s\n", server.Package, server.Name)
    }
    
    // Extract gRPC services
    services := astkratos.ListGrpcServices("./api")
    for _, service := range services {
        fmt.Printf("Service: %s.%s\n", service.Package, service.Name)
    }
}
```

### Struct Analysis

```go
// Analyze all structs in a Go file
structs := astkratos.ListStructsMap("internal/biz/user.go")
for name, def := range structs {
    fmt.Printf("Struct: %s\n", name)
    fmt.Printf("Source: %s\n", def.StructCode)
}
```

### Module Information

```go
// Get comprehensive module information
moduleInfo, err := astkratos.GetModuleInfo(".")
if err == nil {
    fmt.Printf("Module: %s\n", moduleInfo.Module.Path)
    fmt.Printf("Go Version: %s\n", moduleInfo.GetToolchainVersion())
    
    // List dependencies
    for _, req := range moduleInfo.Require {
        status := "direct"
        if req.Indirect {
            status = "indirect"
        }
        fmt.Printf("Dependency: %s@%s (%s)\n", req.Path, req.Version, status)
    }
}
```

### Custom File Traversal

```go
// Create custom suffix matcher for specific file types
matcher := astkratos.NewSuffixMatcher([]string{".proto", "_grpc.pb.go"})

// Walk through files with custom processing
err := astkratos.WalkFiles("./api", matcher, func(path string, info os.FileInfo) error {
    fmt.Printf("Processing: %s\n", path)
    // Your custom file processing logic here
    return nil
})
```

### One-Stop Project Analysis

```go
// Get comprehensive project analysis in one call
analysis := astkratos.AnalyzeProject(".")
fmt.Printf("Project: %s\n", analysis.ModuleInfo.Module.Path)
fmt.Printf("Go Version: %s\n", analysis.ModuleInfo.GetToolchainVersion())
fmt.Printf("gRPC Clients: %d\n", analysis.ClientCount)
fmt.Printf("gRPC Servers: %d\n", analysis.ServerCount)
fmt.Printf("gRPC Services: %d\n", analysis.ServiceCount)

// Check if project has gRPC components
if astkratos.HasGrpcClients("./api") {
    fmt.Println("Project has gRPC clients")
}
```

### Debug Mode

```go
// Enable debug output to see detailed analysis process
astkratos.SetDebugMode(true)

// Run analysis with debug output
analysis := astkratos.AnalyzeProject(".")

// Disable debug output for clean results
astkratos.SetDebugMode(false)
```

## API Reference

### Core Types

- **`GrpcTypeDefinition`**: Represents gRPC type definitions with package and name information
- **`StructDefinition`**: Complete struct analysis with AST type, source code, and code snippets  
- **`ModuleInfo`**: Comprehensive Go module metadata including dependencies and toolchain info
- **`SuffixMatcher`**: Intelligent file filtering with pattern-based matching capabilities

### Primary Functions

- **`ListGrpcClients(root string)`**: Extract all gRPC client interfaces from project
- **`ListGrpcServers(root string)`**: Identify all gRPC server interfaces  
- **`ListGrpcServices(root string)`**: Discover available gRPC services
- **`ListGrpcUnimplementedServers(root string)`**: Find unimplemented server structures
- **`ListStructsMap(path string)`**: Parse and analyze Go structs in specific files
- **`GetModuleInfo(projectPath string)`**: Extract comprehensive module and dependency information

### Convenience Functions

- **`HasGrpcClients(root string)`**: Check if any gRPC clients exist
- **`HasGrpcServers(root string)`**: Check if any gRPC servers exist  
- **`CountGrpcServices(root string)`**: Get total count of gRPC services
- **`AnalyzeProject(projectRoot string)`**: Comprehensive project analysis with aggregated results

### Debug Functions

- **`SetDebugMode(enable bool)`**: Enable or disable debug output for development and troubleshooting

## Use Cases

**🛠 Code Generation Tools**: Automatically generate service implementations for Kratos projects  
**📈 Project Analysis**: Analyze project architecture and generate documentation or metrics  
**🔧 Refactoring Tools**: Understand code dependencies and assist in refactoring decisions  
**🚀 CI/CD Integration**: Validate project structure and ensure consistency in build pipelines  
**📋 Architecture Documentation**: Generate project structure diagrams and API documentation

## Demo Projects

For complete usage examples and advanced patterns, see: [astkratos-demos](https://github.com/orzkratos/astkratos-demos)

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/astkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/astkratos)