package sh

import (
	"bytes"
        "fmt"
	"os/exec"
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
	out = bytes.Trim(out, "\n")
	return string(out[:]), err
}

func (c *Command) Mock() []string {
	return []string{"mock"}
}
