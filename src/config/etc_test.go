package config_test

import (
	"testing"

	"github.com/lotreal/docker-pods/src/config"
)


func TestEtc(t *testing.T) {
	etc, err := config.Etc()

	for _, dir := range etc.Run {
		t.Log(dir)
	}

        if err != nil {
		t.Fatalf("Err: #v", err)
        }
}
