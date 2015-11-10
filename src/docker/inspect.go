package docker

import (
	"fmt"

	"github.com/lotreal/docker-pods/src/sh"
)

func InspectIp(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.NetworkSettings.IPAddress}}' %s", cid)
	cmd := sh.Command{script}
	return cmd.Ok()
}

func InspectPid(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.Config.Labels.name}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Ok()
	if ret == "<no value>" {
		return ""
	}
	return ret
}

func InspectPorts(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Ok()
	return ret
}

func InspectRunning(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.State.Running}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Ok()
	if ret == "true" {
		return "YES"
	}
	return "NO"
}
