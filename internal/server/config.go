package server

import (
	"flag"
	"os"
)

const hostEnvName = "ADDRESS"

var (
	host   string
	parsed bool = false
)

func parseFlags() {
	parsed = true
	u := os.Getenv(hostEnvName)

	if u == "" {
		flag.StringVar(&host, "a", "localhost:8080", "server address")
		flag.Parse()
		return
	}

	host = u
}

type Config struct {
	Host string
}

func NewConfig() *Config {
	if !parsed {
		parseFlags()
	}
	return &Config{Host: host}
}
