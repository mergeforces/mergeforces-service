package router

import (
	"github.com/go-chi/chi"

	"github.com/mergeforces/mergeforces-service/api/handlers"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}