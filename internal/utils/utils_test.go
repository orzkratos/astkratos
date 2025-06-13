package utils_test

import (
	"testing"

	"github.com/orzkratos/astkratos/internal/utils"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

func TestGetTrimmedLines(t *testing.T) {
	data := rese.A1(utils.GetTrimmedLines(runpath.Path()))
	t.Log(neatjsons.S(data))
}

func TestGetSubstringBetween(t *testing.T) {
	res := utils.GetSubstringBetween("abc", "a", "c")
	require.Equal(t, "b", res)
}
