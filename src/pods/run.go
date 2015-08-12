package pods

import (
	"fmt"
	"path"
	"strings"

	"github.com/lotreal/docker-pods/src/config"
	"github.com/lotreal/docker-pods/src/sh"
)

func MakeCmd(file string) (string, error) {
	wd := path.Dir(file)

	p, err := config.Pods(file)
	if err != nil {
		return "", err
	}

	ct := p.Spec.Containers[0]

	args := []string{}
	args = append(args, "/usr/bin/docker run -d")

	args = append(args, fmt.Sprintf("-l \"pid=%v\"", p.Id))

	for _,s := range ct.Ports {
		args = append(args, fmt.Sprintf("-p \"%v:%v\"", s.HostPort, s.ContainerPort))
	}

	for _,s := range ct.Env {
		args = append(args, fmt.Sprintf("-e \"%v=%v\"", s.Name, s.Value))
	}

	for _,s := range ct.VolumeMounts {
		args = append(args, fmt.Sprintf("-v \"%s/%s:%s\"", wd, s.HostPath, s.MountPath))
	}

	args = append(args, ct.Image)

	command := strings.Join(args, " \\\n")
	return command, nil
}

func Run(file string) (string, error) {
	script, err := MakeCmd(file)
	if err != nil {
		return "", err
	}

	cmd := sh.Command{script}
	return cmd.Run()[0], nil
}
