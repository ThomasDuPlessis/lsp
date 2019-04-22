package server

import (
	"net"
	"os/exec"
	"regexp"
)

type ServerCmd struct {
	pat *regexp.Regexp
	cmd string
}

type Server struct {
	cmd ServerCmd
	proc exec.Command
}

var (
	// File regex -> lsp server command
	server_cmds = []ServerCmd{
		{regexp.MustCompile(".*"), "ciderlsp"},
	}

	servers []Server
)

isInit(file string) bool {
	for _, s := range servers {
		if s.cmd.pat.MatchString(file) {
			return true
		}
	}
	return false
}

func newServer(file string) {
	var cmd string
	for _, sc := range server_cmds {
		if sc.MatchString(file) {
			p0, p1 := net.Pipe()
			proc := exec.Command(args[0], args[1:]...)
			proc.Stdin = p0
			proc.Stdout = p0
			proc.Stderr = os.Stderr
			servers = append(servers, Server{cmd:sc, proc: cmd})
			break
		}
	}
}
