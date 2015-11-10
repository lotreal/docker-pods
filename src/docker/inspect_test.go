package docker_test

import (
	"fmt"
	"testing"

	"github.com/lotreal/docker-pods/src/docker"
	"github.com/lotreal/docker-pods/src/sh"
)

var busybox = (func() string {
	cmd := sh.Command{"docker run -d -it busybox sh"}
	return cmd.Ok()[0:12]
})()

func TestInit(t *testing.T) {
	t.Logf("%s", docker.Ps())
	t.Logf("%s", busybox)
}

func TestIp(t *testing.T) {
	t.Log(docker.InspectIp(busybox))
}

func TestPort(t *testing.T) {
	t.Log(docker.InspectPorts(busybox))
}

func TestState(t *testing.T) {
	t.Log(docker.InspectPorts(busybox))
}

func TestRm(t *testing.T) {
	script := fmt.Sprintf("docker rm -f %s", busybox)
	cmd := sh.Command{script}
	t.Log(cmd.Ok())
}
