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

func (c *Command) Raw() []byte {
	script := c.Script

	out, err := exec.Command("sh", "-c", script).Output()
	if err != nil {
		fmt.Printf("cmd: %#v\n", script)
		fmt.Printf("err: %#s", err)
	}
	out = bytes.Trim(out, "\n")
	return out
}

func (c *Command) Run() []string {
	out := string(c.Raw()[:])
	return strings.Split(out, "\n")
}

func (c *Command) Mock() []string {
	return []string{"mock"}
}
