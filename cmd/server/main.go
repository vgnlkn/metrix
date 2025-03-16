package main

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/vgnlkn/metrix/internal/repository/memstorage"
	"github.com/vgnlkn/metrix/internal/router"
	"github.com/vgnlkn/metrix/internal/server"
	"github.com/vgnlkn/metrix/internal/usecase"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("cannot init logger")
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	host := server.NewConfig().Host
	sugar.Infof("Server running on: %s", host)

	memstorage := memstorage.NewMemStorage()
	usecase := usecase.NewMetricsUsecase(memstorage)
	router := router.NewRouter(usecase, logger)
	http.ListenAndServe(host, router.Mux)
}
