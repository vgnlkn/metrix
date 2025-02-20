package metrix

const (
	TypeGauge   string = "gauge"
	TypeCounter string = "counter"
)

type Gauge float64
type Counter int64
