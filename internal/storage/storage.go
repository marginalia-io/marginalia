package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// DefaultPath is the default on-disk location for the book library.
const DefaultPath = "library"

// Info describes the configured library storage directory and the disk space
// on the volume that contains it.
type Info struct {
	Path           string `json:"path"`
	AvailableBytes uint64 `json:"available_bytes"`
	TotalBytes     uint64 `json:"total_bytes"`
}

// Stat resolves path to an absolute location and reports disk space for the
// filesystem that would hold it. If path does not exist yet, the nearest
// existing parent directory is used for the space query.
func Stat(path string) (Info, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return Info{}, fmt.Errorf("resolve storage path: %w", err)
	}

	statPath, err := existingAncestor(abs)
	if err != nil {
		return Info{}, err
	}

	available, total, err := diskUsage(statPath)
	if err != nil {
		return Info{}, fmt.Errorf("disk usage for %s: %w", statPath, err)
	}

	return Info{
		Path:           abs,
		AvailableBytes: available,
		TotalBytes:     total,
	}, nil
}

func existingAncestor(path string) (string, error) {
	for {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		} else if !os.IsNotExist(err) {
			return "", fmt.Errorf("stat %s: %w", path, err)
		}

		parent := filepath.Dir(path)
		if parent == path {
			return "", fmt.Errorf("storage path has no existing ancestor: %s", path)
		}
		path = parent
	}
}
