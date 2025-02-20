package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vgnlkn/metrix/internal/metrix"
)

type UpdateRoute struct {
	storage *metrix.MemStorage
}

func (ur UpdateRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	parts := strings.Split(r.URL.Path, "/")[2:]

	if l := len(parts); l != 3 {
		http.Error(w, "invalid request", http.StatusNotFound)
		return
	}

	metricsType := parts[0]
	metricsName := parts[1]
	metricsValue := parts[2]

	if err := ur.storage.Update(metricsName, metricsValue, metricsType); err != nil {
		http.Error(w, fmt.Sprintf("error: %v\n\r", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Println("======================")
	ur.storage.PrintStorage()

}

func NewUpdateRoute(storage *metrix.MemStorage) (pattern string, ur UpdateRoute) {
	return `/update/`, UpdateRoute{storage}
}
