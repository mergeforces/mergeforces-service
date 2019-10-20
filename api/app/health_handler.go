package app

import (
	"net/http"
)

// HandleLive is an http.HandlerFunc that handles liveness checks by
// immediately responding with an HTTP 200 status.
func (app *App) HandleLive(w http.ResponseWriter, _ *http.Request) {
	_, err := writeHealthy(w)
	if err != nil {
		app.Logger().Fatal().Err(err).Msg("Failed to write healthy response")
	}
}

// HandleReady is an http.HandlerFunc that handles readiness checks by
// responding with an HTTP 200 status if it is healthy, 500 otherwise.
func (app *App) HandleReady(w http.ResponseWriter, r *http.Request) {
	if err := app.db.DB().Ping(); err != nil {
		app.Logger().Fatal().Err(err).Msg("")
		_, err := writeUnhealthy(w)
		if err != nil {
			app.Logger().Fatal().Err(err).Msg("Failed to write unhealthy response")
		}
		return
	}

	_, err := writeHealthy(w)
	if err != nil {
		app.Logger().Fatal().Err(err).Msg("Failed to write healthy response")
	}
}

func writeHealthy(w http.ResponseWriter) (int, error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	return w.Write([]byte("ok"))
}

func writeUnhealthy(w http.ResponseWriter) (int, error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	return w.Write([]byte("ok"))
}
