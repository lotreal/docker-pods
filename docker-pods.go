package main

import (
        "log"

	"github.com/lotreal/docker-pods/command"
	"github.com/lotreal/docker-pods/convention"
)


func main() {
	pods_yaml, err := convention.Pods()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	command, err := command.Run(pods_yaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	command.Exec()
}
