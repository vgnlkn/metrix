package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vgnlkn/metrix/internal/usecase"
)

type Handlers struct {
	mUsecase *usecase.MetricsUsecase
}

func NewHandlers(mUsecase *usecase.MetricsUsecase) Handlers {
	return Handlers{mUsecase: mUsecase}
}

func (h *Handlers) Update(rw http.ResponseWriter, r *http.Request) {
	metricsType := chi.URLParam(r, "type")
	metricsName := chi.URLParam(r, "name")
	metricsValue := chi.URLParam(r, "value")

	err := h.mUsecase.Update(metricsName, metricsValue, metricsType)

	if err != nil {
		http.Error(rw, fmt.Sprintf("error: %v\n\r", err.Error()), http.StatusBadRequest)
		return
	}

	rw.Header().Add("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
}

func (h *Handlers) Value(w http.ResponseWriter, r *http.Request) {
	metricsType := chi.URLParam(r, "type")
	metricsName := chi.URLParam(r, "name")

	v, err := h.mUsecase.Find(metricsName, metricsType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", v)
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	content := `{{define "metrics"}}
	<!DOCTYPE html class="h-100">
	<html lang="ru">
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
	<head>
		<meta charset="UTF-8">
		<title>Metrix</title>
	</head>
	<body class="d-flex flex-column h-100">
		<table class="table">
			<th>Тип</th>
			<th>Название метрики</th>
			<th>Значение</th>
			{{range .}}
			<tr>
				<td>{{.Type}}</td>
				<td>{{.Name}}</td>
				<td>{{.Value}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	{{end}}`

	tmpl, err := template.New("metrics").Parse(content)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = tmpl.ExecuteTemplate(w, "metrics", h.mUsecase.All())
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
