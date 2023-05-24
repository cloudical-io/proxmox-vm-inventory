package exporter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
)

type networkconfig []string

type Vmdata struct {
	Data map[string]interface{} `json:"data"`
}

// get network config of a specific VM
func getNetworks(apiURL string, apiKey string, node string, vmid string) (networkconfig, error) {

	r, err := request(apiURL, apiKey, fmt.Sprint(apiPrefix+"nodes/"+node+"/qemu/"+vmid+"/config"))
	if err != nil {
		return networkconfig{}, err
	}

	log.Debug("Proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	var d Vmdata
	if err := json.Unmarshal(r, &d); err != nil {
		return networkconfig{}, fmt.Errorf("could not Unmarshal json %s", err)
	}

	var nc networkconfig

	for k, v := range d.Data {
		if strings.Contains(k, "ipconfig") {
			if v2, ok := v.(string); ok {
				nc = append(nc, v2)
			}
		}
	}

	log.Debug("Network Devices found", "list", nc)

	return nc, nil
}
