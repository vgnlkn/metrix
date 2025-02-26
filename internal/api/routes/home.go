package routes

import (
	//"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vgnlkn/metrix/internal/metrix"
	//"github.com/go-chi/chi/v5"
)

type HomeHandler struct {
	storage metrix.Storage
}

type HomeRouter struct {
	handl HomeHandler
}

func (hr HomeRouter) Route(r chi.Router) {
	r.Get(`/`, hr.handl.ServeHTTP)
}

func NewHomeRouter(storage metrix.Storage) HomeRouter {
	return HomeRouter{NewHomeHandler(storage)}
}

type Metrics struct {
	Name string
	Val  string
}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	err = tmpl.ExecuteTemplate(w, "metrics", h.storage.GetAll())
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func NewHomeHandler(storage metrix.Storage) HomeHandler {
	return HomeHandler{storage}
}
