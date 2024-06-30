package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":8080"
)

func main() {
	router := http.NewServeMux()

	httpHandler := NewHandler()
	httpHandler.registerRoutes(router)

	log.Printf("starting http server at %s ", httpAddr)
	err := http.ListenAndServe(httpAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}
