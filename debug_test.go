package astkratos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsDebugMode(t *testing.T) {
	require.False(t, IsDebugMode())
}
