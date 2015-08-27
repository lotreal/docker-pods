package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/lotreal/docker-pods/src/sh"
)

func gitRev(dir string) (string, error) {
	os.Chdir(dir)

	script := "git rev-list --count --first-parent HEAD"
	cmd := sh.Command{script}
	out, err := cmd.Run()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("r%s", out), err
}

func gitDescribe(dir string) (string, error) {
	os.Chdir(dir)
	script := "git describe"
	cmd := sh.Command{script}

	out, err := cmd.Run()
	return out, err
}

func Describe(dir string) string {
	desc, err := gitDescribe(dir)

	if err != nil {
		desc, _ = gitRev(dir)
		// return desc
	}

	// /a/maokai => xx-a
	expectPS := strings.Split(strings.Trim(dir, "/"), "/")
	if len(expectPS) == 2 {
		desc = fmt.Sprintf("%s-%s", desc, expectPS[0])
	}

	return desc
}

func main() {
	fmt.Printf("%#v", Describe("."))
	fmt.Printf("%#v", Describe("/a/maokai"))
}
