package server

import (
	"io"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"marginalia/internal/server/embed"
)

// spaHandler serves the embedded frontend, falling back to index.html for
// client-side routes. Requests for missing files that look like assets (i.e.
// have a file extension) return a real 404 instead of masking it with
// index.html.
//
// In dev builds (-tags dev) the frontend is not embedded; this returns a stub
// that points at the Vite dev server, which proxies /api back to this server.
func spaHandler() http.Handler {
	if !embed.Enabled {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			_, _ = io.WriteString(w, "marginalia dev build: the frontend is served by the Vite dev server (run `pnpm dev` or `make dev`). The API is available under /api.\n")
		})
	}

	dist := embed.FS()
	fileServer := http.FileServer(http.FS(dist))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, "/")
		if p == "" {
			p = "index.html"
		}
		if _, err := fs.Stat(dist, p); err != nil {
			if path.Ext(p) != "" {
				http.NotFound(w, r)
				return
			}
			r.URL.Path = "/"
		}
		fileServer.ServeHTTP(w, r)
	})
}
