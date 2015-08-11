package docker

import (
	"fmt"

	"github.com/lotreal/docker-pods/src/sh"
)

func InspectIp(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.NetworkSettings.IPAddress}}' %s", cid)
	cmd := sh.Command{script}
	return cmd.Run()[0]
}

func InspectPid(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.Config.Labels.pid}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Run()[0]
	if ret == "<no value>" {
		return ""
	}
	return ret
}


func InspectPorts(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Run()[0]
	return ret
}

func InspectRunning(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.State.Running}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Run()[0]
	if ret == "true" {
		return "YES"
	}
	return "NO"
}
