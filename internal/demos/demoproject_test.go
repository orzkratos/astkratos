package demos_test

import (
	"testing"

	"github.com/orzkratos/astkratos/internal/demos"
	"github.com/yyle88/zaplog"
)

var projectPath string

func TestMain(m *testing.M) {
	projectPath = demos.CreateDemoProjectWhenNotExist()

	zaplog.SUG.Debugln(projectPath)
	m.Run()
}

func TestPath(t *testing.T) {
	t.Log(projectPath)
}
