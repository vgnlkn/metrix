package main

import (
	"flag"
	"time"
)

const (
	DefaultPollInterval   = 2
	DefaultReportInterval = 10
)

var host *string = flag.String("a", "localhost:8080", "server address")
var pollInterval *int = flag.Int("p", DefaultPollInterval, "metrics poll interval")
var reportInterval *int = flag.Int("r", DefaultReportInterval, "metrics report interval")

func GetPollInt() time.Duration {
	return time.Duration(*pollInterval) * time.Second
}

func GetReportInt() time.Duration {
	return time.Duration(*reportInterval) * time.Second
}
