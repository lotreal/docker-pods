package pods

import (
        "fmt"
	"os"
	"text/tabwriter"

	"github.com/lotreal/docker-pods/src/sh"
)

func Status() {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintln(w, "LABEL\tREPLICAS\tSTATUS\td")
	fmt.Fprintln(w, "maokai\t2\trunning\t123456789")
	fmt.Fprintln(w)
	w.Flush()

	cmd := sh.Command{"docker ps"}
	cmd.Exec()
}
