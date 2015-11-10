package docker

import (
	"io"

	"github.com/lotreal/docker-pods/src/sh"
)

type PsOutput struct {
	ContainerId string `field:"CONTAINER ID"`
	Image       string `field:"IMAGE"`
	Command     string `field:"COMMAND"`
	Created     string `field:"CREATED"`
	Status      string `field:"STATUS"`
	Ports       string `field:"PORTS"`
	Names       string `field:"NAMES"`
}

func Ps() []PsOutput {
	cmd := sh.Command{"docker ps -a"}

	lines, _ := cmd.Lines()
	var reader = sh.NewReader(lines, PsOutput{})

	var status []PsOutput
	status = make([]PsOutput, 0)

	for {
		var s PsOutput
		err := sh.Unmarshal(reader, &s)

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		status = append(status, s)
	}
	return status
}
