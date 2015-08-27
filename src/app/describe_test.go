package app_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/app"
)

func TestPs(t *testing.T) {
	t.Logf("%#v", app.Describe("."))
	t.Logf("%#v", app.Describe("/a/maokai"))
}
