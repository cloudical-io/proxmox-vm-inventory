package exporter

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

func RunAPI(apiURL string, apiKey string) {
	t := time.NewTicker(time.Minute * 5)

	select {
	case <-t.C:
		return //todo
	case <-Q:
		os.Exit(0)
	}
}

func exec(url string, key string, path string) interface{} {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", url, path), nil)

	if err != nil {
		log.Error("Could not instanciate http/request", "err", err)
	}

	req.Header.Add("Authorization", key)

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Proxmox API Response Mallformed", "err", err)
	}

	return resp.Body
}
