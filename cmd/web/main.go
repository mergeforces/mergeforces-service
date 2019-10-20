package main

import (
	"fmt"
	"net/http"

	"github.com/mergeforces/mergeforces-service/api/app"
	r "github.com/mergeforces/mergeforces-service/api/router"
	c "github.com/mergeforces/mergeforces-service/config"
	dbConn "github.com/mergeforces/mergeforces-service/pkg/adapter/gorm"
	l "github.com/mergeforces/mergeforces-service/pkg/util/logger"
	vr "github.com/mergeforces/mergeforces-service/pkg/util/validator"
)

func main() {
	config := c.AppConfig()
	logger := l.New(config.Debug)
	validator := vr.New()

	db, err := dbConn.New(config)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	if config.Debug {
		db.LogMode(true)
	}

	application := app.New(logger, db, validator)
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
