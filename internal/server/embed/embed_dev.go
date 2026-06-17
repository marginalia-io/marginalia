//go:build dev

package embed

import "io/fs"

// Enabled is false in dev builds: the frontend is served by the Vite dev
// server rather than embedded, so no dist directory is required to compile.
const Enabled = false

// FS returns nil in dev builds; callers must guard on Enabled.
func FS() fs.FS { return nil }
