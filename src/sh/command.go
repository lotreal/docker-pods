package sh

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Script string
}

func (c *Command) Print() {
	fmt.Printf("%s", c.Text())
}

func (c *Command) Text() string {
	return c.Script
}

func (c *Command) Run() (string, error) {
	script := c.Script

	out, err := exec.Command("sh", "-c", script).Output()
	if err != nil {
		return "", err
	}

	out = bytes.Trim(out, "\n")
	return string(out[:]), err
}

func (c *Command) Lines() ([]string, error) {
	out, err := c.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")
	return lines, err
}

func (c *Command) Ok() string {
	out, _ := c.Run()
	return out
}

func (c *Command) Mock() []string {
	return []string{"mock"}
}
