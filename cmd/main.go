package main

import (
	"fmt"
	"proxmox-vm-inventory/exporter"

	"github.com/alecthomas/kingpin/v2"
	"github.com/charmbracelet/log"
)

// variables to be parsed
var (
	apiKey = kingpin.
		Flag("api-key", "API Key").
		Short('P').
		Required().
		Envar("API_KEY").
		String()
	apiUser = kingpin.
		Flag("api-user", "API Key").
		Short('U').
		Required().
		Envar("API_USER").
		String()
	apiURL = kingpin.
		Flag("api-host", "API URL / Endpoint").
		Short('H').
		Required().
		Envar("API_HOST").
		String()
	logLevel = kingpin.
			Flag("log-level", "Set the Log Level / verbosity").
			Short('l').
			Envar("LOG_LEVEL").
			Default("WARN").
			Enum("DEBUG", "INFO", "WARN", "ERROR", "FATAL")
)

func init() {
	kingpin.Parse()

	l := log.DebugLevel

	switch *logLevel {
	case "DEBUG":
		l = log.DebugLevel
	case "INFO":
		l = log.InfoLevel
	case "WARN":
		l = log.WarnLevel
	case "ERROR":
		l = log.ErrorLevel
	case "FATAL":
		l = log.FatalLevel
	}

	log.SetLevel(l)
	log.Info("Set debug level to", "level", *logLevel)
}

func main() {
	token := fmt.Sprintf("%s=%s", *apiUser, *apiKey)
	exporter.Exporter(*apiURL, token)
}
