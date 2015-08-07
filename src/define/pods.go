package define

import (

	"errors"

        "gopkg.in/yaml.v2"
)


type Port struct {
        ContainerPort int `yaml:"containerPort"`
	HostPort int `yaml:"hostPort"`
}

// func (p *Port) ContainerPort() int {
// 	fmt.Printf("port: %v", p)
// 	return p.containerPort
// }

// func (p *Port) HostPort() int {
// 	if p.hostPort == 0 {
// 		return p.ContainerPort
// 	}
// 	return p.hostPort
// }

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
