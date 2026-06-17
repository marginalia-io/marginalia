//go:build !dev

// Package embed holds the built frontend assets that are compiled into the
// binary. The dist directory is produced by the frontend build (see the
// Makefile) and embedded here so the server can serve the SPA from a single
// static binary.
package embed

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var assets embed.FS

// Enabled reports whether the frontend assets were embedded in this build. It
// is false in dev builds (-tags dev), which skip the embed entirely.
const Enabled = true

// FS returns the embedded frontend build output (the contents of dist).
func FS() fs.FS {
	sub, err := fs.Sub(assets, "dist")
	if err != nil {
		panic(err)
	}
	return sub
}
