package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vgnlkn/metrix/internal/metrix"
)

func TestUpdateHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "Correct gauge metrics",
			request: "gauge/SomeValue/3.1415",
			want: want{
				code:        http.StatusOK,
				response:    "",
				contentType: "text/plain",
			},
		},
		{
			name:    "Correct counter metrics",
			request: "counter/SomeValue/100",
			want: want{
				code:        http.StatusOK,
				response:    "",
				contentType: "text/plain",
			},
		},
		{
			name:    "Invalid metrics type",
			request: "int/SomeValue/100",
			want: want{
				code:        http.StatusBadRequest,
				response:    "error: invalid metrics type\n\r\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Missing metrics name",
			request: "gauge/",
			want: want{
				code:        http.StatusNotFound,
				response:    "invalid request\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Missing metrics value",
			request: "gauge/SomeValue",
			want: want{
				code:        http.StatusNotFound,
				response:    "invalid request\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Invalid counter metrics value",
			request: "counter/SomeValue/t",
			want: want{
				code:        http.StatusBadRequest,
				response:    "error: invalid metrics value\n\r\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Invalid gauge metrics value",
			request: "counter/SomeValue/t",
			want: want{
				code:        http.StatusBadRequest,
				response:    "error: invalid metrics value\n\r\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage := metrix.NewMemStorage()
			pattern, ur := NewUpdateRoute(&storage)
			request := httptest.NewRequest(http.MethodPost, pattern+test.request, nil)
			w := httptest.NewRecorder()
			ur.ServeHTTP(w, request)

			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, test.want.response, string(resBody))
		})
	}
}
