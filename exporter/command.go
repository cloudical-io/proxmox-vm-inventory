package exporter

import (
	"os"
	"os/signal"
	"syscall"
)

type nodes struct {
	node []string `json:nodes`
}

type vm struct {
	status         string `json:status`
	vmid           string `json:vmid`
	cpu            string `json:cpu`
	lock           string `json:lock`
	maxdisk        string `json:maxdisk`
	maxmem         string `json:maxmem`
	name           string `json:name`
	pid            string `json:pid`
	qmpstatus      string `json:qmpstatus`
	runningmachine string `json:running-machine`
	runningqemu    string `json:running-qemu`
	tags           string `json:tags`
	uptime         string `json:maxdisk`
}

var (
	Q = make(chan os.Signal, 1)
)

func init() {
	signal.Notify(Q, syscall.SIGTERM, syscall.SIGINT)
}
