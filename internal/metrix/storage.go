package metrix

import (
	"fmt"
)

type MemStorage struct {
	data []Metrics
}

func NewMemStorage() MemStorage {
	return MemStorage{}
}

func (m *MemStorage) Update(name, value, vType string) error {
	for _, n := range m.data {
		if n.Name != name {
			continue
		}
		if n.vType != vType {
			return fmt.Errorf("attempt to update value with type %s using %s update request", n.vType, vType)
		}
		return n.Val.Update(value)
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

func (m *MemStorage) PrintStorage() {
	for _, v := range m.data {
		fmt.Println(v.String())
	}
}
