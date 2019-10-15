package router

import (
	"github.com/go-chi/chi"

	"github.com/mergeforces/mergeforces-service/api/app"
	"github.com/mergeforces/mergeforces-service/api/handler"
)

func New(app *app.App)  *chi.Mux {
	l := app.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", handler.NewHandler(app.HandleIndex, l))

	return r
}