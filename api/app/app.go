package app

import (
	"github.com/jinzhu/gorm"

	"github.com/mergeforces/mergeforces-service/pkg/util/logger"
)

type App struct{
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB, ) *App {
	return &App{logger: logger,db: db,}
}

func (a *App) Logger() *logger.Logger {
	return a.logger
}