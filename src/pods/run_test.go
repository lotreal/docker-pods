package pods_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/convention"
	"github.com/lotreal/docker-pods/src/pods"
)


var root = "/root/.go/src/github.com/lotreal/docker-pods/examples"


func TestMakeCmd(t *testing.T) {
	yaml, _ := convention.Pods(root)
	t.Logf("%#v", yaml)
	data, err := pods.MakeCmd(yaml[0])

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%v", data)
}

func TestRun(t *testing.T) {
	yaml, _ := convention.Pods("/data")
	t.Logf("%#v", yaml)
	data, err := pods.Run(yaml[0])

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%v", data)
}
