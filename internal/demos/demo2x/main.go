// Demo program to showcase astkratos gRPC discovery with real Kratos project
// Creates temporary Kratos project and demonstrates gRPC client/server extraction
// Outputs detailed gRPC type definitions including package and name information
//
// 演示程序展示 astkratos 在真实 Kratos 项目中的 gRPC 发现功能
// 创建临时 Kratos 项目并演示 gRPC 客户端/服务器提取
// 输出详细的 gRPC 类型定义，包括包名和类型名信息
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
	output := rese.A1(osexec.NewExecConfig().WithPath(tempDIR).Exec("kratos", "new", "demo2kratos"))
	zaplog.SUG.Debugln(string(output))

	// Get API directory path for gRPC analysis
	// 获取 API 目录路径用于 gRPC 分析
	projectRoot := osmustexist.ROOT(filepath.Join(tempDIR, "demo2kratos"))
	apiPath := osmustexist.ROOT(filepath.Join(projectRoot, "api"))

	// List all gRPC clients in the project
	// 列出项目中的所有 gRPC 客户端
	clients := astkratos.ListGrpcClients(apiPath)
	zaplog.SUG.Debugln("gRPC Clients:")
	zaplog.SUG.Debugln(neatjsons.S(clients))

	// List all gRPC servers in the project
	// 列出项目中的所有 gRPC 服务器
	servers := astkratos.ListGrpcServers(apiPath)
	zaplog.SUG.Debugln("gRPC Servers:")
	zaplog.SUG.Debugln(neatjsons.S(servers))

	// List all gRPC services in the project
	// 列出项目中的所有 gRPC 服务
	services := astkratos.ListGrpcServices(apiPath)
	zaplog.SUG.Debugln("gRPC Services:")
	zaplog.SUG.Debugln(neatjsons.S(services))
}
