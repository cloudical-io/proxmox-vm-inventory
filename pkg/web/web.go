package web

import (
	"encoding/json"
	"net/http"
	"proxmox-vm-inventory/pkg/exporter"

	"github.com/charmbracelet/log"
)

func Run(address string) {
	s := &http.Server{
		Addr: address,
	}

	// default handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Got HTTP request", "path", r.RequestURI, "origin", r.RemoteAddr, "method", r.Method)
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(""))
	})

	// endpoints
	http.HandleFunc("/all", inventoryAll)
	http.HandleFunc("/cluster", inventoryCluster)
	http.HandleFunc("/cluster/", inventoryCluster)

	if err := s.ListenAndServe(); err != nil {
		log.Error("http serve error", "err", err)
	}
}

// retrieves the list of all environments
func inventoryAll(w http.ResponseWriter, r *http.Request) {
	log.Info("Got HTTP request", "path", r.RequestURI, "origin", r.RemoteAddr, "method", r.Method)

	l := exporter.Inv.GetList()

	// marshal the struct to JSON
	if j, err := json.Marshal(l); err != nil {
		log.Error("weg.go", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
}

// get the list of an specific cluster
//
// reqiures you to POST formdata with the 'cluster' field set to the value for "cluster.apihost" you specified in your config.yaml
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
		log.Error("weg.go", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
}
