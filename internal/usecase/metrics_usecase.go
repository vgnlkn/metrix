package usecase

import (
	"fmt"

	"github.com/vgnlkn/metrix/internal/entity"
)

type MetricsAsString struct {
	Name, Value, Type string
}

func NewMetricsAsString(m entity.Metrics) MetricsAsString {
	return MetricsAsString{
		Name:  m.Name,
		Value: m.Val.String(),
		Type:  m.Type,
	}
}

//go:generate mockgen -destination=mocks/mock_metricsrepo.go -package=mocks mypackage MyService
type MetricsRepository interface {
	UpdateMetrics(metric *entity.Metrics) error
	CreateMetrics(metric *entity.Metrics) error
	FindMetrics(name, vType string) (string, error)
	All() []entity.Metrics
}

type MetricsUsecase struct {
	storage MetricsRepository
}

func NewMetricsUsecase(storage MetricsRepository) *MetricsUsecase {
	return &MetricsUsecase{storage: storage}
}

func (u *MetricsUsecase) Update(name, value, vType string) error {
	metric, err := entity.NewMetrics(name, value, vType)
	if err != nil {
		return fmt.Errorf("entity.NewMetrics: %w", err)
	}
	if err = u.storage.UpdateMetrics(&metric); err != nil {
		return fmt.Errorf("storage.Update: %w", err)
	}
	return nil
}

func (u *MetricsUsecase) Find(name, vType string) (string, error) {
	return u.storage.FindMetrics(name, vType)
}

func (u *MetricsUsecase) All() []MetricsAsString {
	all := u.storage.All()
	r := make([]MetricsAsString, 0, len(all))
	for _, metric := range all {
		r = append(r, NewMetricsAsString(metric))
	}
	return r
}
