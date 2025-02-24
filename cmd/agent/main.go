package main

import (
	"fmt"
	"time"

	"github.com/vgnlkn/metrix/internal/client"
	"github.com/vgnlkn/metrix/internal/metrix"
)

const (
	DefaultPollInterval   = 2 * time.Second
	DefaultReportInterval = 10 * time.Second
)

func main() {
	gm := make(metrix.GaugeMetrics)
	cm := make(metrix.CounterMetrics)
	client := client.NewClient("http://localhost:8080")

	lastReport := time.Now()

	for {
		if err := metrix.GrabMetrics(&gm, &cm); err != nil {
			panic(err.Error())
		}
		fmt.Println(gm, cm)
		fmt.Println("=====================")
		time.Sleep(2 * time.Second)

		now := time.Now()
		if now.Sub(lastReport) >= DefaultReportInterval {
			client.UpdateMetrics(gm, cm)
			lastReport = time.Now()
		}
	}

}
