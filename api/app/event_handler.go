package app

import (
	"net/http"
)

func (app *App) HandleListEvents(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
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