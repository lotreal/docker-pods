package sh

import (
        "fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"text/tabwriter"
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

	TableParser(out)

	// fmt.Printf("out: %#v", out)
	// n := SliceIndex(len(out), func(i int) bool { return out[i] == 0xa })
	// s := string(out[:n])
	// fmt.Printf("line 1: %v", s)

	return
	s := string(out[:])
	// fmt.Printf("out: %#v", strings.Split(s, "\n"))

	lines := strings.Split(s, "\n")
	// for _, line := range lines {
	// 	tokens := strings.Fields(line)
	// 	fmt.Printf("%#v", tokens)
	// }
	// fmt.Printf("out: %#s", s.split("\n"))

	w := new(tabwriter.Writer)


	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)

	for _, line := range lines {
		fmt.Fprintln(w, "PODS\t" + line)
	}

	// fmt.Fprintln(w, "maokai\t2\trunning\t123456789")
	fmt.Fprintln(w)
	w.Flush()

}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

type Status struct {
	ContainerId string `field:"CONTAINER ID"`
	Image       string `field:"IMAGE"`
	Command     string `field:"COMMAND"`
	Created     string `field:"CREATED"`
	Status      string `field:"STATUS"`
	Ports       string `field:"PORTS"`
	Names       string `field:"NAMES"`
}

func TableParser(in []byte) {
	lines := strings.Split(string(in[:]), "\n")
	fmt.Printf("in: %#v", lines)

	ins := Status{}
	rt  := reflect.TypeOf(ins)
	fmt.Printf("%#v", rt.NumField())


	// fmt.Printf("in: %#v", in)
	// n := SliceIndex(len(in), func(i int) bool { return in[i] == 0xa })
	// s := string(in[:n])
	// fmt.Printf("line 1: %v", s)
}
