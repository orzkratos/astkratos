// Package astkratos: Advanced Kratos project code structure analysis engine with Go AST
// Provides comprehensive gRPC service detection, client/service interface extraction, and struct parsing
// Features intelligent file walking, pattern matching, and module information analysis in Kratos projects
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
	"path/filepath"
	"strings"

	"github.com/orzkratos/astkratos/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/zaplog"
)

// GrpcTypeDefinition represents a gRPC type definition with its name and package
//
// GrpcTypeDefinition 表示 gRPC 类型定义，包含名称和包信息
type GrpcTypeDefinition struct {
	Name    string // Name of the gRPC type // gRPC 类型名称
	Package string // Package name where the type is defined // 类型定义所在的包名
	SrcPath string // Source file path where the type is defined // 类型定义所在的源文件路径
}

// ListGrpcClients lists gRPC client types in the specified root path
//
// ListGrpcClients 列出指定根目录下的 gRPC 客户端类型
func ListGrpcClients(root string) (definitions []*GrpcTypeDefinition) {
	definitions = make([]*GrpcTypeDefinition, 0)

	must.Done(utils.WalkFiles(root, utils.NewSuffixPattern([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		// Get the package name from the file path
		// 从文件路径获取包名
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		// 读取并修剪文件中的行
		lines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range lines {
			// Check if the line defines a gRPC client interface
			// 检查该行是否定义了 gRPC 客户端接口
			if strings.HasPrefix(s, "type ") && strings.HasSuffix(s, "Client interface {") {
				name := utils.GetSubstringBetween(s, "type ", " interface {")
				// Append the gRPC client definition to the list
				// 将 gRPC 客户端定义添加到列表中
				definitions = append(definitions, &GrpcTypeDefinition{
					Name:    name,
					Package: pkgName,
					SrcPath: rese.V1(filepath.Abs(path)),
				})
			}
		}
		return nil
	}))
	return definitions
}

// ListGrpcServers lists gRPC server types in the specified root path
//
// ListGrpcServers 列出指定根目录下的 gRPC 服务器类型
func ListGrpcServers(root string) (definitions []*GrpcTypeDefinition) {
	definitions = make([]*GrpcTypeDefinition, 0)

	must.Done(utils.WalkFiles(root, utils.NewSuffixPattern([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		// Get the package name from the file path
		// 从文件路径获取包名
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		// 读取并修剪文件中的行
		lines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range lines {
			// Check if the line defines a gRPC server interface
			// 检查该行是否定义了 gRPC 服务器接口
			if strings.HasPrefix(s, "type ") && strings.HasSuffix(s, "Server interface {") && !strings.HasPrefix(s, "type Unsafe") {
				name := utils.GetSubstringBetween(s, "type ", " interface {")
				// Append the gRPC server definition to the list
				// 将 gRPC 服务器定义添加到列表中
				definitions = append(definitions, &GrpcTypeDefinition{
					Name:    name,
					Package: pkgName,
					SrcPath: rese.V1(filepath.Abs(path)),
				})
			}
		}
		return nil
	}))
	return definitions
}

// ListGrpcUnimplementedServers lists unimplemented gRPC server types in the specified root path
//
// ListGrpcUnimplementedServers 列出指定根目录下的未实现 gRPC 服务器类型
func ListGrpcUnimplementedServers(root string) (definitions []*GrpcTypeDefinition) {
	if debugModeOpen {
		zaplog.SUG.Debugln("discovering unimplemented gRPC servers in project:", root)
	}

	definitions = make([]*GrpcTypeDefinition, 0)

	must.Done(utils.WalkFiles(root, utils.NewSuffixPattern([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		if debugModeOpen {
			zaplog.SUG.Debugln("examining generated protobuf source:", path)
		}
		// Get the package name from the file path
		// 从文件路径获取包名
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		// 读取并修剪文件中的行
		lines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range lines {
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
						Name:    name,
						Package: pkgName,
						SrcPath: rese.V1(filepath.Abs(path)),
					})
				}
			}
		}
		return nil
	}))

	if debugModeOpen {
		zaplog.SUG.Debugln("discovered unimplemented server definitions:", neatjsons.S(definitions))
	}
	return definitions
}

// ListGrpcServices lists gRPC services in the specified root path
//
// ListGrpcServices 列出指定根目录下的 gRPC 服务
func ListGrpcServices(root string) (definitions []*GrpcTypeDefinition) {
	if debugModeOpen {
		zaplog.SUG.Debugln("resolving service definitions in project:", root)
	}

	definitions = make([]*GrpcTypeDefinition, 0)
	// Iterate through unimplemented gRPC servers and extract service names
	// 遍历未实现的 gRPC 服务器并提取服务名称
	for _, unimplement := range ListGrpcUnimplementedServers(root) {
		if debugModeOpen {
			zaplog.SUG.Debugln("identified service:", unimplement.Name, "within package:", unimplement.Package)
		}
		serviceName := utils.GetSubstringBetween(unimplement.Name, "Unimplemented", "Server")
		must.OK(serviceName)
		// Append the gRPC service definition to the list
		// 将 gRPC 服务定义添加到列表中
		definitions = append(definitions, &GrpcTypeDefinition{
			Name:    serviceName,
			Package: unimplement.Package,
			SrcPath: unimplement.SrcPath,
		})
	}

	if debugModeOpen {
		zaplog.SUG.Debugln("resolved service definitions:", neatjsons.S(definitions))
	}
	return definitions
}

