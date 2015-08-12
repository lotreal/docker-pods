package config

import (
	"io/ioutil"

        "gopkg.in/yaml.v2"
)


type Config struct {
        Run []string
}

func (p *Config) Parse(data []byte) error {
	if err := yaml.Unmarshal(data, p); err != nil {
		return err
	}
	return nil
}

func Etc() (Config, error) {
	file := "/etc/docker-pods/config.yaml"

	var p Config

	data, err := ioutil.ReadFile(file)
        if err != nil {
		return p, err
        }

	p.Parse(data)
	return p, nil
}
