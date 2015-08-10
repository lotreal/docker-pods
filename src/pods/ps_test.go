package pods_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/pods"
)

func TestPs(t *testing.T) {
	status := pods.Ps()
	t.Log("$ docker-pods ps")
	for i := 0; i < len(status); i++ {
		t.Logf("%#v", status[i])
	}
}
