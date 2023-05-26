package main

import (
	"os"
	"os/signal"
	"proxmox-vm-inventory/pkg/config"
	"proxmox-vm-inventory/pkg/exporter"
	"proxmox-vm-inventory/pkg/web"
	"syscall"

	"github.com/charmbracelet/log"
)

func main() {
	log.Print(`
	___ _________  __  ______  _   _  __  _   __  __  ______  _____________   _ 
	|__]|__/|  | \/ |\/||  | \/    |  ||\/|   ||\ ||  ||___|\ | | |  ||__/ \_/  
	|   |  \|__|_/\_|  ||__|_/\_    \/ |  |   || \| \/ |___| \| | |__||  \  |   
`)

	c := config.New()

	go web.Run()

	go exporter.Run(*c)

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGTERM, syscall.SIGINT)

	<-q
	log.Info("Caught Shutdown Signal, Terminating.")
}
