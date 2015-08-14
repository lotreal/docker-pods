package pods_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/config"
	"github.com/lotreal/docker-pods/src/convention"
	"github.com/lotreal/docker-pods/src/pods"
)


var root = "/root/.go/src/github.com/lotreal/docker-pods/examples"


func TestMakeCmd(t *testing.T) {
	yaml, _ := convention.Pods(root)
	t.Logf("%#v", yaml)
	p, _ := config.Pods(yaml[0])
	data, err := pods.MakeCmd(p)

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%v", data)
}

func TestRunPods(t *testing.T) {
	yaml, _ := convention.Pods("/data")
	t.Logf("%#v", yaml)

	data, err := pods.RunPods(yaml[0])

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%v", data)
}

func TestRun(t *testing.T) {
	data, err := pods.Run()

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%v", data)
}
