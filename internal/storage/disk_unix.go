//go:build unix

package storage

import "golang.org/x/sys/unix"

func diskUsage(path string) (available, total uint64, err error) {
	var stat unix.Statfs_t
	if err := unix.Statfs(path, &stat); err != nil {
		return 0, 0, err
	}

	bsize := uint64(stat.Bsize)
	return stat.Bavail * bsize, stat.Blocks * bsize, nil
}
