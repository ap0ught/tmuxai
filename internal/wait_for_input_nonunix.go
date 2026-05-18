//go:build !unix

package internal

import "time"

func waitForInput(fd int, timeout time.Duration) (bool, error) {
	_ = fd
	if timeout <= 0 {
		timeout = 10 * time.Millisecond
	}
	time.Sleep(timeout)
	return false, nil
}
