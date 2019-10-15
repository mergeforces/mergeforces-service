package main

import (
	"fmt"
	"log"
	"net/http"

	r "github.com/mergeforces/mergeforces-service/api/router"
	c "github.com/mergeforces/mergeforces-service/config"
)

func main() {
	config := c.AppConfig()
	router := r.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Ping)

	address := fmt.Sprintf(":%d", config.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      router,
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