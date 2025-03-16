package memstorage

import (
	"fmt"

	"github.com/vgnlkn/metrix/internal/entity"
	"github.com/vgnlkn/metrix/internal/usecase"
)

type stored struct {
	Val  entity.Value
	Type string
}

type storage map[string]stored

var _ usecase.MetricsRepository = &MemStorage{}

type MemStorage struct {
	gaugeMetrics   storage
	counterMetrics storage
}

func (m MemStorage) getMap(vType string) (*storage, error) {
	switch vType {
	case entity.TypeCounter:
		return &m.counterMetrics, nil
	case entity.TypeGauge:
		return &m.gaugeMetrics, nil
	default:
		return nil, fmt.Errorf("invalid metrics type")
	}
}

func NewMemStorage() MemStorage {
	return MemStorage{
		gaugeMetrics:   make(storage, 0),
		counterMetrics: make(storage, 0),
	}
}

func (m MemStorage) UpdateMetrics(metric *entity.Metrics) error {
	storage, err := m.getMap(metric.Type)
	if err != nil {
		return err
	}

	if _, exists := (*storage)[metric.Name]; !exists {
		return m.CreateMetrics(metric)
	}

	return (*storage)[metric.Name].Val.Update(metric.Val.String())
}

func (m MemStorage) CreateMetrics(metric *entity.Metrics) error {
	storage, err := m.getMap(metric.Type)
	if err != nil {
		return err
	}

	if _, exists := (*storage)[metric.Name]; exists {
		return fmt.Errorf("MemStorage.CreateMetrics: metrics exist")
	}
	(*storage)[metric.Name] = stored{
		Val:  metric.Val,
		Type: metric.Type,
	}

	return nil
}

func (m MemStorage) FindMetrics(name, vType string) (string, error) {
	storage, err := m.getMap(vType)
	if err != nil {
		return "", err
	}

	if value, exists := (*storage)[name]; exists {
		return value.Val.String(), nil
	}
	return "", fmt.Errorf("MemStorage.FindMetrics: not found %s", name)
}

func (m MemStorage) All() []entity.Metrics {
	list := make([]entity.Metrics, 0, len(m.gaugeMetrics)+len(m.counterMetrics))

	allOf := func(storage map[string]stored) {
		for k, v := range storage {
			list = append(list, entity.Metrics{
				Name: k,
				Val:  v.Val,
				Type: v.Type,
			})
		}
	}

	allOf(m.gaugeMetrics)
	allOf(m.counterMetrics)

	return list
}
