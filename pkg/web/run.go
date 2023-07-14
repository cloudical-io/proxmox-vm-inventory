package web

import (
	"net/http"

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
		w.Write([]byte("Not Implemented"))
	})

	// endpoints
	http.HandleFunc("/all", inventoryAll)
	http.HandleFunc("/cluster", inventoryCluster)
	http.HandleFunc("/cluster/", inventoryCluster)
	http.HandleFunc("/html", serveContent)

	if err := s.ListenAndServe(); err != nil {
		log.Error("http serve error", "err", err)
	}
}
