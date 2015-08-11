package convention_test

import (
	"fmt"
	"testing"

	"github.com/lotreal/docker-pods/src/convention"
)


var root = "/root/.go/src/github.com/lotreal/docker-pods/examples"


func TestPodsEmpty(t *testing.T) {
	pods := root + "/empty.yaml"
	_, err := convention.CheckPods(pods)
	if err == nil ||  err.Error() != fmt.Sprintf("%s is empty", pods) {
		t.Fatalf("%s should notice empty error", pods)
	}
}

func TestPodsNotExist(t *testing.T) {
	pods := root + "/null.yaml"
	_, err := convention.CheckPods(pods)
	if err == nil ||  err.Error() != fmt.Sprintf("%s not exist", pods) {
		t.Fatalf("%s should notice not exist error", pods)
	}
}

func TestPodsIsDir(t *testing.T) {
	pods := root
	_, err := convention.CheckPods(pods)
	if err == nil ||  err.Error() != fmt.Sprintf("%s is dir", pods) {
		t.Fatalf("%s should notice is dir error", pods)
	}
}


func TestPods2(t *testing.T) {
	t.Log(convention.FindPodsInDir("/data/maokai"))
}

func TestPods3(t *testing.T) {
	t.Log(convention.FindPodsInDir("."))
}
