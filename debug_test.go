package astkratos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIsDebugMode tests debug mode status checking
//
// TestIsDebugMode 测试调试模式状态检查
func TestIsDebugMode(t *testing.T) {
	require.False(t, IsDebugMode())
}
