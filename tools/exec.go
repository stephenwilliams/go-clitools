package tools

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func (c *CommandContext) createCommand(opts []CommandOption) *command {
	cmd := &command{
		cmd:    c.Path,
		args:   c.Args,
		logger: DefaultLogger,
	}

	for _, opt := range opts {
		opt(cmd)
	}

	expandMapper := func(k string) string {
		if v, ok := cmd.env[k]; ok {
			return v
		}

		return os.Getenv(k)
	}

	for i, arg := range cmd.args {
		cmd.args[i] = os.Expand(arg, expandMapper)
	}

	return cmd
}

func (c *command) execute(ctx context.Context, stdout, stderr io.Writer) (ran bool, exitCode int, err error) {
	cmd := exec.CommandContext(ctx, c.cmd, c.args...)

	if !c.ignoreOSEnv {
		cmd.Env = os.Environ()
	}

	for k, v := range c.env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Stdin = os.Stdin
	c.logger.Log(cmd.Path, c.args, cmd.Env)
	err = cmd.Run()
	return CmdRan(err), GetExitStatus(err), err
}
