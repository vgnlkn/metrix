package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vgnlkn/metrix/internal/entity"
	"github.com/vgnlkn/metrix/internal/usecase"
	"github.com/vgnlkn/metrix/internal/usecase/mocks"
	"go.uber.org/mock/gomock"
)

func TestRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMetricsRepository(ctrl)
	mockRepo.EXPECT().FindMetrics("test", entity.TypeGauge).Return("1.3", nil)
	mockRepo.EXPECT().UpdateMetrics(gomock.Any()).Return(nil)

	usecase := usecase.NewMetricsUsecase(mockRepo)
	router := NewRouter(usecase)

	ts := httptest.NewServer(router.Mux)
	defer ts.Close()

	testRequest := func(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
		req, err := http.NewRequest(method, ts.URL+path, body)
		if err != nil {
			t.Fatal(err)
			return nil, ""
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
			return nil, ""
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
			return nil, ""
		}
		defer resp.Body.Close()

		return resp, string(respBody)
	}

	if _, body := testRequest(t, ts, "GET", "/value/gauge/test", nil); body != "1.3" {
		t.Errorf("expected 1.3, got %s", body)
	}

	if resp, _ := testRequest(t, ts, "POST", "/update/gauge/test/1.3", nil); resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	if resp, _ := testRequest(t, ts, "POST", "/update/gauge/test/invalid_value", nil); resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code 400, got %d", resp.StatusCode)
	}

	if resp, _ := testRequest(t, ts, "POST", "/update/invalid_type/test/1", nil); resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code 400, got %d", resp.StatusCode)
	}
}
