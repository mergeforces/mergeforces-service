package main

import (
	"fmt"
	"net/http"

	a "github.com/mergeforces/mergeforces-service/api/app"
	r "github.com/mergeforces/mergeforces-service/api/router"
	c "github.com/mergeforces/mergeforces-service/config"
	l "github.com/mergeforces/mergeforces-service/pkg/util/logger"
)

func main() {
	config := c.AppConfig()
	logger := l.New(config.Debug)
	application := a.New(logger)
	router := r.New(application)

	address := fmt.Sprintf(":%d", config.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
