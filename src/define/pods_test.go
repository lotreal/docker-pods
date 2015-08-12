package define_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/convention"
	"github.com/lotreal/docker-pods/src/define"
)


var root = "/root/.go/src/github.com/lotreal/docker-pods/examples"


func TestPods(t *testing.T) {
	yaml, _ := convention.Pods(root)
	data, err := define.Pods(yaml[0])

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%#v", data)
}
