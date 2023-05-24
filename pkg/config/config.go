package config

import (
	"io"
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Clusters      []Cluster `mapstructure:"clusters"`
	FetchInterval int       `mapstructure:"fetchInterval"`
}

type Cluster struct {
	ApiKey  string `mapstructure:"apikey"`
	ApiUser string `mapstructure:"apiuser"`
	ApiHost string `mapstructure:"apihost"`
}

// variables to be parsed
var (
	filePath = kingpin.
			Flag("cluster-file", "API Key").
			Short('f').
			Required().
			Envar("CLUSTER_FILE").
			String()
	logLevel = kingpin.
			Flag("log-level", "Set the Log Level / verbosity").
			Short('l').
			Envar("LOG_LEVEL").
			Default("INFO").
			Enum("DEBUG", "INFO", "WARN", "ERROR", "FATAL")
	fetchInterval = kingpin.
			Flag("fetch-interval", "Interval at whicht to refetch all VMs").
			Short('i').
			Default("300").
			Int()
)

func init() {
	kingpin.Parse()

	s := strings.ToLower(*logLevel)

	l := log.ParseLevel(s)

	log.SetLevel(l)

	log.Info("log level", "level", log.GetLevel().String())
}

func New() *Config {
	c := &Config{
		FetchInterval: *fetchInterval,
	}

	f, err := os.Open(*filePath)
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Error("closing file failed", "err", err)
		}
	}(f)

	if err != nil {
		log.Fatal("could not open config file", "err", err)
	}

	r, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("cold not read content", "err", err)
	}

	var y interface{}
	if err := yaml.Unmarshal(r, &y); err != nil {
		log.Fatal("cold not read YAML", "err", err)
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &c})
	if err := decoder.Decode(y); err != nil {
		log.Fatal("error unmarshalling YAML", "err", err)
	}

	return c
}
