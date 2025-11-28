// Demo program to showcase astkratos module info extraction with real Kratos project
// Creates temporary Kratos project and demonstrates module metadata parsing
// Outputs module path, Go version, toolchain info and dependency analysis
//
// 演示程序展示 astkratos 在真实 Kratos 项目中的模块信息提取功能
// 创建临时 Kratos 项目并演示模块元数据解析
// 输出模块路径、Go 版本、工具链信息和依赖分析
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
	output := rese.A1(osexec.NewExecConfig().WithPath(tempDIR).Exec("kratos", "new", "demo3kratos"))
	zaplog.SUG.Debugln(string(output))

	// Get project root path for module analysis
	// 获取项目根目录路径用于模块分析
	projectRoot := osmustexist.ROOT(filepath.Join(tempDIR, "demo3kratos"))

	// Extract module information from go.mod
	// 从 go.mod 提取模块信息
	moduleInfo := rese.P1(astkratos.GetModuleInfo(projectRoot))
	zaplog.SUG.Debugln("Module Info:")
	zaplog.SUG.Debugln(neatjsons.S(moduleInfo))

	// Display toolchain version
	// 显示工具链版本
	zaplog.SUG.Debugln("Toolchain Version:", moduleInfo.GetToolchainVersion())

	// Check gRPC component existence
	// 检查 gRPC 组件是否存在
	apiPath := osmustexist.ROOT(filepath.Join(projectRoot, "api"))
	zaplog.SUG.Debugln("Has gRPC Clients:", astkratos.HasGrpcClients(apiPath))
	zaplog.SUG.Debugln("Has gRPC Servers:", astkratos.HasGrpcServers(apiPath))
	zaplog.SUG.Debugln("gRPC Service Count:", astkratos.CountGrpcServices(apiPath))
}
