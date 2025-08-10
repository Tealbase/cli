package integration

import (
	"log"
	"os"
	"testing"

	"github.com/docker/docker/client"
	"github.com/spf13/viper"
	"github.com/tealbase/cli/internal/utils"
	"github.com/tealbase/cli/test/mocks/docker"
	"github.com/tealbase/cli/test/mocks/tealbase"
)

const (
	DockerPort   = ":2375"
	TealbasePort = ":2376"
)

var (
	TempDir string
)

var (
	Logger     *log.Logger
	DockerMock *docker.Server
	SupaMock   *tealbase.Server
)

func TestMain(m *testing.M) {
	Logger := log.New(os.Stdout, "", 0)

	Logger.Println("Global tests setup")

	DockerMock = newDockerMock(Logger)
	SupaMock = newTealbaseMock(Logger)
	TempDir = NewTempDir(Logger, "")

	// redirect clients to mock servers
	os.Setenv("DOCKER_HOST", "tcp://127.0.0.1"+DockerPort)
	utils.Docker = utils.NewDocker()
	if err := client.WithVersion(docker.APIVersion)(utils.Docker); err != nil {
		Logger.Fatal(err)
	}
	viper.Set("INTERNAL_API_HOST", "http://127.0.0.1"+TealbasePort)
	os.Setenv("TEALBASE_ACCESS_TOKEN", tealbase.AccessToken)
	os.Setenv("HOME", TempDir)

	// run tests
	exitVal := m.Run()

	Logger.Println("Global teardown")
	os.RemoveAll(TempDir)

	// exit process with tests exit code
	os.Exit(exitVal)
}

func newDockerMock(Logger *log.Logger) *docker.Server {
	dockerMock := docker.NewServer()
	dockerRouter := dockerMock.NewRouter()
	go func() {
		err := dockerRouter.Run(DockerPort)
		if err != nil {
			Logger.Fatal(err)
		}
	}()

	return dockerMock
}

func newTealbaseMock(Logger *log.Logger) *tealbase.Server {
	supaMock := tealbase.NewServer()
	supaRouter := supaMock.NewRouter()
	go func() {
		err := supaRouter.Run(TealbasePort)
		if err != nil {
			Logger.Fatal(err)
		}
	}()

	return supaMock
}

func NewTempDir(Logger *log.Logger, baseDir string) string {
	wd := baseDir
	var err error
	if baseDir == "" {
		wd, err = os.Getwd()
	}
	if err != nil {
		Logger.Fatal(err)
	}
	tempDir, err := os.MkdirTemp(wd, "cli-test-")
	if err != nil {
		Logger.Fatal(err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		Logger.Fatal(err)
	}
	return tempDir
}
