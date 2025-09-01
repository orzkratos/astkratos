// Package astkratos: Advanced Kratos project code structure analysis engine with Go AST
// Provides comprehensive gRPC service discovery, client/server interface extraction, and struct parsing
// Features intelligent file traversal, pattern matching, and module information analysis for Kratos projects
// Supports automated code generation, project analysis, and development workflow optimization
//
// astkratos: 基于 Go AST 的高级 Kratos 项目代码结构分析引擎
// 提供全面的 gRPC 服务发现、客户端/服务器接口提取和结构体解析功能
// 具有智能文件遍历、模式匹配和 Kratos 项目模块信息分析特性
// 支持自动化代码生成、项目分析和开发工作流程优化
package astkratos

import (
	"go/ast"
	"os"
	"strings"

	"github.com/orzkratos/astkratos/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/zaplog"
)

// GrpcTypeDefinition represents a gRPC type definition with its name and package
// GrpcTypeDefinition 表示 gRPC 类型定义，包含名称和包信息
type GrpcTypeDefinition struct {
	Name    string // Name of the gRPC type // gRPC 类型名称
	Package string // Package name where the type is defined // 类型定义所在的包名
}

// ListGrpcClients lists all gRPC client types in the specified root directory
// ListGrpcClients 列出指定根目录下的所有 gRPC 客户端类型
func ListGrpcClients(root string) (definitions []*GrpcTypeDefinition) {
	definitions = make([]*GrpcTypeDefinition, 0)

	must.Done(WalkFiles(root, NewSuffixMatcher([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		// Get the package name from the file path
		// 从文件路径获取包名
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		// 读取并修剪文件中的行
		sLines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range sLines {
			// Check if the line defines a gRPC client interface
			// 检查该行是否定义了 gRPC 客户端接口
			if strings.HasPrefix(s, "type ") && strings.HasSuffix(s, "Client interface {") {
				name := utils.GetSubstringBetween(s, "type ", " interface {")
				// Append the gRPC client definition to the list
				// 将 gRPC 客户端定义添加到列表中
				definitions = append(definitions, &GrpcTypeDefinition{
					Package: pkgName,
					Name:    name,
				})
			}
		}
		return nil
	}))
	return definitions
}

// ListGrpcServers lists all gRPC server types in the specified root directory
// ListGrpcServers 列出指定根目录下的所有 gRPC 服务器类型
func ListGrpcServers(root string) (definitions []*GrpcTypeDefinition) {
	definitions = make([]*GrpcTypeDefinition, 0)

	must.Done(WalkFiles(root, NewSuffixMatcher([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		// Get the package name from the file path
		// 从文件路径获取包名
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		// 读取并修剪文件中的行
		sLines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range sLines {
			// Check if the line defines a gRPC server interface
			// 检查该行是否定义了 gRPC 服务器接口
			if strings.HasPrefix(s, "type ") && strings.HasSuffix(s, "Server interface {") && !strings.HasPrefix(s, "type Unsafe") {
				name := utils.GetSubstringBetween(s, "type ", " interface {")
				// Append the gRPC server definition to the list
				// 将 gRPC 服务器定义添加到列表中
				definitions = append(definitions, &GrpcTypeDefinition{
					Package: pkgName,
					Name:    name,
				})
			}
		}
		return nil
	}))
	return definitions
}

// ListGrpcUnimplementedServers lists all unimplemented gRPC server types in the specified root directory
// ListGrpcUnimplementedServers 列出指定根目录下的所有未实现的 gRPC 服务器类型
func ListGrpcUnimplementedServers(root string) (definitions []*GrpcTypeDefinition) {
	zaplog.SUG.Debugln("list-grpc-unimplemented-servers in path:", root)

	definitions = make([]*GrpcTypeDefinition, 0)

	must.Done(WalkFiles(root, NewSuffixMatcher([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		zaplog.SUG.Debugln("xyz_grpc.pb.go path:", path)
		// Get the package name from the file path
		// 从文件路径获取包名
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		// 读取并修剪文件中的行
		sLines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range sLines {
			// Check if the line defines an unimplemented gRPC server
			// 检查该行是否定义了未实现的 gRPC 服务器
			if strings.HasPrefix(s, "type Unimplemented") {
				var name string
				switch {
				case strings.HasSuffix(s, "Server struct {"): // Old version // 旧版本格式
					name = utils.GetSubstringBetween(s, "type ", " struct {")
					must.OK(name)
				case strings.HasSuffix(s, "Server struct{}"): // New version // 新版本格式
					name = utils.GetSubstringBetween(s, "type ", " struct{}")
					must.OK(name)
				}
				if name != "" { // Match old version or new version // 匹配旧版本或新版本格式
					// Append the unimplemented gRPC server definition to the list
					// 将未实现的 gRPC 服务器定义添加到列表中
					definitions = append(definitions, &GrpcTypeDefinition{
						Package: pkgName,
						Name:    name,
					})
				}
			}
		}
		return nil
	}))

	zaplog.SUG.Debugln("list-grpc-unimplemented-servers definitions:", neatjsons.S(definitions))
	return definitions
}

// ListGrpcServices lists all gRPC services in the specified root directory
// ListGrpcServices 列出指定根目录下的所有 gRPC 服务
func ListGrpcServices(root string) (definitions []*GrpcTypeDefinition) {
	zaplog.SUG.Debugln("list-grpc-services in path:", root)

	definitions = make([]*GrpcTypeDefinition, 0)
	// Iterate through unimplemented gRPC servers and extract service names
	// 遍历未实现的 gRPC 服务器并提取服务名称
	for _, x := range ListGrpcUnimplementedServers(root) {
		zaplog.SUG.Debugln("service-name:", x.Name, "package-name:", x.Package)
		one := &GrpcTypeDefinition{
			Name:    utils.GetSubstringBetween(x.Name, "Unimplemented", "Server"),
			Package: x.Package,
		}
		must.OK(one.Name)
		// Append the gRPC service definition to the list
		// 将 gRPC 服务定义添加到列表中
		definitions = append(definitions, one)
	}

	zaplog.SUG.Debugln("list-grpc-services definitions:", neatjsons.S(definitions))
	return definitions
}

// StructDefinition represents a struct definition with its name, type, source code, and code snippet
// StructDefinition 表示结构体定义，包含名称、类型、源码和代码片段
type StructDefinition struct {
	Name       string          // Struct name // 结构体名称
	Type       *ast.StructType // AST representation of the struct // 结构体的 AST 表示
	FileSource []byte          // The entire source code of the source file // 源文件的完整源码
	StructCode string          // The code snippet defining the struct // 定义结构体的代码片段
}

// ListStructsMap lists all struct definitions in the specified file and returns them as a map
// ListStructsMap 列出指定文件中的所有结构体定义并返回映射表
func ListStructsMap(path string) map[string]*StructDefinition {
	zaplog.SUG.Debugln("list-structs-map in path:", path)

	var structMap = map[string]*StructDefinition{}

	// Read the entire source code of the file
	// 读取文件的完整源代码
	fileSource := rese.V1(os.ReadFile(path))
	// Parse the source code into an AST bundle
	// 将源代码解析为 AST 包
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(fileSource))
	astFile, _ := astBundle.GetBundle()

	// Map struct types by their names
	// 按名称映射结构体类型
	structTypes := syntaxgo_search.MapStructTypesByName(astFile)
	for structName, structType := range structTypes {
		// Get the code snippet defining the struct
		// 获取定义结构体的代码片段
		structCode := syntaxgo_astnode.GetText(fileSource, structType)
		zaplog.SUG.Debugln("struct-name:", structName, "struct-code:", structCode)
		// Add the struct definition to the map
		// 将结构体定义添加到映射表中
		structMap[structName] = &StructDefinition{
			Name:       structName,
			Type:       structType,
			FileSource: fileSource,
			StructCode: structCode,
		}
	}

	zaplog.SUG.Debugln("list-structs-map struct-map-size:", len(structMap))
	return structMap
}
