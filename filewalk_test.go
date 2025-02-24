package astkratos_test

import (
	"os"
	"testing"

	"github.com/orzkratos/astkratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestWalkFiles(t *testing.T) {
	require.NoError(t, astkratos.WalkFiles(runpath.PARENT.Path(), astkratos.NewSuffixMatcher([]string{".go"}), func(path string, info os.FileInfo) error {
		t.Log(path)
		return nil
	}))
}
