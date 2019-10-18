package router

import (
	"github.com/go-chi/chi"

	"github.com/mergeforces/mergeforces-service/api/app"
	"github.com/mergeforces/mergeforces-service/api/handler"
)

func New(app *app.App)  *chi.Mux {
	l := app.Logger()

	r := chi.NewRouter()
	r.Get("/health/live", app.HandleLive)
	r.Method("GET", "/health/ready", handler.NewHandler(app.HandleReady, l))

	r.Method("GET", "/events", handler.NewHandler(app.HandleListEvents, l))
	r.Method("POST", "/events", handler.NewHandler(app.HandleCreateEvent, l))
	r.Method("GET", "/events/{id}", handler.NewHandler(app.HandleReadEvent, l))
	r.Method("PUT", "/events/{id}", handler.NewHandler(app.HandleUpdateEvent, l))
	r.Method("DELETE", "/events/{id}", handler.NewHandler(app.HandleDeleteEvent, l))

	r.Method("GET", "/", handler.NewHandler(app.HandleIndex, l))

	return r
}