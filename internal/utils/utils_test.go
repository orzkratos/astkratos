package utils_test

import (
	"testing"

	"github.com/orzkratos/astkratos/internal/utils"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

// TestGetTrimmedLines tests file reading with line trimming
//
// TestGetTrimmedLines 测试带行修剪的文件读取
func TestGetTrimmedLines(t *testing.T) {
	data := rese.A1(utils.GetTrimmedLines(runpath.Path()))
	t.Log(neatjsons.S(data))
}

// TestGetSubstringBetween tests substring extraction between bounds
//
// TestGetSubstringBetween 测试分隔符之间的子字符串提取
func TestGetSubstringBetween(t *testing.T) {
	res := utils.GetSubstringBetween("abc", "a", "c")
	require.Equal(t, "b", res)
}
