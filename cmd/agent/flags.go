package main

import (
	"flag"
	"os"
	"strconv"
	"time"
)

const (
	AddressEnvName   = "ADDRESS"
	ReportIntEnvName = "REPORT_INTERVAL"
	PollIntEnvName   = "POLL_INTERVAL"

	DefaultPollInterval   = 2
	DefaultReportInterval = 10
)

var (
	host           string
	pollInterval   int
	reportInterval int
)

func parseFlags() {
	flag.StringVar(&host, "a", "localhost:8080", "server address")
	flag.IntVar(&pollInterval, "p", DefaultPollInterval, "metrics poll interval")
	flag.IntVar(&reportInterval, "r", DefaultReportInterval, "metrics report interval")

	flag.Parse()

	if addr := os.Getenv(AddressEnvName); addr != "" {
		host = addr
	}
	if pollInt := os.Getenv(PollIntEnvName); pollInt != "" {
		num, err := strconv.Atoi(pollInt)
		if err != nil {
			panic("invalid pollInt value")
		}
		pollInterval = num
	}
	if reportInt := os.Getenv(ReportIntEnvName); reportInt != "" {
		num, err := strconv.Atoi(reportInt)
		if err != nil {
			panic("invalid reportInt value")
		}
		reportInterval = num
	}

}

func GetPollInt() time.Duration {
	return time.Duration(pollInterval) * time.Second
}

func GetReportInt() time.Duration {
	return time.Duration(reportInterval) * time.Second
}
