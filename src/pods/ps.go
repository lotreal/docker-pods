package pods

import (
        //"fmt"
	// "os"
	// "text/tabwriter"

	"github.com/lotreal/docker-pods/src/docker"
	"github.com/lotreal/docker-pods/src/sh"
)


type PsOutput struct {
	Pid         string
	ContainerId string
	Ip          string
}


func Ps() []PsOutput {
	var status []PsOutput
	status = make([]PsOutput, 0)

	ps := docker.Ps()
	for i := 0; i < len(ps); i++ {
		p := ps[i]
		cid := p.ContainerId

		status = append(status, PsOutput{
			Pid:         docker.InspectPid(cid),
			ContainerId: cid,
			Ip:          docker.InspectIp(cid),
		})
	}

	sh.TabWrite(status)

	return status

	// w := new(tabwriter.Writer)

	// // Format in tab-separated columns with a tab stop of 8.
	// w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	// fmt.Fprintln(w, "LABEL\tREPLICAS\tSTATUS\td")
	// fmt.Fprintln(w, "maokai\t2\trunning\t123456789")
	// fmt.Fprintln(w)
	// w.Flush()

	// cmd := sh.Command{"docker ps"}
	// cmd.Exec()

	// xs := []int{2, 4, 6, 8}
	// ys := []string{"C", "B", "K", "A"}
	// fmt.Println(
	// 	SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }),
	// 	SliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }),
	// 	SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }),
	// 	SliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }))
}
