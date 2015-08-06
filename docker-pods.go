package main

import (
        "log"
	"os"

	"github.com/lotreal/docker-pods/command"
)


func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	command, err := command.Run(pwd + "/pods.yaml1")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	command.Print()

}
