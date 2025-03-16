package main

import (
	"log"
	"time"

	"github.com/vgnlkn/metrix/internal/agent"
	"github.com/vgnlkn/metrix/internal/entity"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("cannot init logger")
	}
	defer logger.Sync()

	config := agent.NewConfig()
	gaugeMetrics := make(entity.GaugeMetrics)
	counterMetrics := make(entity.CounterMetrics)
	client := agent.NewClient("http://"+config.Host, logger)

	lastReport := time.Now()

	log.Println("Launch params:")
	log.Printf(" - Host: %s\r\n", config.Host)
	log.Printf(" - Poll interval: %d sec\r\n", config.PollInterval)
	log.Printf(" - Report interval: %d sec\r\n", config.ReportInterval)

	for {
		if err := entity.CollectMetrics(&gaugeMetrics, &counterMetrics); err != nil {
			panic(err.Error())
		}
		time.Sleep(time.Duration(config.PollInterval) * time.Second)

		now := time.Now()
		if now.Sub(lastReport) >= time.Duration(config.ReportInterval)*time.Second {
			client.UpdateMetrics(gaugeMetrics, counterMetrics)
			lastReport = time.Now()
		}
	}

}
