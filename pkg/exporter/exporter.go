package exporter

import (
	"proxmox-vm-inventory/pkg/config"
	"time"
)

var (
	Inv = &Inventory{}
)

func Run(conf config.Config) {
	Inv.inventory = make(map[string]*[]vm, 0)

	t := time.NewTicker(time.Second * time.Duration(conf.FetchInterval))
	for {
		for _, v := range conf.Clusters {
			go Inv.createInventory(v, conf.RequestTimeout)
		}
		<-t.C
	}
}
