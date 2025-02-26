package metrix

import (
	"fmt"
)

type MetricsAsString struct {
	Name, Value, Type string
}

type Storage interface {
	Update(name, value, vType string) error
	Create(name, value, vType string) error
	GetAll() []MetricsAsString
	Find(name, vType string) (string, error)
}

type MemStorage struct {
	data []Metrics
}

var _ Storage = &MemStorage{}

func NewMemStorage() MemStorage {
	return MemStorage{
		data: make([]Metrics, 0),
	}
}

func (m *MemStorage) Update(name, value, vType string) error {
	for _, n := range m.data {
		if n.Name == name && n.vType == vType {
			return n.Val.Update(value)
		}
	}

	return m.Create(name, value, vType)
}

func (m *MemStorage) Create(name, value, vType string) error {
	new, err := NewMetrics(name, value, vType)
	if err != nil {
		return err
	}
	m.data = append(m.data, new)
	return nil
}

func (m *MemStorage) GetAll() []MetricsAsString {
	l := make([]MetricsAsString, 0, len(m.data))
	for _, v := range m.data {
		l = append(l, MetricsAsString{
			Name:  v.Name,
			Value: v.Val.String(),
			Type:  v.vType,
		})
	}
	return l
}

func (m *MemStorage) Find(name, vType string) (string, error) {
	for _, n := range m.data {
		if n.Name == name && n.vType == vType {
			return n.Val.String(), nil
		}
	}
	return "", fmt.Errorf("there'is no metrics with name %s and type %s", name, vType)
}

func (m *MemStorage) PrintStorage() {
	for _, v := range m.data {
		fmt.Println(v.String())
	}
}
