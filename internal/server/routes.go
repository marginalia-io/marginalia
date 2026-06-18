package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func apiRouter(a *api) http.Handler {
	r := chi.NewRouter()
	r.Get("/health", a.handleHealth)
	r.Get("/setup", a.handleSetupStatus)
	return r
}
