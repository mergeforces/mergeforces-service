package app

import (
	"github.com/mergeforces/mergeforces-service/util/logger"
)

type App struct{
	logger *logger.Logger
}

func New(logger *logger.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Logger() *logger.Logger {
	return a.logger
}