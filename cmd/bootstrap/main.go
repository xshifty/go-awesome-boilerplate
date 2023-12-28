package main

import (
	"flag"
	"fmt"
	"github.com/xshifty/go-awesome-boilerplate/handlers"
	"net/http"
)

func main() {
	var host string
	var port int

	flag.StringVar(&host, "h", "localhost", "Listening host (Default: localhost)")
	flag.IntVar(&port, "p", 8000, "Listening port (Default: 8000)")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, port)
	router := handlers.Register(
		handlers.WithLoggerEnabled(true),
	)

	fmt.Printf("Starting application server at http://%s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
