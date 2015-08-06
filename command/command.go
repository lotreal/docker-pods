package command

import (
        "fmt"
	"os/exec"
)

type Command struct {
	command string
}

func (c *Command) Print() {
	fmt.Printf("%s", c.command)
}

func (c *Command) Exec() {
	command := c.command
	fmt.Printf("%s\n", command)

	command = "echo 1"
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	fmt.Printf("out: %s", out)
}
