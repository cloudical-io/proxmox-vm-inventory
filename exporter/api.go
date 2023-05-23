package exporter

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	n := getNodes(apiURL, apiKey)

	list := make([]vm, 3)

	for _, v := range n.Data {
		r := getVMs(apiURL, apiKey, v.Node)
		list = append(list, r...)
	}

	log.Debug("Got VM List", "list", list)

	Inventory = &list

	// t := time.NewTicker(time.Minute * 5)

	// select {
	// case <-t.C:
	// 	return //todo
	// case <-Q:
	// 	os.Exit(0)
	// }

}

func request(url string, key string, path string) []byte {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	log.Debug("Requesting Proxmox API at", "url", fmt.Sprintf("%s%s", url, path))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Mallformed response body", "err", err)
		return nil
	}

	return body
}
