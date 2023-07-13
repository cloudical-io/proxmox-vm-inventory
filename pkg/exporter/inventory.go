package exporter

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/cloudical-io/proxmox-vm-inventory/pkg/config"

	"github.com/charmbracelet/log"
)

// struct to hold all cluster information
//
// to access the content use the given methods
// since concurrency problems might appear if you do otherwise
type Inventory struct {
	inventory map[string]*[]Vm
	mu        sync.Mutex
}

func (i *Inventory) AddList(key string, vms *[]Vm) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.inventory[key] = vms
}

func (i *Inventory) GetList() map[string]*[]Vm {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.inventory
}

func (i *Inventory) GetClusterVM(s string) []Vm {
	i.mu.Lock()
	defer i.mu.Unlock()

	if v, ok := i.inventory[s]; ok {
		return *v
	} else {
		return nil
	}
}

// Takes a Cluster config object to generate a new VM List for given cluster
func (i *Inventory) createInventory(c config.Cluster, timeout int) {
	log.Info("Fetching Inventory...")

	apiKey := fmt.Sprintf("%s=%s", c.ApiUser, c.ApiKey)

	n, err := getNodes(c.ApiHost, apiKey, timeout)
	if err != nil {
		log.Warn("could not get nodes", "host", c.ApiHost)
		return
	}

	list := make([]Vm, 0)

	for _, v := range n.Data {
		if r, err := getVMs(c.ApiHost, apiKey, v.Node, timeout); err != nil {
			log.Warn("could not get nodes on node", "node", v.Node)
		} else {
			list = append(list, r...)
		}
	}

	log.Debug("Got VM List", "list", list)

	for i, v := range list {
		vmid := strconv.Itoa(v.Vmid)
		list[i].Networks, err = getNetworks(c.ApiHost, apiKey, v.Node, vmid, timeout)
		if err != nil {
			log.Warn("could not get network for vm", "vmid", vmid)
			log.Debug("error", "err", err)
		}
	}

	i.AddList(c.Name, &list)

	log.Info("Completed fetching.")
}
