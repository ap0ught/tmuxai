//go:build windows

package internal

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows"
)

func waitForInput(fd int, timeout time.Duration) (bool, error) {
	if timeout <= 0 {
		timeout = 10 * time.Millisecond
	}
	pollTimeout := uint32(timeout / time.Millisecond)
	if pollTimeout == 0 {
		pollTimeout = 1
	}

	result, err := windows.WaitForSingleObject(windows.Handle(fd), pollTimeout)
	if err != nil {
		return false, err
	}

	switch result {
	case windows.WAIT_OBJECT_0:
		return true, nil
	case uint32(windows.WAIT_TIMEOUT):
		return false, nil
	default:
		return false, fmt.Errorf("WaitForSingleObject returned %d", result)
	}
}
