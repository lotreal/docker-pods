package docker_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/docker"
)

func TestStatus(t *testing.T) {
	status := docker.Status()
	t.Log("$ docker ps -a")
	for i := 0; i < len(status); i++ {
		t.Logf("%s", status[i])
	}
}
