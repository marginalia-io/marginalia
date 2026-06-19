//go:build windows

package storage

import (
	"syscall"
	"unsafe"
)

func diskUsage(path string) (available, total uint64, err error) {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, 0, err
	}

	var freeBytesAvailable, totalBytes, totalFreeBytes uint64
	if err := syscall.GetDiskFreeSpaceEx(
		pathPtr,
		(*uint64)(unsafe.Pointer(&freeBytesAvailable)),
		(*uint64)(unsafe.Pointer(&totalBytes)),
		(*uint64)(unsafe.Pointer(&totalFreeBytes)),
	); err != nil {
		return 0, 0, err
	}

	return freeBytesAvailable, totalBytes, nil
}
