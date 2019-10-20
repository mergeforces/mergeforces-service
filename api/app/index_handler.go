package app

import "net/http"

func (app *App) HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Hello World!"))

	if err != nil {
		app.logger.Warn().Err(err).Msg("Handling Index Failed")
	}
}
