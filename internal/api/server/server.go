package server

import (
	"net/http"

	"github.com/vgnlkn/metrix/internal/api/routes"
	"github.com/vgnlkn/metrix/internal/metrix"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	host    string
	mux     *chi.Mux
	storage metrix.MemStorage
}

func (s *Server) Run() {
	http.ListenAndServe(s.host, s.mux)
}

func NewServer(host string) Server {
	s := Server{
		host,
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
