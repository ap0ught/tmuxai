//go:build !unix

package mcp

import "os/exec"

func configureProcessGroup(cmd *exec.Cmd) {
	_ = cmd
}

func killCmdProcessGroup(cmd *exec.Cmd) {
	if cmd == nil || cmd.Process == nil {
		return
	}
	_ = cmd.Process.Kill()
}
