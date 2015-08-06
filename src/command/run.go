package command

import (
        "fmt"
	"io/ioutil"
	"strings"

	"github.com/lotreal/docker-pods/src/pods"
)



func Run(yaml string) (Command, error) {
	data, err := ioutil.ReadFile(yaml)
        if err != nil {
		return Command{}, err
        }

	var p pods.Pod
	if err = p.Parse(data); err != nil {
		return Command{}, err
	}

	ct := p.DesiredState.Manifest.Containers[0]

	args := []string{}
	args = append(args, "/usr/bin/docker run -d")

	for _,s := range ct.Ports {
		args = append(args, fmt.Sprintf("-p \"%v:%v\"", s.HostPort, s.ContainerPort))
	}

	for _,s := range ct.Env {
		args = append(args, fmt.Sprintf("-e \"%v=%v\"", s.Key, s.Value))
	}

	for _,s := range ct.VolumeMounts {
		args = append(args, fmt.Sprintf("-v \"%v:%v\"", s.HostPath, s.MountPath))
	}

	args = append(args, ct.Image)

	command := Command{ command: strings.Join(args, " \\\n") }
	return command, nil
}
