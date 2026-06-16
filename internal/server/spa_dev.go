//go:build dev

package server

import (
	"io"
	"net/http"
)

// spaHandler in dev builds does not embed web/dist, so the binary compiles
// without a prior frontend build. Serve the frontend from the Vite dev server
// (run `pnpm dev` / `make dev`); it proxies /api back to this server.
func spaHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = io.WriteString(w, "marginalia dev build: the frontend is served by the Vite dev server (run `pnpm dev` or `make dev`). The API is available under /api.\n")
	})
}
