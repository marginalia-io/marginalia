package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func apiRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", handleHealth)
	return r
}
