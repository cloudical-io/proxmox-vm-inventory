package exporter

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
)

var (
	apiPrefix string = "/api2/json/"
	q                = make(chan os.Signal, 1)
	Inventory *[]vm
)

func init() {
	signal.Notify(q, syscall.SIGTERM, syscall.SIGINT)
}

func Exporter(apiURL string, apiKey string) {
	t := time.NewTicker(time.Second * 60)

	for {
		select {
		case <-t.C:
			go func() {
				log.Info("Fetching Inventory...")

				n := getNodes(apiURL, apiKey)

				list := make([]vm, 0)

				for _, v := range n.Data {
					r := getVMs(apiURL, apiKey, v.Node)
					list = append(list, r...)
				}

				log.Debug("Got VM List", "list", list)

				for i, v := range list {
					vmid := strconv.Itoa(v.Vmid)
					list[i].Networks = getNetworks(apiURL, apiKey, v.Node, vmid)
				}

				Inventory = &list

				log.Info("Completed fetching.")
				log.Debug("Got inventory", "data", Inventory)
			}()
		case <-q:
			log.Info("Caught Shutdown Signal, Terminating...")
			return
		}
	}

}

// general request handler
func request(url string, key string, path string) []byte {
	log.Debug("Requesting Proxmox API at", "url", fmt.Sprintf("%s%s", url, path))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", url, path), nil)
	if err != nil {
		log.Error("Could not instanciate http/request", "err", err)
	}

	// adding the Proxmox API Token to the request
	t := fmt.Sprintf("PVEAPIToken=%s", key)
	req.Header.Add("Authorization", t)

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Proxmox API Response Mallformed", "err", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("Proxmox API returned non 200 status code", "status", resp.StatusCode, "message", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Mallformed response body", "err", err)
		log.Debug("Response", "response", resp.Body, "body", body)
		return []byte{}
	}

	return body
}
