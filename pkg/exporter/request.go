package exporter

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"proxmox-vm-inventory/pkg/config"
	"strconv"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
)

var (
	apiPrefix string = "/api2/json/"
	q                = make(chan os.Signal, 1)
	Inventory map[string]*[]vm
)

func init() {
	signal.Notify(q, syscall.SIGTERM, syscall.SIGINT)
}

func Exporter(conf config.Config) {
	t := time.NewTicker(time.Second * time.Duration(conf.FetchInterval))

	Inventory = make(map[string]*[]vm, len(conf.Clusters))

	for {
		select {
		case <-t.C:
			for _, v := range conf.Clusters {
				go fetchNode(v)
			}
		case <-q:
			log.Info("Caught Shutdown Signal, Terminating...")
			return
		}
	}

}

func fetchNode(c config.Cluster) {
	log.Info("Fetching Inventory...")

	apiKey := fmt.Sprintf("%s=%s", c.ApiUser, c.ApiKey)

	n, err := getNodes(c.ApiHost, apiKey)
	if err != nil {
		log.Warn("could not get nodes", "host", c.ApiHost)
		return
	}

	list := make([]vm, 0)

	for _, v := range n.Data {

		if r, err := getVMs(c.ApiHost, apiKey, v.Node); err != nil {
			log.Warn("could not get nodes on node", "node", v.Node)
		} else {
			list = append(list, r...)
		}

	}

	log.Debug("Got VM List", "list", list)

	for i, v := range list {
		vmid := strconv.Itoa(v.Vmid)
		list[i].Networks, err = getNetworks(c.ApiHost, apiKey, v.Node, vmid)
		if err != nil {
			log.Warn("could not get network for vm", "vmid", vmid)
			log.Debug("error", "err", err)
		}
	}

	j, _ := json.Marshal(Inventory)

	Inventory[c.ApiHost] = &list

	log.Info("Completed fetching.")
	log.Debug("Got inventory", "json", string(j))
}

// general request handler
func request(url string, key string, path string) ([]byte, error) {
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
		return nil, fmt.Errorf("Could not instanciate http/request %s", err)
	}

	// adding the Proxmox API Token to the request
	t := fmt.Sprintf("PVEAPIToken=%s", key)
	req.Header.Add("Authorization", t)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Proxmox API Response Mallformed %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("Proxmox API returned non 200 status code", "status", resp.StatusCode, "message", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Debug("Response", "response", resp.Body, "body", body)
		return nil, fmt.Errorf("Mallformed response body %s", err)
	}

	return body, nil
}
