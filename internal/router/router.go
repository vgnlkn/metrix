package router

import (
	"github.com/vgnlkn/metrix/internal/handler"
	"github.com/vgnlkn/metrix/internal/usecase"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Mux     *chi.Mux
	handler handler.Handlers
}

func NewRouter(mUsecase *usecase.MetricsUsecase) Router {
	s := Router{
		chi.NewRouter(),
		handler.NewHandlers(mUsecase),
	}

	s.Mux.Route(`/`, func(r chi.Router) {
		r.Get(`/`, s.handler.Home)
		r.Get(`/value/{type}/{name}`, s.handler.Value)
		r.Post(`/update/{type}/{name}/{value}`, s.handler.Update)
	})

	return s
}
