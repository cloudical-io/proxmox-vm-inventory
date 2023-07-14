package exporter

import (
	"time"

	"github.com/cloudical-io/proxmox-vm-inventory/pkg/config"
)

var (
	Inv = &Inventory{}
)

func Run(conf config.Config) {
	Inv.inventory = make(map[string]*[]Vm, 0)

	t := time.NewTicker(time.Second * time.Duration(*conf.FetchInterval))
	for {
		for _, v := range *conf.Clusters {
			go Inv.createInventory(v, *conf.RequestTimeout)
		}
		<-t.C
	}
}
