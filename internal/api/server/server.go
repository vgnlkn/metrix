package server

import (
	"net/http"

	"github.com/vgnlkn/metrix/internal/api/server/routes"
	"github.com/vgnlkn/metrix/internal/metrix"
)

const (
	home string = "/"
)

type Server struct {
	mux     *http.ServeMux
	storage metrix.MemStorage
}

func (s *Server) Run() {
	http.ListenAndServe(`localhost:8080`, s.mux)
}

func NewServer() Server {
	s := Server{
		http.NewServeMux(),
		metrix.NewMemStorage(),
	}

	p, u := routes.NewUpdateRoute(&s.storage)
	s.mux.Handle(p, u)
	s.mux.Handle(home, http.NotFoundHandler())
	return s
}
