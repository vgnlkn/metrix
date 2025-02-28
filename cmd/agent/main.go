package main

import (
	"fmt"
	"time"

	"github.com/vgnlkn/metrix/internal/client"
	"github.com/vgnlkn/metrix/internal/metrix"
)

func main() {
	parseFlags()
	gaugeMetrics := make(metrix.GaugeMetrics)
	counterMetrics := make(metrix.CounterMetrics)
	client := client.NewClient("http://" + host)

	lastReport := time.Now()

	reportInt := GetReportInt()
	pollInt := GetPollInt()

	fmt.Println("Launch params:")
	fmt.Printf(" - Host: %s\r\n", host)
	fmt.Printf(" - Poll interval: %d sec\r\n", pollInt/time.Second)
	fmt.Printf(" - Report interval: %d sec\r\n", reportInt/time.Second)
	fmt.Println("=====================")

	for {
		if err := metrix.GrabMetrics(&gaugeMetrics, &counterMetrics); err != nil {
			panic(err.Error())
		}
		fmt.Println(gaugeMetrics, counterMetrics)
		fmt.Println("=====================")
		time.Sleep(pollInt)

		now := time.Now()
		if now.Sub(lastReport) >= reportInt {
			client.UpdateMetrics(gaugeMetrics, counterMetrics)
			lastReport = time.Now()
		}
	}

}
