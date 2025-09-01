// Package astkratos module utilities: Advanced Go module information extraction and analysis
// Provides comprehensive module metadata parsing and toolchain version resolution capabilities
// Features intelligent go.mod parsing with JSON-based configuration analysis workflows
// Optimized for Kratos project dependency analysis and version compatibility verification
//
// astkratos 模块工具：高级 Go 模块信息提取和分析
// 提供全面的模块元数据解析和工具链版本解析功能
// 具有智能的 go.mod 解析和基于 JSON 的配置分析工作流程
// 针对 Kratos 项目依赖分析和版本兼容性验证优化
package astkratos

import (
	"encoding/json"

	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
	"github.com/yyle88/tern/zerotern"
)

// Module represents the core module information from go.mod
// Contains essential module path and identification data
//
// Module 代表来自 go.mod 的核心模块信息
// 包含基本的模块路径和识别数据
type Module struct {
	Path string `json:"Path"` // Module path identifier // 模块路径标识符
}

// Require represents a module dependency with version and indirect status
// Contains comprehensive dependency metadata for analysis workflows
//
// Require 代表具有版本和间接状态的模块依赖
// 包含用于分析工作流程的全面依赖元数据
type Require struct {
	Path     string `json:"Path"`     // Dependency module path // 依赖模块路径
	Version  string `json:"Version"`  // Required version // 所需版本
	Indirect bool   `json:"Indirect"` // Whether dependency is indirect // 是否为间接依赖
}

// ModuleInfo provides comprehensive Go module analysis with toolchain information
// Aggregates module metadata, dependencies, and version configuration for project analysis
// Supports intelligent toolchain version resolution and compatibility verification workflows
//
// ModuleInfo 提供带有工具链信息的全面 Go 模块分析
// 聚合模块元数据、依赖和版本配置，用于项目分析
// 支持智能的工具链版本解析和兼容性验证工作流程
type ModuleInfo struct {
	Module    *Module    `json:"Module"`    // Core module information // 核心模块信息
	Go        string     `json:"Go"`        // Go version requirement // Go 版本要求
	Toolchain string     `json:"Toolchain"` // Toolchain version if specified // 指定的工具链版本
	Require   []*Require `json:"Require"`   // Module dependencies list // 模块依赖列表
}

// GetToolchainVersion resolves the effective Go toolchain version for the module
// Returns configured toolchain version or falls back to module Go version with prefix
// Provides consistent version formatting for compatibility verification operations
//
// GetToolchainVersion 解析模块的有效 Go 工具链版本
// 返回配置的工具链版本，或回退到带前缀的模块 Go 版本
// 为兼容性验证操作提供一致的版本格式
func (a *ModuleInfo) GetToolchainVersion() string {
	return zerotern.VF(a.Toolchain, func() string {
		return "go" + a.Go // Add go prefix to version number for consistency // 为版本号添加 go 前缀以保持一致性
	})
}

// GetModuleInfo extracts comprehensive module information from the specified project
// Executes go mod edit command to retrieve JSON-formatted module metadata
// Parses and unmarshals module configuration including dependencies and toolchain info
// Returns complete ModuleInfo structure using assertion-style error handling
//
// GetModuleInfo 从指定项目提取全面的模块信息
// 执行 go mod edit 命令检索 JSON 格式的模块元数据
// 解析并反序列化包括依赖和工具链信息的模块配置
// 使用断言风格的错误处理返回完整的 ModuleInfo 结构
func GetModuleInfo(projectPath string) (*ModuleInfo, error) {
	output := rese.V1(osexec.ExecInPath(projectPath, "go", "mod", "edit", "-json"))
	var moduleInfo ModuleInfo
	must.Done(json.Unmarshal(output, &moduleInfo))
	return &moduleInfo, nil
}
