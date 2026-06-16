package server

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

//go:embed all:web/dist
var frontendFiles embed.FS

func distFS() fs.FS {
	sub, err := fs.Sub(frontendFiles, "web/dist")
	if err != nil {
		panic(err)
	}
	return sub
}

// spaHandler serves static assets, falling back to index.html for client-side
// routes. Requests for missing files that look like assets (i.e. have a file
// extension) return a real 404 instead of masking it with index.html.
func spaHandler() http.Handler {
	dist := distFS()
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
