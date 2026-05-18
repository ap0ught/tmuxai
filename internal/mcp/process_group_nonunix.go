//go:build !unix

package mcp

import "os/exec"

func configureProcessGroup(cmd *exec.Cmd) {
_ = cmd
}

func killCmdProcessGroup(cmd *exec.Cmd) {
_ = cmd
}
