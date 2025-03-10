package main

import (
	"fmt"
	"net/http"

	"github.com/vgnlkn/metrix/internal/repository/memstorage"
	"github.com/vgnlkn/metrix/internal/router"
	"github.com/vgnlkn/metrix/internal/usecase"
)

func main() {
	//flag.Parse()
	parseFlags()
	fmt.Println("Server running on:", host)

	memstorage := memstorage.NewMemStorage()
	usecase := usecase.NewMetricsUsecase(memstorage)
	router := router.NewRouter(usecase)
	http.ListenAndServe(host, router.Mux)
}
