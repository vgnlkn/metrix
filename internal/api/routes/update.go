package routes

import (
	"fmt"
	"net/http"

	"github.com/vgnlkn/metrix/internal/metrix"

	"github.com/go-chi/chi/v5"
)

type UpdateRoute struct {
	storage metrix.Storage
}

type UpdateRouter struct {
	updateRouter UpdateRoute
}

func (ur UpdateRouter) Route(r chi.Router) {
	r.Route(`/{type}/{name}/{value}`, func(r chi.Router) {
		r.Post(`/`, ur.updateRouter.ServeHTTP)
	})
}

func NewUpdateRouter(storage metrix.Storage) UpdateRouter {
	u := NewUpdateRoute(storage)
	return UpdateRouter{u}
}

func (ur UpdateRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	metricsType := chi.URLParam(r, "type")
	metricsName := chi.URLParam(r, "name")
	metricsValue := chi.URLParam(r, "value")

	if err := ur.storage.Update(metricsName, metricsValue, metricsType); err != nil {
		http.Error(w, fmt.Sprintf("error: %v\n\r", err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

}

func NewUpdateRoute(storage metrix.Storage) UpdateRoute {
	return UpdateRoute{storage}
}
