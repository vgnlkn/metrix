package agent

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	AddressEnvName   = "ADDRESS"
	ReportIntEnvName = "REPORT_INTERVAL"
	PollIntEnvName   = "POLL_INTERVAL"

	DefaultPollInterval   = 2
	DefaultReportInterval = 10
)

type Config struct {
	Host           string
	PollInterval   int
	ReportInterval int
}

var (
	host           string
	pollInterval   int
	reportInterval int
	parsed         bool = false
)

func NewConfig() *Config {
	if !parsed {
		flag.StringVar(&host, "a", "localhost:8080", "server address")
		flag.IntVar(&pollInterval, "p", DefaultPollInterval, "metrics poll interval")
		flag.IntVar(&reportInterval, "r", DefaultReportInterval, "metrics report interval")

		flag.Parse()
		parsed = true
		fmt.Println(host, reportInterval, pollInterval)
	}
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

	return &Config{
		Host:           host,
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
	}
}
