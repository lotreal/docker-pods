package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/lotreal/docker-pods/src/command"
)

func TestReverse(t *testing.T) {
	fmt.Print("test")
	pods_yaml := "./examples/hello-world.yaml"

	command, err := command.Run(pods_yaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	command.Exec()
	// t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
}
