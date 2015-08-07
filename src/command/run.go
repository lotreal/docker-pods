package command

import (
        "fmt"
	"io/ioutil"
	"strings"

	"github.com/lotreal/docker-pods/src/define"
)



func Run(yaml string) (Command, error) {
	data, err := ioutil.ReadFile(yaml)
        if err != nil {
		return Command{}, err
        }

	var p define.Pod
	if err = p.Parse(data); err != nil {
		return Command{}, err
	}

	ct := p.Spec.Containers[0]

	args := []string{}
	args = append(args, "/usr/bin/docker run -d")

	for _,s := range ct.Ports {
		args = append(args, fmt.Sprintf("-p \"%v:%v\"", s.HostPort, s.ContainerPort))
	}

	fmt.Printf("env: %v", ct.Env)
	for _,s := range ct.Env {
		args = append(args, fmt.Sprintf("-e \"%v=%v\"", s.Name, s.Value))
	}

	for _,s := range ct.VolumeMounts {
		args = append(args, fmt.Sprintf("-v \"%v:%v\"", s.HostPath, s.MountPath))
	}

	args = append(args, ct.Image)

	command := Command{ command: strings.Join(args, " \\\n") }
	return command, nil
}
