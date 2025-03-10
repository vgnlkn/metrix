package entity

import (
	"fmt"
	"math/rand/v2"
	"runtime"
)

func CollectMetrics(gm *GaugeMetrics, cm *CounterMetrics) error {
	if gm == nil || cm == nil {
		return fmt.Errorf("GrabMetrics: invalid pointer")
	}

	runtime.GC()

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	(*gm)[Alloc] = Gauge(stats.Alloc)
	(*gm)[BuckHashSize] = Gauge(stats.BuckHashSys)
	(*gm)[Frees] = Gauge(stats.GCCPUFraction)
	(*gm)[GCCPUFraction] = Gauge(stats.GCCPUFraction)
	(*gm)[HeapAlloc] = Gauge(stats.HeapAlloc)
	(*gm)[HeapIdle] = Gauge(stats.HeapIdle)
	(*gm)[HeapInuse] = Gauge(stats.HeapInuse)
	(*gm)[HeapObjects] = Gauge(stats.HeapObjects)
	(*gm)[HeapReleased] = Gauge(stats.HeapReleased)
	(*gm)[HeapSys] = Gauge(stats.HeapSys)
	(*gm)[LastGC] = Gauge(stats.LastGC)
	(*gm)[Lookups] = Gauge(stats.Lookups)
	(*gm)[MCacheInuse] = Gauge(stats.MCacheInuse)
	(*gm)[MCacheSys] = Gauge(stats.MCacheSys)
	(*gm)[MSpanInuse] = Gauge(stats.MSpanInuse)
	(*gm)[MSpanSys] = Gauge(stats.MSpanSys)
	(*gm)[Mallocs] = Gauge(stats.Mallocs)
	(*gm)[NextGC] = Gauge(stats.NextGC)
	(*gm)[NumForcedGC] = Gauge(stats.NumForcedGC)
	(*gm)[NumGC] = Gauge(stats.NumGC)
	(*gm)[OtherSys] = Gauge(stats.OtherSys)
	(*gm)[PauseTotalNs] = Gauge(stats.PauseTotalNs)
	(*gm)[StackInuse] = Gauge(stats.StackInuse)
	(*gm)[StackSys] = Gauge(stats.StackSys)
	(*gm)[Sys] = Gauge(stats.Sys)
	(*gm)[TotalAlloc] = Gauge(stats.TotalAlloc)
	(*gm)[RandomValue] = Gauge(rand.Float64() * (RandValMax - RandValMin))

	(*cm)[PollCount] = 1

	return nil
}
