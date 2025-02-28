package main

import (
	"fmt"

	"github.com/vgnlkn/metrix/internal/api/server"
)

func main() {
	//flag.Parse()
	parseFlags()
	fmt.Println("Server running on:", host)

	server := server.NewServer(host)
	server.Run()
}
