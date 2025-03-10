package entity

const (
	TypeGauge   string = "gauge"
	TypeCounter string = "counter"
)

const (
	/*	Some gauge metrics */
	Alloc         = "Alloc"
	BuckHashSize  = "BuckHashSize"
	Frees         = "Frees"
	GCCPUFraction = "GCCPUFraction"
	GCSys         = "GCSys"
	HeapAlloc     = "HeapAlloc"
	HeapIdle      = "HeapIdle"
	HeapInuse     = "HeapInuse"
	HeapObjects   = "HeapObjects"
	HeapReleased  = "HeapReleased"
	HeapSys       = "HeapSys"
	LastGC        = "LastGC"
	Lookups       = "Lookups"
	MCacheInuse   = "MCacheInuse"
	MCacheSys     = "MCacheSys"
	MSpanInuse    = "MSpanInuse"
	MSpanSys      = "MSpanSys"
	Mallocs       = "Mallocs"
	NextGC        = "NextGC"
	NumForcedGC   = "NumForcedGC"
	NumGC         = "NumGC"
	OtherSys      = "OtherSys"
	PauseTotalNs  = "PauseTotalNs"
	StackInuse    = "StackInuse"
	StackSys      = "StackSys"
	Sys           = "Sys"
	TotalAlloc    = "TotalAlloc"
	RandomValue   = "RandomValue"

	/* Some counter Metrics	*/
	PollCount = "PollCount"
)

const (
	RandValMax float64 = 5.0
	RandValMin float64 = -5.0
)

type Gauge float64
type Counter int64

type GaugeMetrics map[string]Gauge
type CounterMetrics map[string]Counter
