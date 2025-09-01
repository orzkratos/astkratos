// Package astkratos file traversal utilities: Intelligent file system navigation for code analysis
// Provides pattern-based file matching and recursive directory traversal capabilities
// Features customizable suffix matching and callback-based file processing workflows
// Optimized for Kratos project structure analysis with flexible filtering options
//
// astkratos 文件遍历工具：用于代码分析的智能文件系统导航
// 提供基于模式的文件匹配和递归目录遍历功能
// 具有可定制的后缀匹配和基于回调的文件处理工作流程
// 针对 Kratos 项目结构分析优化，具有灵活的过滤选项
package astkratos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/erero"
)

// SuffixMatcher provides intelligent file extension matching for selective file processing
// Contains configurable suffix patterns for flexible file filtering operations
// Supports multiple suffix matching with optimized string comparison algorithms
//
// SuffixMatcher 提供智能的文件扩展名匹配，用于选择性文件处理
// 包含可配置的后缀模式，支持灵活的文件过滤操作
// 支持多后缀匹配，具有优化的字符串比较算法
type SuffixMatcher struct {
	suffixes []string // List of file suffixes for matching // 用于匹配的文件后缀列表
}

// NewSuffixMatcher creates a new SuffixMatcher with specified suffix patterns
// Initializes matcher with custom suffix list for targeted file selection
// Returns configured matcher ready for file pattern matching operations
//
// NewSuffixMatcher 创建一个具有指定后缀模式的新 SuffixMatcher
// 使用自定义后缀列表初始化匹配器，用于目标文件选择
// 返回配置好的匹配器，准备进行文件模式匹配操作
func NewSuffixMatcher(suffixes []string) *SuffixMatcher {
	return &SuffixMatcher{
		suffixes: suffixes,
	}
}

// Match performs suffix-based string matching against configured patterns
// Tests if input string ends with any of the predefined suffixes
// Returns true if match found, false otherwise for efficient filtering
//
// Match 对配置的模式执行基于后缀的字符串匹配
// 测试输入字符串是否以任何预定义后缀结尾
// 如果找到匹配则返回 true，否则返回 false 以实现高效过滤
func (sm *SuffixMatcher) Match(s string) bool {
	for _, suffix := range sm.suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// WalkFiles performs recursive directory traversal with intelligent file filtering
// Applies callback function to files matching the specified suffix patterns
// Provides comprehensive error handling and skip non-matching files automatically
// Returns aggregated error from traversal or callback execution failures
//
// WalkFiles 执行带有智能文件过滤的递归目录遍历
// 对匹配指定后缀模式的文件应用回调函数
// 提供全面的错误处理并自动跳过不匹配的文件
// 返回来自遍历或回调执行失败的聚合错误
func WalkFiles(root string, suffixMatcher *SuffixMatcher, run func(path string, info os.FileInfo) error) error {
	if err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return erero.Wro(err)
			}
			if info == nil || info.IsDir() {
				return nil
			}
			if suffixMatcher.Match(path) {
				return run(path, info)
			}
			return nil
		},
	); err != nil {
		return erero.Wro(err)
	}
	return nil
}
