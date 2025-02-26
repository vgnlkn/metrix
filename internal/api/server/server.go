package server

import (
	"net/http"

	"github.com/vgnlkn/metrix/internal/api/routes"
	"github.com/vgnlkn/metrix/internal/metrix"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	mux     *chi.Mux
	storage metrix.MemStorage
}

func (s *Server) Run() {
	http.ListenAndServe(`localhost:8080`, s.mux)
}

func NewServer() Server {
	s := Server{
		chi.NewRouter(),
		metrix.NewMemStorage(),
	}

	ur := routes.NewUpdateRouter(&s.storage)
	hr := routes.NewHomeRouter(&s.storage)
	vr := routes.NewValueRouter(&s.storage)

	s.mux.Route(`/`, func(r chi.Router) {
		r.Route(`/`, hr.Route)
		r.Route(`/update`, ur.Route)
		r.Route(`/value`, vr.Route)
	})

	return s
}
