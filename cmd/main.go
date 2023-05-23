package cmd

import (
	"proxmox-vm-inventory/exporter"

	"github.com/alecthomas/kingpin/v2"
)

type authMethod string

const (
	authShell authMethod = "pvesh"
	authAPI   authMethod = "api"
)

func (a authMethod) toString() string {
	return string(a)
}

// variables to be parsed
var (
	authType = kingpin.Flag("auth-method", "Authentication Method").
			Short('a').
			Required().
			Enum(authShell.toString(), authAPI.toString())
	apiKey = kingpin.Flag("api-key", "API Key").Short('P').String()
	apiURL = kingpin.Flag("api-host", "API URL / Endpoint").Short('H').String()
)

func init() {
	kingpin.Parse()
}

func main() {
	switch *authType {
	case authShell.toString():
		exporter.RunShell()
		return
	case authAPI.toString():
		exporter.RunAPI(*apiURL, *apiKey)
	}
}
