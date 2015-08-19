package config

import (
	"io/ioutil"
	"errors"
	"fmt"
	"path"

        "gopkg.in/yaml.v2"
)


type Port struct {
        ContainerPort int `yaml:"containerPort"`
	HostPort int `yaml:"hostPort"`
}

type VolumeMounts struct {
        Name string
	HostPath string `yaml:"hostPath"`
	MountPath string `yaml:"mountPath"`
}

type Env struct {
        Name string
	Value string
}

type Container struct {
        Name string
	Image string

	Ports []Port
	VolumeMounts []VolumeMounts `yaml:"volumeMounts"`
	Env []Env
}

type Pod struct {
	Id   string
	Kind string
	Spec struct {
		Containers []Container
	}
}

func (p *Pod) Parse(data []byte) error {
	if err := yaml.Unmarshal(data, p); err != nil {
		return err
	}
	if p.Kind != "Pod" {
		return errors.New("pods: invalid `kind`")
	}
	return nil
}

func Pods(file string) (Pod, error) {
	var p Pod
	etc, _ := Etc()

	data, err := ioutil.ReadFile(file)
        if err != nil {
		return p, err
        }

	p.Parse(data)

	wd := path.Dir(file)
	vm := p.Spec.Containers[0].VolumeMounts

	for i := 0; i < len(vm); i++ {
		if vm[i].Name == "config" {
			vm[i].HostPath = fmt.Sprintf("%s/etc/%s/%s", wd,
				etc.Env[0].Value,
				vm[i].HostPath)
		} else {
			vm[i].HostPath = fmt.Sprintf("%s/%s", wd,
				vm[i].HostPath)
		}
	}

	return p, nil
}
