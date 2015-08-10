package docker_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/docker"
)

func TestPs(t *testing.T) {
	t.Errorf("%s", docker.Ps())
}
