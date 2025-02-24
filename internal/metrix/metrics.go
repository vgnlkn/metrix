package metrix

import (
	"errors"
	"fmt"
	"reflect"
)

type Metrics struct {
	Name  string
	Val   Value
	vType string
}

func NewMetrics(name, value, vType string) (Metrics, error) {
	var v Value = nil

	if vType == TypeGauge {
		v = NewGaugeValue(value)
	} else if vType == TypeCounter {
		v = NewCounterValue(value)
	} else {
		return Metrics{}, errors.New("invalid metrics type")
	}

	if !reflect.ValueOf(v).IsNil() {
		return Metrics{name, v, vType}, nil
	}

	return Metrics{}, errors.New("invalid metrics value")
}

func (m Metrics) String() string {
	return fmt.Sprintf("%s : %v", m.Name, m.Val.String())
}
