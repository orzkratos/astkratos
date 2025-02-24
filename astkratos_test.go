package astkratos_test

import (
	"path/filepath"
	"testing"

	"github.com/orzkratos/astkratos"
	"github.com/orzkratos/astkratos/internal/demos"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/zaplog"
)

var projectPath string

func TestMain(m *testing.M) {
	projectPath = demos.CreateDemoProjectWhenNotExist()

	zaplog.SUG.Debugln(projectPath)
	m.Run()
}

func TestListGrpcClients(t *testing.T) {
	definitions := astkratos.ListGrpcClients(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))
}

func TestListGrpcServers(t *testing.T) {
	definitions := astkratos.ListGrpcServers(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))
}

func TestListGrpcServices(t *testing.T) {
	definitions := astkratos.ListGrpcServices(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))
}

func TestListGrpcUnimplementedServers(t *testing.T) {
	definitions := astkratos.ListGrpcUnimplementedServers(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))
}

func TestListStructsMap(t *testing.T) {
	structsMap := astkratos.ListStructsMap(filepath.Join(projectPath, "api/helloworld/v1/greeter.pb.go"))
	t.Log(len(structsMap))

	for name, definition := range structsMap {
		t.Log(name)

		t.Log(definition.Name, definition.StructCode)
	}
}
