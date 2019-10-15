package main

import (
	"fmt"
	"github.com/mergeforces/mergeforces-service/config"
	"log"
	"net/http"
)

func main() {
	config := config.AppConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Ping)

	address := fmt.Sprintf(":%d", config.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}