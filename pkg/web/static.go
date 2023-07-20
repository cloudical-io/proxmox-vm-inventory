package web

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/cloudical-io/proxmox-vm-inventory/pkg/exporter"

	"github.com/charmbracelet/log"
)

func serveContent(w http.ResponseWriter, r *http.Request) {
	log.Info("Got HTTP request", "path", r.RequestURI, "origin", r.RemoteAddr, "method", r.Method)

	l := exporter.Inv.GetList()

	t := template.New("").Funcs(template.FuncMap{
		"joinNetworks": func(n exporter.NetworkConfig) string {
			return strings.Join(n, ", ")
		},
	})

	t, err := t.Parse(table)
	if err != nil {
		log.Error("Failed Parsing Template", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := t.Execute(w, l); err != nil {
		log.Error("Failed Compiling Template", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serveSortableJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")
	if _, err := w.Write([]byte(sortablejs)); err != nil {
		log.Error("Could Not Serve File", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
