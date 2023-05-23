package exporter

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

type nodes struct {
	Data []node `json:data`
}

type node struct {
	Maxcpu          int     `json:maxcpu`
	Level           string  `json:level`
	Mem             int     `json:mem`
	Uptime          int     `json:uptime`
	Maxmem          int     `json:maxmem`
	Status          string  `json:status`
	Node            string  `json:node`
	Disk            int     `json:disk`
	Maxdisk         int     `json:maxdisk`
	Ssl_fingerprint string  `json:ssl_fingerprint`
	Cpu             float64 `json:cpu`
	Node_type       string  `json:type`
	Id              string  `json:id`
}

func getNodes(apiURL string, apiKey string) nodes {

	//get the nodes in cluster
	r := request(apiURL, apiKey, fmt.Sprint(apiPrefix+"nodes"))

	log.Debug("Proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	//unmarshal the json object into struct
	n := &nodes{}
	if err := json.Unmarshal(r, n); err != nil {
		log.Warn("Could not Unmarshal json", "err", err)
	}

	return *n
}
