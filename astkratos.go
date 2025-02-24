package astkratos

import (
	"go/ast"
	"os"
	"strings"

	"github.com/orzkratos/astkratos/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/zaplog"
)

// GrpcTypeDefinition represents a gRPC type definition with its name and package.
type GrpcTypeDefinition struct {
	Name    string
	Package string
}

// ListGrpcClients lists all gRPC client types in the specified root directory.
func ListGrpcClients(root string) (definitions []*GrpcTypeDefinition) {
	must.Done(WalkFiles(root, NewSuffixMatcher([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		// Get the package name from the file path
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		sLines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range sLines {
			// Check if the line defines a gRPC client interface
			if strings.HasPrefix(s, "type ") && strings.HasSuffix(s, "Client interface {") {
				name := utils.GetSubstringBetween(s, "type ", " interface {")
				// Append the gRPC client definition to the list
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

// ListGrpcServers lists all gRPC server types in the specified root directory.
func ListGrpcServers(root string) (definitions []*GrpcTypeDefinition) {
	must.Done(WalkFiles(root, NewSuffixMatcher([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		// Get the package name from the file path
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		// Read and trim lines from the file
		sLines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range sLines {
			// Check if the line defines a gRPC server interface
			if strings.HasPrefix(s, "type ") && strings.HasSuffix(s, "Server interface {") && !strings.HasPrefix(s, "type Unsafe") {
				name := utils.GetSubstringBetween(s, "type ", " interface {")
				// Append the gRPC server definition to the list
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

// ListGrpcUnimplementedServers lists all unimplemented gRPC server types in the specified root directory.
func ListGrpcUnimplementedServers(root string) (definitions []*GrpcTypeDefinition) {
	must.Done(WalkFiles(root, NewSuffixMatcher([]string{"_grpc.pb.go"}), func(path string, info os.FileInfo) error {
		zaplog.SUG.Debugln(path)
		// Get the package name from the file path
		pkgName := syntaxgo_ast.GetPackageNameFromPath(path)
		zaplog.SUG.Debugln(pkgName)
		// Read and trim lines from the file
		sLines := rese.V1(utils.GetTrimmedLines(path))
		for _, s := range sLines {
			// Check if the line defines an unimplemented gRPC server
			if strings.HasPrefix(s, "type Unimplemented") {
				var name string
				switch {
				case strings.HasSuffix(s, "Server struct {"): // Old version
					name = utils.GetSubstringBetween(s, "type ", " struct {")
					must.OK(name)
				case strings.HasSuffix(s, "Server struct{}"): // New version
					name = utils.GetSubstringBetween(s, "type ", " struct{}")
					must.OK(name)
				}
				if name != "" { // Match old version or new version
					// Append the unimplemented gRPC server definition to the list
					definitions = append(definitions, &GrpcTypeDefinition{
						Package: pkgName,
						Name:    name,
					})
				}
			}
		}
		return nil
	}))
	return definitions
}

// ListGrpcServices lists all gRPC services in the specified root directory.
func ListGrpcServices(root string) (definitions []*GrpcTypeDefinition) {
	zaplog.SUG.Debugln(root)
	// Iterate through unimplemented gRPC servers and extract service names
	for _, x := range ListGrpcUnimplementedServers(root) {
		zaplog.SUG.Debugln(x.Name, x.Package)
		one := &GrpcTypeDefinition{
			Name:    utils.GetSubstringBetween(x.Name, "Unimplemented", "Server"),
			Package: x.Package,
		}
		must.OK(one.Name)
		// Append the gRPC service definition to the list
		definitions = append(definitions, one)
	}
	return definitions
}

// StructDefinition represents a struct definition with its name, type, source code, and code snippet.
type StructDefinition struct {
	Name       string
	Type       *ast.StructType
	FileSource []byte // The entire source code of the source file
	StructCode string // The code snippet defining the struct
}

// ListStructsMap lists all struct definitions in the specified file and returns them as a map.
func ListStructsMap(path string) map[string]*StructDefinition {
	var structMap = map[string]*StructDefinition{}

	// Read the entire source code of the file
	fileSource := rese.V1(os.ReadFile(path))
	// Parse the source code into an AST bundle
	astBundle := rese.P1(syntaxgo_ast.NewAstBundleV1(fileSource))
	astFile, _ := astBundle.GetBundle()

	// Map struct types by their names
	structTypes := syntaxgo_search.MapStructTypesByName(astFile)
	zaplog.SUG.Debugln(len(structTypes))
	for structName, structType := range structTypes {
		// Get the code snippet defining the struct
		structCode := syntaxgo_astnode.GetText(fileSource, structType)
		zaplog.SUG.Debugln(structName, structCode)
		// Add the struct definition to the map
		structMap[structName] = &StructDefinition{
			Name:       structName,
			Type:       structType,
			FileSource: fileSource,
			StructCode: structCode,
		}
	}
	return structMap
}
