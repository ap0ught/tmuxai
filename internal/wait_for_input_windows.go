//go:build windows

package internal

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows"
)

func waitForInput(fd int, timeout time.Duration) (bool, error) {
	_ = fd // readEscapeSequence passes stdin; Windows waits on the stdin handle directly.

	if timeout <= 0 {
		timeout = 10 * time.Millisecond
	}
	pollTimeout := uint32(timeout / time.Millisecond)
	if pollTimeout == 0 {
		pollTimeout = 1
	}

	handle, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
	if err != nil {
		return false, fmt.Errorf("failed to get stdin handle: %w", err)
	}

	result, err := windows.WaitForSingleObject(handle, pollTimeout)
	if err != nil {
		return false, fmt.Errorf("WaitForSingleObject failed: %w", err)
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
