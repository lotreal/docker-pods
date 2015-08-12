package config_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/config"
)


func TestEtc(t *testing.T) {
	data, err := config.Etc()

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
	t.Logf("%#v", data)
}
