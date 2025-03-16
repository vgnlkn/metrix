package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgnlkn/metrix/internal/entity"
	"github.com/vgnlkn/metrix/internal/usecase"
	"github.com/vgnlkn/metrix/internal/usecase/mocks"
	"go.uber.org/mock/gomock"
)

func TestHandlersHome(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMetricsRepository(ctrl)
	mockRepo.EXPECT().All().Return([]entity.Metrics{
		{Name: "test", Val: entity.NewGaugeValue("1.3"), Type: entity.TypeGauge},
		{Name: "counter", Val: entity.NewCounterValue("1234"), Type: entity.TypeGauge},
	})

	usecase := usecase.NewMetricsUsecase(mockRepo)
	handler := NewHandlers(usecase)

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Home(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "test")
	assert.Contains(t, rr.Body.String(), "counter")
	assert.Contains(t, rr.Body.String(), "1.3")
	assert.Contains(t, rr.Body.String(), "1234")
}
