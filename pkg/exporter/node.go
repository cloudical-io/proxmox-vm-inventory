package exporter

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

type nodes struct {
	Data []Node `json:"data"`
}

type Node struct {
	Id              string  `json:"id"`
	Node            string  `json:"node"`
	Node_type       string  `json:"type,omitempty"`
	Maxcpu          int     `json:"maxcpu,omitempty"`
	Level           string  `json:"level,omitempty"`
	Mem             int     `json:"mem,omitempty"`
	Uptime          int     `json:"uptime,omitempty"`
	Maxmem          int     `json:"maxmem,omitempty"`
	Status          string  `json:"status,omitempty"`
	Disk            int     `json:"disk,omitempty"`
	Maxdisk         int     `json:"maxdisk,omitempty"`
	Ssl_fingerprint string  `json:"ssl_fingerprint,omitempty"`
	Cpu             float64 `json:"cpu,omitempty"`
}

// retrieves the node list from a Proxmox Cluster
//
// returns a struct from type Node.
// On error returns the message and an empty struct.
func getNodes(apiURL string, apiKey string, timeout int) (nodes, error) {

	//get the nodes in cluster
	uri := fmt.Sprintf("%s/nodes", apiPrefix)
	r, err := request(apiURL, apiKey, uri, timeout)
	if err != nil {
		return nodes{}, err
	}

	log.Debug("Proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	//unmarshal the json object into struct
	n := &nodes{}
	if err := json.Unmarshal(r, n); err != nil {
		return nodes{}, fmt.Errorf("could not unmarshal json %s", err)
	}

	return *n, nil
}
