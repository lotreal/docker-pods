package convention_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/convention"
)


var root = "/root/.go/src/github.com/lotreal/docker-pods/examples"


func TestPodsEmpty(t *testing.T) {
	pods := root + "/empty.yaml"
	_, err := convention.Pods(pods)
	if err == nil ||  err != convention.ErrNoPods {
		t.Fatalf("%s should throw error", pods)
	}
}

func TestPodsNotExist(t *testing.T) {
	pods := root + "/null.yaml"
	_, err := convention.Pods(pods)
	if err == nil ||  err != convention.ErrNoPods {
		t.Fatalf("%s should throw error", pods)
	}
}

func TestPodsDir1(t *testing.T) {
	pods := root
	p, err := convention.Pods(pods)
	if err == nil && len(p) != 1 {
		t.Fatalf("%s should return []string{1}", pods)
	}
}

func TestPodsDir2(t *testing.T) {
	pods := root + "/opt"
	p, err := convention.Pods(pods)
	if err == nil && len(p) != 2 {
		t.Fatalf("%s should return []string{2}", pods)
	}
}

func TestPodsDir3(t *testing.T) {
	pods := root + "/empty"
	_, err := convention.Pods(pods)
	if err == nil ||  err != convention.ErrNoPods {
		t.Fatalf("%s should throw error", pods)
	}
}
