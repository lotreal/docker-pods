package config

import (
	"io/ioutil"
	"errors"

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

	data, err := ioutil.ReadFile(file)
        if err != nil {
		return p, err
        }

	p.Parse(data)
	return p, nil
}
