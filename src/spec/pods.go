package spec

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type spec struct {
	Name string
}

type base struct {
	spec
}

func (p *spec) Parse(data []byte) error {
	if err := yaml.Unmarshal(data, p); err != nil {
		return err
	}
	return nil
}

func Pods(file string) (spec, error) {
	fmt.Println(file)

	var p spec

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return p, err
	}

	p.Parse(data)

	return p, err
}
