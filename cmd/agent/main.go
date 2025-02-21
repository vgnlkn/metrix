package main

import (
	"fmt"
	"time"

	"github.com/vgnlkn/metrix/internal/client"
	"github.com/vgnlkn/metrix/internal/metrix"
)

func main() {
	gm := make(metrix.GaugeMetrics)
	cm := make(metrix.CounterMetrics)
	client := client.NewClient("http://localhost:8080")

	for {
		if err := metrix.GrabMetrics(&gm, &cm); err != nil {
			panic(err.Error())
		}
		fmt.Println(gm, cm)
		fmt.Println("=====================")
		time.Sleep(2 * time.Second)

		client.UpdateMetrics(gm, cm)
	}

}
