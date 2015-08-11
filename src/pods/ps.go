package pods

import (
        //"fmt"
	// "os"
	// "text/tabwriter"

	"github.com/lotreal/docker-pods/src/docker"
	"github.com/lotreal/docker-pods/src/sh"
)


type PsOutput struct {
	Running     string `field:"RUNNING"`
	Pid         string `field:"PODS ID"`
	ContainerId string `field:"CONTAINER ID"`
	Ip          string `field:"IP ADDR"`
	Image       string `field:"IMAGE"`
	Command     string `field:"COMMAND"`
	Status      string `field:"STATUS"`
	Ports       string `field:"PORTS"`
}


func Ps() []PsOutput {
	var status []PsOutput
	status = make([]PsOutput, 0)

	ps := docker.Ps()
	for i := 0; i < len(ps); i++ {
		p := ps[i]
		cid := p.ContainerId

		status = append(status, PsOutput{
			Running:      docker.InspectRunning(cid),
			Pid:         docker.InspectPid(cid),
			ContainerId: cid,
			Ip:          docker.InspectIp(cid),
			Image:       p.Image,
			Command:     p.Command,
			Ports:       p.Ports,
			Status:      p.Status,
		})
	}

	sh.TabWrite(status)

	return status
}
