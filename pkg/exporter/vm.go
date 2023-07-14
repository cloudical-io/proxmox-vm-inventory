package exporter

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

type vms struct {
	Data []Vm `jason:"data"`
}

type Vm struct {
	Node           string        `json:"node"`
	Status         string        `json:"status,omitempty"`
	Vmid           int           `json:"vmid"`
	Cpu            float64       `json:"cpu,omitempty"`
	Lock           string        `json:"lock,omitempty"`
	Maxdisk        int           `json:"maxdisk,omitempty"`
	Maxmem         int           `json:"maxmem,omitempty"`
	Name           string        `json:"name,omitempty"`
	Pid            int           `json:"pid,omitempty"`
	Qmpstatus      string        `json:"qmpstatus,omitempty"`
	Runningmachine string        `json:"running-machine,omitempty"`
	Runningqemu    string        `json:"running-qemu,omitempty"`
	Tags           string        `json:"tags,omitempty"`
	Uptime         int           `json:"uptime,omitempty"`
	Networks       NetworkConfig `json:"ipconfigs"`
}

// get vm list
func getVMs(apiURL string, apiKey string, node string, timeout int) ([]Vm, error) {

	uri := fmt.Sprintf("%s/nodes/%s/qemu", apiPrefix, node)
	r, err := request(apiURL, apiKey, uri, timeout)
	if err != nil {
		return nil, err
	}

	log.Debug("Proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	v := &vms{}
	if err := json.Unmarshal(r, v); err != nil {
		return nil, err
	}

	for i := range v.Data {
		v.Data[i].Node = node
	}

	return v.Data, nil
}
