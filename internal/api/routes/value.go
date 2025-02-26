package routes

import (
	"fmt"
	"net/http"

	"github.com/vgnlkn/metrix/internal/metrix"

	"github.com/go-chi/chi/v5"
)

type ValueHandler struct {
	storage metrix.Storage
}

type ValueRouter struct {
	handl ValueHandler
}

func (vr ValueRouter) Route(r chi.Router) {
	r.Route(`/{type}/{name}`, func(r chi.Router) {
		r.Get(`/`, vr.handl.ServeHTTP)
	})
}

func NewValueRouter(storage metrix.Storage) ValueRouter {
	return ValueRouter{NewValueHandler(storage)}
}

func (vh ValueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	metricsType := chi.URLParam(r, "type")
	metricsName := chi.URLParam(r, "name")

	v, err := vh.storage.Find(metricsName, metricsType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", v)
}

func NewValueHandler(storage metrix.Storage) ValueHandler {
	return ValueHandler{storage}
}
