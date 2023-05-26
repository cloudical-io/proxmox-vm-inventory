package exporter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
)

type networkConfig []string

type VmData struct {
	Data map[string]interface{} `json:"data"`
}

// get network config of a specific VM
func getNetworks(apiURL string, apiKey string, node string, vmid string) (networkConfig, error) {

	r, err := request(apiURL, apiKey, fmt.Sprint(apiPrefix+"nodes/"+node+"/qemu/"+vmid+"/config"))
	if err != nil {
		return networkConfig{}, err
	}

	log.Debug("proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	var vmData VmData
	if err := json.Unmarshal(r, &vmData); err != nil {
		return networkConfig{}, err
	}

	var networkConfig networkConfig
	for k, v := range vmData.Data {
		if strings.Contains(k, "ipconfig") {
			if s, ok := v.(string); ok {
				networkConfig = append(networkConfig, s)
			}
		}
	}

	log.Debug("network Devices found", "list", networkConfig)

	return networkConfig, nil
}
