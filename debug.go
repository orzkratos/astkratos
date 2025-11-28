// Package astkratos debug utilities: Debug mode switch and logging functions
// Provides centralized debug output management with elegant logging capabilities
// Features configurable debug mode switching and optimized development workflow
// Designed with clean production output while maintaining detailed debugging support
//
// astkratos 调试工具：调试模式控制和日志功能
// 提供集中式调试输出管理和优雅的日志记录功能
// 具有可配置的调试模式切换和优化的开发工作流程
// 专为干净的生产输出而设计，同时保持详细的调试支持
package astkratos

// Debug mode switch state
//
// 调试模式开关状态
var debugModeOpen = false

// SetDebugMode enables or disables debug output
//
// SetDebugMode 启用或禁用调试输出
func SetDebugMode(enable bool) {
	debugModeOpen = enable
}

// IsDebugMode returns the current debug mode status
//
// IsDebugMode 返回当前调试模式状态
func IsDebugMode() bool {
	return debugModeOpen
}
