package entity

import (
	"errors"
	"fmt"
	"reflect"
)

type Metrics struct {
	Name string
	Val  Value
	Type string
}

func NewMetrics(name, value, vType string) (Metrics, error) {
	var v Value = nil

	switch vType {
	case TypeGauge:
		v = NewGaugeValue(value)
	case TypeCounter:
		v = NewCounterValue(value)
	default:
		return Metrics{}, errors.New("invalid metrics type")
	}

	if !reflect.ValueOf(v).IsNil() {
		return Metrics{name, v, vType}, nil
	}

	return Metrics{}, errors.New("invalid metrics value")
}

func (m Metrics) String() string {
	return fmt.Sprintf("%s : %v %v", m.Name, m.Val.String(), m.Type)
}
