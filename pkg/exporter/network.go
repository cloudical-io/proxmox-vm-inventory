package exporter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
)

type NetworkConfig []string

type VmData struct {
	Data map[string]interface{} `json:"data"`
}

// get network config of a specific VM
func getNetworks(apiURL string, apiKey string, node string, vmid string, timeout int) (NetworkConfig, error) {

	uri := fmt.Sprintf("%s/nodes/%s/qemu/%s/config", apiPrefix, node, vmid)
	r, err := request(apiURL, apiKey, uri, timeout)
	//r, err := request(apiURL, apiKey, fmt.Sprint(apiPrefix+"nodes/"+node+"/qemu/"+vmid+"/config"), timeout)
	if err != nil {
		return NetworkConfig{}, err
	}

	log.Debug("proxmox API returned json", "json", fmt.Sprintf("%v", string(r)))

	var vmData VmData
	if err := json.Unmarshal(r, &vmData); err != nil {
		return NetworkConfig{}, err
	}

	var networkConfig NetworkConfig
	for k, v := range vmData.Data {
		if strings.Contains(k, "ipconfig") {
			if s, ok := v.(string); ok {
				networkConfig = append(networkConfig, s)
			}
		}
	}

	log.Debug("network devices found", "list", networkConfig)

	return networkConfig, nil
}
