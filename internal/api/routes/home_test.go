package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/vgnlkn/metrix/internal/metrix"
)

func TestHomeHandler(t *testing.T) {
	c := chi.NewRouter()
	storage := metrix.NewMemStorage()
	hr := NewHomeRouter(&storage)
	c.Route(`/`, func(r chi.Router) {
		hr.Route(r)
	})
	ts := httptest.NewServer(c)
	defer ts.Close()

	type fields struct {
		storage metrix.Storage
	}

	tests := []struct {
		name     string
		request  string
		fields   fields
		wantCode int
	}{
		{
			name:    "Check if html generated",
			request: "/",
			fields: fields{
				storage: &storage,
			},
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, got := testRequest(t, ts, "GET", tt.request)
			assert.Equal(t, tt.wantCode, code)
			assert.Greater(t, len(got), 0)
		})
	}
}
