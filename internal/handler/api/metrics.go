package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/vgnlkn/metrix/internal/entity"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func NewMetricsFromString(name, vType, value string) (Metrics, error) {
	m, err := entity.NewMetrics(name, value, vType)
	if err != nil {
		return Metrics{}, err
	}

	apiMetrics := Metrics{
		ID:    m.Name,
		MType: m.Type,
	}

	if gv, ok := m.Val.(*entity.GaugeValue); ok {
		apiMetrics.Value = new(float64)
		*apiMetrics.Value = gv.AsNumber()
	} else if cv, ok := m.Val.(*entity.CounterValue); ok {
		apiMetrics.Delta = new(int64)
		*apiMetrics.Delta = cv.AsNumber()
	}

	return apiMetrics, nil
}

func NewMetricsFromStream(stream io.ReadCloser) *Metrics {
	m := Metrics{}
	dec := json.NewDecoder(stream)
	if err := dec.Decode(&m); err != nil {
		return nil
	}

	return &m
}

func (m *Metrics) ToString() (name, vType, value string) {
	if m.Value != nil {
		value = fmt.Sprintf("%f", *m.Value)
	} else if m.Delta != nil {
		value = fmt.Sprintf("%d", *m.Delta)
	}

	name, vType = m.ID, m.MType
	return
}

func (m *Metrics) ToBytes() *[]byte {
	if byteArray, err := json.Marshal(m); err == nil {
		return &byteArray
	}
	return nil
}
