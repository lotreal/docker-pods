package sh_test

import (
	"github.com/lotreal/docker-pods/src/sh"
	"testing"
)

func TestRun(t *testing.T) {
	cmd := sh.Command{"echo -n 12; echo 34"}
	out, _ := cmd.Run()
	if out != "1234" {
		t.Errorf("command(%s) should output 1234, but is: %#v", cmd.Text(), out)
	}
}

func TestLines(t *testing.T) {
	cmd := sh.Command{"echo 12; echo 34"}
	out, _ := cmd.Lines()
	if len(out) != 2 {
		t.Errorf("command(%s) should output 2 line, but is: %#v", cmd.Text(), out)
	}
}

func TestFoo(t *testing.T) {
	cmd := sh.Command{"foo"}
	_, err := cmd.Run()
	if err.Error() != "exit status 127" {
		t.Errorf("command(%s) should exit with 127, but is: %#v", cmd.Text(), err.Error())
	}
}
