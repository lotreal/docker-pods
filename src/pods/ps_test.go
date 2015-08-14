package pods_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/pods"
)

func TestPs(t *testing.T) {
	status := pods.GetStatus()
	t.Log("$ docker-pods ps")
	for i := 0; i < len(status); i++ {
		t.Logf("%#v", status[i])
	}
}

func TestPs2(t *testing.T) {
	pods.Ps()
}
