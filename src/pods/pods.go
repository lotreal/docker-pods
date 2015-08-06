package pods

import (
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
        Key string
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
	Id string
	Kind string
	DesiredState struct {
		Manifest struct {
			Containers []Container
		}
	} `yaml:"desiredState"`
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
