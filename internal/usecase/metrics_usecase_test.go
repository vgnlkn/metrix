package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vgnlkn/metrix/internal/entity"
	"github.com/vgnlkn/metrix/internal/usecase/mocks"
	"go.uber.org/mock/gomock"
)

func TestUpdateMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMetricsRepository(ctrl)
	mockRepo.EXPECT().UpdateMetrics(gomock.Any()).Return(nil)

	usecase := NewMetricsUsecase(mockRepo)
	err := usecase.Update("test", "123", "gauge")
	assert.NoError(t, err)
}

func TestFindMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMetricsRepository(ctrl)
	mockRepo.EXPECT().FindMetrics("test", "gauge").Return("123", nil)

	usecase := NewMetricsUsecase(mockRepo)
	value, err := usecase.Find("test", "gauge")
	assert.NoError(t, err)
	assert.Equal(t, "123", value)
}

func TestAllMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMetricsRepository(ctrl)
	mockRepo.EXPECT().All().Return([]entity.Metrics{
		{Name: "test1", Val: entity.NewGaugeValue("123"), Type: "gauge"},
		{Name: "test2", Val: entity.NewCounterValue("123"), Type: "counter"},
	})

	usecase := NewMetricsUsecase(mockRepo)
	allMetrics := usecase.All()
	assert.Len(t, allMetrics, 2)
	assert.Equal(t, "test1", allMetrics[0].Name)
	assert.Equal(t, "test2", allMetrics[1].Name)
}
