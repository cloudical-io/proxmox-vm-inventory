package exporter

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

type vms struct {
	Data []vm `jason:data`
}

type vm struct {
	Node           string  `json:node,omitempty`
	Status         string  `json:status`
	Vmid           int     `json:vmid`
	Cpu            float64 `json:cpu`
	Lock           string  `json:lock`
	Maxdisk        int     `json:maxdisk`
	Maxmem         int     `json:maxmem`
	Name           string  `json:name`
	Pid            int     `json:pid`
	Qmpstatus      string  `json:qmpstatus`
	Runningmachine string  `json:running-machine`
	Runningqemu    string  `json:running-qemu`
	Tags           string  `json:tags`
	Uptime         int     `json:uptime`
}

func getVMs(apiURL string, apiKey string, node string) []vm {

	r := request(apiURL, apiKey, fmt.Sprint(apiPrefix+"nodes/"+node+"/qemu"))

	log.Debug("Proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	v := &vms{}
	if err := json.Unmarshal(r, v); err != nil {
		log.Error("Could not Unmarshal json", "err", err)
	}

	for i := range v.Data {
		v.Data[i].Node = node
	}

	return v.Data
}
