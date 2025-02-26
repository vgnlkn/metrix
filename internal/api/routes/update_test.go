package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vgnlkn/metrix/internal/metrix"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestUpdateHandler(t *testing.T) {
	c := chi.NewRouter()
	storage := metrix.NewMemStorage()
	ur := NewUpdateRouter(&storage)
	c.Route(`/`, func(r chi.Router) {
		ur.Route(r)
	})
	ts := httptest.NewServer(c)
	defer ts.Close()

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
			request: "/gauge/1/3.1415",
			want: want{
				code:        http.StatusOK,
				response:    "",
				contentType: "text/plain",
			},
		},
		{
			name:    "Correct counter metrics",
			request: "/counter/2/100",
			want: want{
				code:        http.StatusOK,
				response:    "",
				contentType: "text/plain",
			},
		},
		{
			name:    "Invalid metrics type",
			request: "/int/3/100",
			want: want{
				code:        http.StatusBadRequest,
				response:    "error: invalid metrics type\n\r\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Missing metrics name",
			request: "/gauge/",
			want: want{
				code:        http.StatusNotFound,
				response:    "404 page not found\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Missing metrics value",
			request: "/gauge/SomeValue",
			want: want{
				code:        http.StatusNotFound,
				response:    "404 page not found\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Invalid counter metrics value",
			request: "/counter/SomeValue/t",
			want: want{
				code:        http.StatusBadRequest,
				response:    "error: invalid metrics value\n\r\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Invalid gauge metrics value",
			request: "/counter/SomeValue/t",
			want: want{
				code:        http.StatusBadRequest,
				response:    "error: invalid metrics value\n\r\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, get := testRequest(t, ts, "POST", test.request)
			assert.Equal(t, test.want.code, resp.StatusCode)
			assert.Equal(t, test.want.response, get)
		})
	}
}
