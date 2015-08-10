package sh

import (
        "testing"

	"github.com/lotreal/docker-pods/src/sh"
)




func TestPs(t *testing.T) {
	ps := []string{"CONTAINER ID        IMAGE                             COMMAND                CREATED             STATUS              PORTS                    NAMES", "89ac6e81003f        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   3 days ago          Up 3 days           80/tcp                   insane_kowalevski   ", "9788e773bb20        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   6 days ago          Up 6 days           80/tcp                   drunk_engelbart     ", "0b7b96fd0a2a        cedvan/toran-proxy:1.1.7-1        \"/bin/sh -c /bin/tor   11 days ago         Up 9 days           80/tcp, 443/tcp          dreamy_hypatia      ", "cc5cd7551cb6        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   12 days ago         Up 9 days           80/tcp                   stoic_pasteur       ", "889ea5f46cf9        docker.web.dm/apache-php:latest   \"/usr/local/bin/star   12 days ago         Up 9 days           80/tcp                   insane_turing       ", "90808059ce6e        docker.web.dm/leesin:latest       \"/usr/local/bin/star   3 weeks ago         Up 9 days           0.0.0.0:3000->3000/tcp   sleepy_einstein     "}
	cmd := sh.Command{"docker ps"}
	t.Logf("%#v", cmd.Run())
	t.Logf("%#v", ps)
}

func TestPsaq(t *testing.T) {
	cmd := sh.Command{"docker ps -aq"}
	t.Logf("%#v", cmd.Run())
}

func TestEcho(t *testing.T) {
	cmd := sh.Command{"echo 12; echo 34"}
	out := cmd.Run()
	if len(out) != 2 {
		t.Errorf("command(%s) should output 2 line, but is: %#v", cmd.Text(), out)
	}
}

func TestEchon(t *testing.T) {
	cmd := sh.Command{"echo -n 12; echo 34"}
	out := cmd.Run()
	if len(out) != 1 {
		t.Errorf("command(%s) should output 1 line, but is: %#v", cmd.Text(), out)
	}
}
