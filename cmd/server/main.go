package main

import (
	"github.com/vgnlkn/metrix/internal/api/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
