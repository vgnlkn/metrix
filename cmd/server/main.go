package main

import (
	"flag"
	"fmt"

	"github.com/vgnlkn/metrix/internal/api/server"
)

func main() {
	flag.Parse()

	fmt.Println("Server running on:", *host)

	server := server.NewServer(*host)
	server.Run()
}
