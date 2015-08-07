package sh

import (
        "fmt"
	"os/exec"
)

type Command struct {
	Script string
}

func (c *Command) Print() {
	fmt.Printf("%s", c.Script)
}

func (c *Command) Exec() {
	script := c.Script
	fmt.Printf("cmd: %#v\n", script)

	out, err := exec.Command("sh", "-c", script).Output()
	if err != nil {
		fmt.Printf("err: %#s", err)
	}
	fmt.Printf("out: %#v", out)
}
