package astkratos_test

import (
	"testing"

	"github.com/orzkratos/astkratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestGetModuleInfo(t *testing.T) {
	moduleInfo, err := astkratos.GetModuleInfo(runpath.PARENT.Path())
	require.NoError(t, err)
	t.Log(neatjsons.S(moduleInfo))
	require.Equal(t, "go1.22.8", moduleInfo.GetToolchainVersion())
}
