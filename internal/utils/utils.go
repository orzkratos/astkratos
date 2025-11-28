// Package utils: Advanced text processing utilities in code analysis workflows
// Provides intelligent file reading and string manipulation capabilities in AST operations
// Features line-based file processing with comprehensive whitespace trimming and substring extraction
// Optimized in efficient text parsing and code structure analysis in Kratos projects
//
// utils: 用于代码分析工作流程的高级文本处理工具
// 为 AST 操作提供智能的文件读取和字符串操作功能
// 具有基于行的文件处理，包含全面的空白字符修剪和子字符串提取
// 针对 Kratos 项目中高效的文本解析和代码结构分析优化
package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/rese"
)

// GetTrimmedLines performs intelligent line-based file reading with whitespace normalization
// Reads complete file content and returns list of trimmed lines used in efficient text processing
// Provides comprehensive EOF handling and ensures last line inclusion with exception management
// Returns clean line list suitable in code analysis and pattern matching operations
//
// GetTrimmedLines 执行智能的逐行文件读取，带有空白字符标准化
// 读取完整文件内容并返回修剪后的行数组，用于高效的文本处理
// 提供全面的 EOF 处理，确保最后一行包含和正确的错误管理
// 返回适合代码分析和模式匹配操作的清洁行数组
func GetTrimmedLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, erero.Wro(err)
	}
	defer rese.F0(file.Close)

	var lines []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, strings.TrimSpace(scan.Text()))
	}
	if err := scan.Err(); err != nil {
		return nil, erero.Wro(err)
	}
	return lines, nil
}

// GetSubstringBetween performs intelligent substring extraction between specified bounds
// Locates start and end points within input string and extracts content between them
// Excludes bound points from result and handles edge cases with robust checking
// Returns extracted content as blank string if no valid substring pattern found
//
// GetSubstringBetween 执行指定分隔符之间的智能子字符串提取
// 在输入字符串中定位开始和结束标记，并提取它们之间的内容
// 从结果中排除分隔符标记，并通过强健的边界检查处理边缘情况
// 返回提取的内容，如果没有找到有效的子字符串模式则返回空字符串
func GetSubstringBetween(s string, sSub, eSub string) string {
	if sIdx, eIdx := strings.Index(s, sSub), strings.LastIndex(s, eSub); sIdx >= 0 && eIdx >= 0 && eIdx >= sIdx+len(sSub) {
		return s[sIdx+len(sSub) : eIdx]
	}
	return "" // Return blank string if no substring is found // 如果没有找到子字符串则返回空字符串
}
