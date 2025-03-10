package main

import (
	"log"
	"time"

	"github.com/vgnlkn/metrix/internal/client"
	"github.com/vgnlkn/metrix/internal/entity"
)

func main() {
	config := client.NewConfig()
	gaugeMetrics := make(entity.GaugeMetrics)
	counterMetrics := make(entity.CounterMetrics)
	client := client.NewClient("http://" + config.Host)

	lastReport := time.Now()

	log.Println("Launch params:")
	log.Printf(" - Host: %s\r\n", config.Host)
	log.Printf(" - Poll interval: %d sec\r\n", config.PollInterval)
	log.Printf(" - Report interval: %d sec\r\n", config.ReportInterval)

	for {
		if err := entity.CollectMetrics(&gaugeMetrics, &counterMetrics); err != nil {
			panic(err.Error())
		}
		log.Println(gaugeMetrics, counterMetrics)
		time.Sleep(time.Duration(config.PollInterval) * time.Second)

		now := time.Now()
		if now.Sub(lastReport) >= time.Duration(config.ReportInterval)*time.Second {
			client.UpdateMetrics(gaugeMetrics, counterMetrics)
			lastReport = time.Now()
		}
	}

}
