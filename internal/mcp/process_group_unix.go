//go:build unix

package mcp

import (
"os/exec"
"syscall"
)

func configureProcessGroup(cmd *exec.Cmd) {
cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func killCmdProcessGroup(cmd *exec.Cmd) {
if cmd == nil || cmd.Process == nil {
return
}
pgid, err := syscall.Getpgid(cmd.Process.Pid)
if err == nil {
_ = syscall.Kill(-pgid, syscall.SIGKILL)
}
}
