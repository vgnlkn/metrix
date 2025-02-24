package metrix

import (
	"strconv"
)

type Value interface {
	Update(string) error
	String() string
}

type GaugeValue struct {
	value Gauge
}

func NewGaugeValue(value string) *GaugeValue {
	v := new(GaugeValue)
	if v.Update(value) != nil {
		return nil
	}
	return v
}

func (v *GaugeValue) Update(new string) error {
	floatValue, err := strconv.ParseFloat(new, 64)
	if err != nil {
		return err
	}
	v.value = Gauge(floatValue)
	return nil
}

func (v *GaugeValue) String() string {
	return strconv.FormatFloat(float64(v.value), 'f', -1, 64)
}

type CounterValue struct {
	value Counter
}

func NewCounterValue(value string) *CounterValue {
	v := CounterValue{Counter(0)}
	if v.Update(value) != nil {
		return nil
	}

	return &v
}

func (v *CounterValue) Update(new string) error {
	intValue, err := strconv.ParseInt(new, 10, 64)
	if err != nil {
		return err
	}
	v.value += Counter(intValue)
	return nil
}

func (v CounterValue) String() string {
	return strconv.FormatInt(int64(v.value), 10)
}