// StructDefinition represents a struct definition with its name, type, source code, and code snippet
//
// StructDefinition 表示结构体定义，包含名称、类型、源码和代码片段
type StructDefinition struct {
	Name       string          // Struct name // 结构体名称
	Type       *ast.StructType // AST representation of the struct // 结构体的 AST 表示
	FileSource []byte          // The entire source code of the source file // 源文件的完整源码
	StructCode string          // The code snippet defining the struct // 定义结构体的代码片段
}

// GetStructsMap gets struct definitions in the specified file and returns them as a map
//
// GetStructsMap 获取指定文件中的结构体定义并返回映射表
func GetStructsMap(path string) map[string]*StructDefinition {
	if debugModeOpen {
		zaplog.SUG.Debugln("parsing Go struct definitions from:", path)
	}

	structMap := map[string]*StructDefinition{}

	// Read the entire source code of the file
	// 读取文件的完整源代码
	fileSource := rese.V1(os.ReadFile(path))
	// Parse the source code into an AST bundle
	// 将源代码解析为 AST 包
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(fileSource))
	astFile, _ := astBundle.GetBundle()

	// Map struct types based on name
	// 按名称映射结构体类型
	structTypes := syntaxgo_search.MapStructTypesByName(astFile)
	for structName, structType := range structTypes {
		// Get the code snippet defining the struct
		// 获取定义结构体的代码片段
		structCode := syntaxgo_astnode.GetText(fileSource, structType)
		if debugModeOpen {
			zaplog.SUG.Debugln("extracted struct:", structName, "with source:", structCode)
		}
		// Add the struct definition to the map
		// 将结构体定义添加到映射表中
		structMap[structName] = &StructDefinition{
			Name:       structName,
			Type:       structType,
			FileSource: fileSource,
			StructCode: structCode,
		}
	}

	if debugModeOpen {
		zaplog.SUG.Debugln("struct parsing completed, discovered", len(structMap), "definitions")
	}
	return structMap
}

// HasGrpcClients checks if gRPC clients exist in the specified root path
// Returns true when at least one gRPC client is found
//
// HasGrpcClients 检查指定根目录下是否存在 gRPC 客户端
// 找到至少一个 gRPC 客户端时返回 true
func HasGrpcClients(root string) bool {
	return len(ListGrpcClients(root)) > 0
}

// HasGrpcServers checks if gRPC services exist in the specified root path
// Returns true when at least one gRPC service is found
//
// HasGrpcServers 检查指定根目录下是否存在 gRPC 服务器
// 找到至少一个 gRPC 服务器时返回 true
func HasGrpcServers(root string) bool {
	return len(ListGrpcServers(root)) > 0
}

// CountGrpcServices returns the count of gRPC services in the specified root path
// Provides quick statistics without returning the complete service list
//
// CountGrpcServices 返回指定根目录下 gRPC 服务的总数
// 提供快速统计信息而无需返回完整的服务列表
func CountGrpcServices(root string) int {
	return len(ListGrpcServices(root))
}

// ProjectReport provides comprehensive Kratos project analysis results
// Aggregates analysis data including gRPC services, module info, and file counts
//
// ProjectReport 提供全面的 Kratos 项目分析结果
// 聚合分析数据，包括 gRPC 服务、模块信息和文件统计
type ProjectReport struct {
	ModuleInfo *ModuleInfo           `json:"moduleInfo"` // Module and dependency information // 模块和依赖信息
	Clients    []*GrpcTypeDefinition `json:"clients"`    // List of gRPC clients // gRPC 客户端列表
	Servers    []*GrpcTypeDefinition `json:"servers"`    // List of gRPC servers // gRPC 服务器列表
	Services   []*GrpcTypeDefinition `json:"services"`   // List of gRPC services // gRPC 服务列表
}

// AnalyzeProject performs comprehensive Kratos project analysis
// Scans gRPC definitions and extracts complete module information in one operation
// Returns aggregated project analysis with discovered components and metadata
//
// AnalyzeProject 执行 Kratos 项目的全面分析
// 扫描 gRPC 定义并在一次操作中提取完整的模块信息
// 返回包含发现组件和元数据的聚合项目分析
func AnalyzeProject(projectRoot string) *ProjectReport {
	// Get module information first
	// 首先获取模块信息
	moduleInfo := rese.P1(GetModuleInfo(projectRoot))

	// Analyze gRPC components in API path
	// 分析 API 目录中的 gRPC 组件
	apiPath := osmustexist.ROOT(filepath.Join(projectRoot, "api"))

	// Build comprehensive report
	// 构建全面报告
	return &ProjectReport{
		ModuleInfo: moduleInfo,
		Clients:    ListGrpcClients(apiPath),
		Servers:    ListGrpcServers(apiPath),
		Services:   ListGrpcServices(apiPath),
	}
}
