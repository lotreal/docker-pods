package docker_test

import (
	"fmt"
	"testing"

	"github.com/lotreal/docker-pods/src/docker"
	"github.com/lotreal/docker-pods/src/sh"
)

var busybox = (func() string {
	return "5ccac93e27e25b836878a0cc7d18a2882605d85ae3c8bc347924aef343022c26"[0:12]

	cmd := sh.Command{"docker run -d -it busybox sh"}
	return cmd.Run()[0][0:12]
})()

func TestInit(t *testing.T) {
	t.Logf("%s", docker.Ps())
	t.Logf("%s", busybox)
}

func TestIp(t *testing.T) {
	script := fmt.Sprintf("docker inspect --format='{{.NetworkSettings.IPAddress}}' %s", busybox)
	cmd := sh.Command{script}
	t.Log(cmd.Run())
}

func TestPort(t *testing.T) {
	script := fmt.Sprintf("docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' %s", busybox)
	cmd := sh.Command{script}
	t.Log(cmd.Run())
}

func TestState(t *testing.T) {
	script := fmt.Sprintf("docker inspect --format='{{.State.Running}}' %s", busybox)
	cmd := sh.Command{script}
	t.Log(cmd.Run())
}


func TestSn2(t *testing.T) {
	// script := fmt.Sprintf("", busybox)
	// cmd := sh.Command{script}
	// t.Log(cmd.Run())

	t.Logf("%#v", 2)
}

func TestRm(t *testing.T) {
	script := fmt.Sprintf("docker rm -f %s", busybox)
	t.Log(script)
	// cmd := sh.Command{script}
	// t.Log(cmd.Run())
}
