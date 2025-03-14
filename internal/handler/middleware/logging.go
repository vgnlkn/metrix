package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type contextKey string

const loggerKey contextKey = "logger"

var _ ResponseWriterProxy = ResponseWriterProxy{}

type ResponseWriterProxy struct {
	writer       http.ResponseWriter
	StatusCode   int
	ResponseSize int
}

func NewResponseWriterProxy(w http.ResponseWriter) *ResponseWriterProxy {
	return &ResponseWriterProxy{
		writer:       w,
		StatusCode:   -1,
		ResponseSize: 0,
	}
}

func (rw *ResponseWriterProxy) Header() http.Header {
	return rw.writer.Header()
}

func (rw *ResponseWriterProxy) Write(data []byte) (int, error) {
	rw.ResponseSize += len(data)
	return rw.writer.Write(data)
}

func (rw *ResponseWriterProxy) WriteHeader(statusCode int) {
	if rw.StatusCode == -1 {
		rw.writer.WriteHeader(statusCode)
		rw.StatusCode = statusCode
	}
}

func LoggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rwProxy := NewResponseWriterProxy(w)

			ctx := context.WithValue(r.Context(), loggerKey, logger)
			start := time.Now()
			next.ServeHTTP(rwProxy, r.WithContext(ctx))
			duration := time.Since(start).Microseconds()

			routeCtx := chi.RouteContext(r.Context())
			if routeCtx == nil {
				logger.Sugar().Errorf("Invalid routeCtx")
				return
			}

			endpoint := routeCtx.RoutePath

			logger.Info("Request received",
				zap.String("Endpoint", endpoint),
				zap.Int("Duration", int(duration)),
				zap.Int("StatusCode", rwProxy.StatusCode),
				zap.Int("ResponseSize", rwProxy.ResponseSize),
			)
		})
	}
}
