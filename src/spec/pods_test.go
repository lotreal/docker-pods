package spec_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/spec"
)

func TestPs(t *testing.T) {
	// file := "~/.go/src/github.com/lotreal/docker-pods/examples/v2.yaml"
	file := "/Users/luotao/.go/src/github.com/lotreal/docker-pods/examples/v2.yaml"
	p, _ := spec.Pods(file)
	t.Logf("%#v", p)
}
