package tools

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

type CommandContext struct {
	Path string
	Args []string
}

type command struct {
	cmd         string
	args        []string
	env         map[string]string
	wd          string
	ignoreOSEnv bool
	logger      Logger
	stdin       io.Reader
}

type CommandOption = func(c *command)

func SetEnvironmentVariables(env map[string]string) CommandOption {
	return func(c *command) {
		c.env = env
	}
}

func SetWorkingDirectory(wd string) CommandOption {
	return func(c *command) {
		c.wd = wd
	}
}

func IgnoreOSEnvironment() CommandOption {
	return func(c *command) {
		c.ignoreOSEnv = true
	}
}

func SetLogger(logger Logger) CommandOption {
	return func(c *command) {
		c.logger = logger
	}
}

func SetNOOPLogger() CommandOption {
	return func(c *command) {
		c.logger = NOOPLogger{}
	}
}

func SetStdin(stdin io.Reader) CommandOption {
	return func(c *command) {
		c.stdin = stdin
	}
}

func (c *CommandContext) Run(opts ...CommandOption) error {
	return c.RunContext(context.Background(), opts...)
}

func (c *CommandContext) RunContext(ctx context.Context, opts ...CommandOption) error {
	var output io.Writer

	cmd := c.createCommand(opts)

	if _, _, err := cmd.execute(ctx, output, os.Stderr); err != nil {
		return err
	}

	return nil
}

// RunV always outputs to Stdout and Stderr
func (c *CommandContext) RunV(opts ...CommandOption) error {
	return c.RunVContext(context.Background(), opts...)
}

// RunVContext always outputs to Stdout and Stderr
func (c *CommandContext) RunVContext(ctx context.Context, opts ...CommandOption) error {
	_, _, err := c.ExecContext(ctx, os.Stdout, os.Stderr, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandContext) RunOutput(opts ...CommandOption) (string, error) {
	return c.RunOutputContext(context.Background(), opts...)
}

func (c *CommandContext) RunOutputContext(ctx context.Context, opts ...CommandOption) (string, error) {
	builder := strings.Builder{}

	_, _, err := c.ExecContext(ctx, &builder, os.Stderr, opts...)
	if err != nil {
		return builder.String(), err
	}

	return builder.String(), nil
}

func (c *CommandContext) Exec(stdout, stderr io.Writer, opts ...CommandOption) (ran bool, exitCode int, err error) {
	return c.ExecContext(context.Background(), stdout, stderr, opts...)
}

func (c *CommandContext) ExecContext(ctx context.Context, stdout, stderr io.Writer, opts ...CommandOption) (ran bool, exitCode int, err error) {
	cmd := c.createCommand(opts)

	ran, exitCode, err = cmd.execute(ctx, stdout, stderr)
	if err == nil {
		return true, exitCode, err
	}

	if ran {
		return ran, exitCode, fmt.Errorf(`running "%s %s" failed with exit code %d`, cmd.cmd, fmtArgs(cmd.args), exitCode)
	}

	return ran, -1, fmt.Errorf(`failed to run "%s %s: %v"`, cmd.cmd, fmtArgs(cmd.args), err)
}
