// Demo program to showcase astkratos functionality with real Kratos project
// Creates temporary Kratos project and demonstrates comprehensive analysis capabilities
// Outputs complete project analysis including gRPC clients, servers, and services
//
// 演示程序展示 astkratos 在真实 Kratos 项目中的功能
// 创建临时 Kratos 项目并演示全面的分析能力
// 输出完整的项目分析结果，包括 gRPC 客户端、服务器和服务
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/orzkratos/astkratos"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
)

func main() {
	// Create temporary directory to hold the demo project
	// 创建临时目录存放演示项目
	tempDIR := rese.V1(os.MkdirTemp("", "go-kratos-test-*"))
	defer func() {
		// Clean up temporary files when done
		// 完成后清理临时文件
		must.Done(os.RemoveAll(tempDIR))
	}()

	// Generate new Kratos project using kratos CLI tool
	// 使用 kratos CLI 工具生成新的 Kratos 项目
	output := rese.A1(osexec.NewExecConfig().WithPath(tempDIR).Exec("kratos", "new", "demo1kratos"))
	fmt.Println(string(output))

	// Analyze the generated project structure and gRPC definitions
	// 分析生成的项目结构和 gRPC 定义
	projectRoot := osmustexist.ROOT(filepath.Join(tempDIR, "demo1kratos"))
	result := astkratos.AnalyzeProject(projectRoot)
	fmt.Println(neatjsons.S(result))
}
