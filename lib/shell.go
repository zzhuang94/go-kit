package lib

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func RunCmd(cmd string, timeout ...int) (string, int, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	if len(timeout) > 0 && timeout[0] > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout[0])*time.Second)
		defer cancel()
	} else {
		ctx = context.Background()
	}

	args := []string{}
	cmds := []string{"/bin/sh", "-c", cmd}
	if len(cmds) > 1 {
		args = cmds[1:]
	}

	c := exec.CommandContext(ctx, cmds[0], args...)

	var stdout bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stdout

	err := c.Run()
	if err != nil {
		code := c.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		return stdout.String(), code, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return "", -1, fmt.Errorf("timeout")
	}

	return strings.TrimSpace(stdout.String()), 0, err
}
