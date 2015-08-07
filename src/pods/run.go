package pods

import (
        "log"

	"github.com/lotreal/docker-pods/src/command"
	"github.com/lotreal/docker-pods/src/convention"
)

func Run() {
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
