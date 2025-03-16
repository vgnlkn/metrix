package router

import (
	"github.com/vgnlkn/metrix/internal/handler"
	"github.com/vgnlkn/metrix/internal/handler/middleware"
	"github.com/vgnlkn/metrix/internal/usecase"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	Mux     *chi.Mux
	handler handler.Handlers
	logger  *zap.Logger
}

func NewRouter(mUsecase *usecase.MetricsUsecase, logger *zap.Logger) Service {
	s := Service{
		chi.NewRouter(),
		handler.NewHandlers(mUsecase),
		logger,
	}

	if s.logger != nil {
		s.Mux.Use(middleware.LoggerMiddleware(logger))
	}

	s.Mux.Route(`/`, func(r chi.Router) {
		r.Get(`/`, s.handler.Home)
		r.Route(`/value`, func(r chi.Router) {
			r.Post(`/`, s.handler.GetMetricValueViaJSON)
			r.Get(`/{type}/{name}`, s.handler.GetMetricValueViaURL)
		})
		r.Route(`/update`, func(r chi.Router) {
			r.Post(`/`, s.handler.UpdateMetricsViaJSON)
			r.Post(`/{type}/{name}/{value}`, s.handler.UpdateMetricsViaURL)
		})
	})

	return s
}
