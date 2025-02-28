package main

import (
	"flag"
	"os"
)

const hostEnvName = "ADDRESS"

var host string

func parseFlags() {
	u := os.Getenv(hostEnvName)

	if u == "" {
		flag.StringVar(&host, "a", "localhost:8080", "server address")
		flag.Parse()
		return
	}

	host = u
}
