package sh

import (
	"io"
        "testing"

	"github.com/lotreal/docker-pods/src/sh"
)


var ps = []string{
	"CONTAINER ID        IMAGE                             COMMAND                CREATED             STATUS                      PORTS                    NAMES",
	"5ccac93e27e2        busybox:latest                    \"sh\"                   21 minutes ago      Up 21 minutes                                        lonely_goldstine        ",
	"f6362ffdb3f2        busybox:latest                    \"sh\"                   27 minutes ago      Up 27 minutes                                        desperate_perlman       ",
	"20fd25ff8f35        busybox:latest                    \"sh\"                   31 minutes ago      Up 31 minutes                                        sick_torvalds           ",
	"560a6b21ce77        busybox:latest                    \"/bin/sh\"              37 minutes ago      Up 37 minutes                                        cocky_sinoussi          ",
	"47df8dbf4de7        busybox:latest                    \"sh\"                   38 minutes ago      Exited (0) 38 minutes ago                            hungry_turing           ",
	"3c81863e1887        busybox:latest                    \"bash\"                 40 minutes ago                                                           pensive_mccarthy        ",
	"5bc06a60cd99        busybox:latest                    \"bush\"                 40 minutes ago                                                           prickly_engelbart       ",
	"7f393777bdfb        busybox:latest                    \"/bin/sh\"              41 minutes ago      Exited (0) 41 minutes ago                            distracted_babbage      ",
	"89ac6e81003f        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   3 days ago          Up 3 days                   80/tcp                   insane_kowalevski       ",
	"9788e773bb20        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   6 days ago          Up 6 days                   80/tcp                   drunk_engelbart         ",
	"006d0481dea4        hello-world:latest                \"bash\"                 9 days ago                                                               nostalgic_ardinghelli   ",
	"0b7b96fd0a2a        cedvan/toran-proxy:1.1.7-1        \"/bin/sh -c /bin/tor   11 days ago         Up 9 days                   80/tcp, 443/tcp          dreamy_hypatia          ",
	"cc5cd7551cb6        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   12 days ago         Up 9 days                   80/tcp                   stoic_pasteur           ",
	"889ea5f46cf9        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   12 days ago         Up 9 days                   80/tcp                   insane_turing           ",
	"90808059ce6e        docker.web.dm/leesin:latest       \"/usr/local/bin/star   3 weeks ago         Up 9 days                   0.0.0.0:3000->3000/tcp   sleepy_einstein         ",
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


func TestPs(t *testing.T) {
	var reader = sh.NewReader(ps, Status{})

	var status []Status
	status = make([]Status, 0)

	for {
		var s Status
		err := sh.Unmarshal(reader, &s)

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		t.Logf("%#v", s)
		status = append(status, s)
	}
	// t.Logf("%#v", status)
}
