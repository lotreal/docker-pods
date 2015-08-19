package pods

import (
	"fmt"
	"strings"

	"github.com/lotreal/docker-pods/src/config"
	"github.com/lotreal/docker-pods/src/convention"
	"github.com/lotreal/docker-pods/src/sh"
)


type RunOutput struct {
	Pid         string `field:"PODS ID"`
	Pods        string `field:"PODS"`
	ContainerId string `field:"CONTAINER ID"`
}

func MakeCmd(p config.Pod) (string, error) {
	etc, _ := config.Etc()

	ct := p.Spec.Containers[0]

	args := []string{}
	args = append(args, "/usr/bin/docker run -d")

	args = append(args, fmt.Sprintf("-l \"pid=%v\"", p.Id))

	for _,s := range ct.Ports {
		args = append(args, fmt.Sprintf("-p \"%v:%v\"", s.HostPort, s.ContainerPort))
	}

	for _,s := range etc.Env {
		args = append(args, fmt.Sprintf("-e \"%v=%v\"", s.Name, s.Value))
	}

	for _,s := range ct.Env {
		args = append(args, fmt.Sprintf("-e \"%v=%v\"", s.Name, s.Value))
	}

	for _,s := range ct.VolumeMounts {
		args = append(args, fmt.Sprintf("-v \"%s:%s\"", s.HostPath, s.MountPath))
	}

	args = append(args, ct.Image)

	command := strings.Join(args, " \\\n")
	return command, nil
}

func RunPods(file string) (RunOutput, error) {
	var out RunOutput
	p, err := config.Pods(file)
	if err != nil {
		return out, err
	}

	script, err := MakeCmd(p)
	if err != nil {
		return out, err
	}

	cmd := sh.Command{script}
	fmt.Println(file)
	fmt.Println(script)
	fmt.Println("=============================")
	out = RunOutput{
		Pid: p.Id,
		Pods: file,
		ContainerId: cmd.Run()[0],
		// ContainerId: cmd.Mock()[0],
	}

	return out, nil
}

func GetStart() ([]RunOutput, error) {
	out := make([]RunOutput, 0)

	pods := []string{}
	etc, err := config.Etc()
	if err != nil {
		return out, err
	}

	for _, dir := range etc.Run {
		p, _ := convention.Pods(dir)
		pods = append(pods, p...)
	}

	for _, p := range pods {
		o, _ := RunPods(p)
		out = append(out, o)
		// out = append(out, RunOutput{
		// 	ContainerId: "x",
		// })

		// ps := GetStatus()
		// fmt.Println(ps)
	}

	return out, nil
}

func Run() (string, error) {
	out, _ := GetStart()

	sh.TabWrite(out)
	return "ok", nil
}
