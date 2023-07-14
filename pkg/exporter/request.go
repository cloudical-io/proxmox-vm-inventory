package exporter

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

var (
	apiPrefix string = "/api2/json"
)

// general request handler
func request(url string, key string, path string, timeout int) ([]byte, error) {
	log.Debug("Requesting Proxmox API at", "url", fmt.Sprintf("%s%s", url, path))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * time.Duration(timeout),
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", url, path), nil)
	if err != nil {
		return nil, fmt.Errorf("could not instantiate http-request %s", err)
	}

	// adding the Proxmox API Token to the request
	t := fmt.Sprintf("PVEAPIToken=%s", key)
	req.Header.Add("Authorization", t)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("proxmox API response malformed %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("proxmox API returned non 200 status code, code %d, message %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Debug("Response", "response", resp.Body, "body", body)
		return nil, fmt.Errorf("malformed response body %s", err)
	}

	return body, nil
}
