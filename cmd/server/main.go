package main

import (
	"log"
	"net/http"

	"github.com/vgnlkn/metrix/internal/repository/memstorage"
	"github.com/vgnlkn/metrix/internal/router"
	"github.com/vgnlkn/metrix/internal/server"
	"github.com/vgnlkn/metrix/internal/usecase"
)

func main() {
	host := server.NewConfig().Host
	log.Println("Server running on:", host)

	memstorage := memstorage.NewMemStorage()
	usecase := usecase.NewMetricsUsecase(memstorage)
	router := router.NewRouter(usecase)
	http.ListenAndServe(host, router.Mux)
}
