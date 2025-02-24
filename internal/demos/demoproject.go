package demos

import (
	"github.com/orzkratos/astkratos"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
)

const projectName = "demoproject"

func CreateDemoProjectWhenNotExist() string {
	projectPath := runpath.PARENT.Join(projectName)
	if osmustexist.IsRoot(projectPath) {
		return projectPath
	}
	CreateDemoProject()
	osmustexist.MustRoot(projectPath)
	return projectPath
}

func CreateDemoProject() {
	execNewDemoProject()

	projectPath := runpath.PARENT.Join(projectName)
	osmustexist.MustRoot(projectPath)
	moduleInfo := rese.P1(astkratos.GetModuleInfo(projectPath))
	execConfig := osexec.NewExecConfig().
		WithDebug().
		WithEnvs([]string{"GOTOOLCHAIN=" + moduleInfo.GetToolchainVersion()}).
		WithPath(projectPath).
		WithMatchPipe(func(line string) bool {
			return true
		}).
		WithMatchMore(true)

	{
		output := rese.V1(execConfig.ExecInPipe("make", "init"))
		zaplog.SUG.Debugln(string(output))
	}

	{
		output := rese.V1(execConfig.ExecInPipe("make", "all"))
		zaplog.SUG.Debugln(string(output))
	}
}

func execNewDemoProject() {
	output := rese.V1(osexec.NewExecConfig().
		WithPath(runpath.PARENT.Path()).
		Exec("kratos", "new", projectName))
	zaplog.SUG.Debugln(string(output))
}
