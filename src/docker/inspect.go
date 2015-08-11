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


// func TestPort() {
// 	script := fmt.Sprintf("docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' %s", cid)
// 	cmd := sh.Command{script}
// 	t.Log(cmd.Run())
// }

func InspectRunning(cid string) string {
	script := fmt.Sprintf("docker inspect --format='{{.State.Running}}' %s", cid)
	cmd := sh.Command{script}
	ret := cmd.Run()[0]
	if ret == "true" {
		return "YES"
	}
	return "NO"
}


// func TestSn2() {
// 	// script := fmt.Sprintf("", cid)
// 	// cmd := sh.Command{script}
// 	// t.Log(cmd.Run())

// 	t.Logf("%#v", 2)
// }

// func TestRm() {
// 	script := fmt.Sprintf("docker rm -f %s", cid)
// 	t.Log(script)
// 	// cmd := sh.Command{script}
// 	// t.Log(cmd.Run())
// }
