package exporter

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

type nodes struct {
	Data []node `json:"data"`
}

type node struct {
	Maxcpu          int     `json:"maxcpu,omitempty"`
	Level           string  `json:"level,omitempty"`
	Mem             int     `json:"mem,omitempty"`
	Uptime          int     `json:"uptime,omitempty"`
	Maxmem          int     `json:"maxmem,omitempty"`
	Status          string  `json:"status,omitempty"`
	Node            string  `json:"node,omitempty"`
	Disk            int     `json:"disk,omitempty"`
	Maxdisk         int     `json:"maxdisk,omitempty"`
	Ssl_fingerprint string  `json:"ssl_fingerprint,omitempty"`
	Cpu             float64 `json:"cpu,omitempty"`
	Node_type       string  `json:"type,omitempty"`
	Id              string  `json:"id,omitempty"`
}

func getNodes(apiURL string, apiKey string) nodes {

	//get the nodes in cluster
	r := request(apiURL, apiKey, fmt.Sprint(apiPrefix+"nodes"))

	log.Debug("Proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	//unmarshal the json object into struct
	n := &nodes{}
	if err := json.Unmarshal(r, n); err != nil {
		log.Error("Could not Unmarshal json", "err", err)
		return nodes{}
	}

	return *n
}
