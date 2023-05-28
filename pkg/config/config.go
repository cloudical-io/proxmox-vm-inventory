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
	RequestTimeout    int       `mapstructure:"timeout"`
	FetchInterval     int       `mapstructure:"fetchInterval"`
	Clusters          []Cluster `mapstructure:"clusters"`
	HttpListenAddress string    `mapstructure:"httpAddress"`
}

type Cluster struct {
	Name    string `mapstructure:"name"`
	ApiKey  string `mapstructure:"apikey"`
	ApiUser string `mapstructure:"apiuser"`
	ApiHost string `mapstructure:"apihost"`
}

// variables to be parsed
var (
	filePath = kingpin.
			Flag("config-file", "YAML file containing your config values. Values set here override all commandline flags and environment vars").
			Short('f').
			Required().
			Envar("INV_CONFIG_FILE").
			String()
	logLevel = kingpin.
			Flag("log-level", "Set the Log Level / verbosity").
			Short('L').
			Envar("INV_LOG_LEVEL").
			Default("INFO").
			Enum("DEBUG", "INFO", "WARN", "ERROR", "FATAL")
	requestTimeout = kingpin.
			Flag("timeout", "Time in seconds before a request times out").
			Short('t').
			Default("10").
			Envar("INV_TIMEOUT").
			Int()
	fetchInterval = kingpin.
			Flag("fetch-interval", "Interval at whicht to refetch all VMs").
			Short('i').
			Default("300").
			Envar("INV_INTERVAL").
			Int()
	httpListenAddress = kingpin.
				Flag("listen-address", "The http port to listen on").
				HintOptions(":8080", "127.0.0.1:8080", "[::]:8080").
				Default(":8080").
				Short('l').
				Envar("INV_HTTP_LISTEN").
				String()
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
		FetchInterval:     *fetchInterval,
		RequestTimeout:    *requestTimeout,
		HttpListenAddress: *httpListenAddress,
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
