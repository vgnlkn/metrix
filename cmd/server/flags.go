package main

import (
	"flag"
)

var host *string = flag.String("a", "localhost:8080", "server address")
