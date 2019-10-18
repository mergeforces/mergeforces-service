package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mergeforces/mergeforces-service/pkg/repository"
)

func (app *App) HandleListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := repository.ListEvents(app.db)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}

	if events == nil {
		fmt.Fprint(w, "[]")
		return
	}

	dtos := events.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}
}

func (app *App) HandleCreateEvent(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
}

func (app *App) HandleReadEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
}

func (app *App) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func (app *App) HandleDeleteEvent(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusAccepted)
}