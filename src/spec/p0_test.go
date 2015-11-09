package config_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/config"
	// "github.com/lotreal/docker-pods/src/convention"
)

var root = "/root/.go/src/github.com/lotreal/docker-pods/examples"

func TestPods(t *testing.T) {
	// yaml, _ := convention.Pods(root)

	yaml := "/etc/docker-pods/maokai.yaml"
	data, err := config.Pods(yaml)

	if err != nil {
		t.Fatalf("Err: #v", err)
	}
	t.Logf("%#v", data)
	t.Logf("%v", data)
}
