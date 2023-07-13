package web

import (
	"encoding/json"
	"net/http"

	"github.com/cloudical-io/proxmox-vm-inventory/pkg/config"
	"github.com/cloudical-io/proxmox-vm-inventory/pkg/exporter"

	"github.com/charmbracelet/log"
)

// unused can be used later for advanced querries
// returns a list of Proxmox Clusters being Scraped
/*
func clusterList(w http.ResponseWriter, r *http.Request) {
	log.Info("Got HTTP request", "path", r.RequestURI, "origin", r.RemoteAddr, "method", r.Method)
	l := config.ClusterList()

	// marshal the struct to JSON
	if j, err := json.Marshal(l); err != nil {
		log.Error("web.go", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
}
*/

// retrieves the list of all environments
func inventoryAll(w http.ResponseWriter, r *http.Request) {
	log.Info("Got HTTP request", "path", r.RequestURI, "origin", r.RemoteAddr, "method", r.Method)

	l := exporter.Inv.GetList()

	// marshal the struct to JSON
	if j, err := json.Marshal(l); err != nil {
		log.Error("web.go", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
}

// get the list of an specific cluster
//
// requires you to POST formdata with the 'cluster' field set to the value for "cluster.apihost" you specified in your config.yaml
// returns JSON
func inventoryCluster(w http.ResponseWriter, r *http.Request) {
	log.Info("Got HTTP request", "path", r.RequestURI, "origin", r.RemoteAddr, "method", r.Method)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Try to POST your data with the 'cluster=<ClusterURL>' form-field set"))
		return
	}

	l := exporter.Inv.GetClusterVM(r.PostFormValue("cluster"))

	if l == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// marshal the struct to JSON
	if j, err := json.Marshal(l); err != nil {
		log.Error("web.go", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
}
