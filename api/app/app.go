package app

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"github.com/mergeforces/mergeforces-service/pkg/util/logger"
)

const (
	appErrDataCreationFailure = "data creation failure"
	appErrDataAccessFailure   = "data access failure"
	appErrDataUpdateFailure   = "data update failure"
	appErrJsonCreationFailure = "json creation failure"
	appErrFormDecodingFailure = "form decoding failure"
)

type App struct {
	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func New(logger *logger.Logger, db *gorm.DB, validator *validator.Validate, ) *App {
	return &App{logger: logger, db: db, validator: validator,}
}

func (a *App) Logger() *logger.Logger {
	return a.logger
}
