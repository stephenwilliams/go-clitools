// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package env

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

// Env prints Go environment information.
//
// By default env prints information as a shell script
// (on Windows, a batch file). If one or more variable
// names is given as arguments, env prints the value of
// each named variable on its own line.
func Env(opts ...EnvOpt) *tools.CommandContext {
	options := &envOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "env")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Env prints Go environment information.
//
// By default env prints information as a shell script
// (on Windows, a batch file). If one or more variable
// names is given as arguments, env prints the value of
// each named variable on its own line.
func EnvWithVars(vars []string, opts ...EnvOpt) *tools.CommandContext {
	options := &envOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "env")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, vars...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type EnvOpt func(*envOptions)

func SetToolProvider(p tools.ToolProvider) EnvOpt {
	return func(opts *envOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) EnvOpt {
	return func(opts *envOptions) {
		opts.ToolVersion = v
	}
}

// prints the environment in JSON format
// instead of as a shell script.
func JSON() EnvOpt {
	return func(opts *envOptions) {
		opts.json = ptrhelpers.Bool(true)
	}
}

// requires one or more arguments and unsets
// the default setting for the named environment variables,
// if one has been set with 'go env -w'.
func Unset() EnvOpt {
	return func(opts *envOptions) {
		opts.unset = ptrhelpers.Bool(true)
	}
}

// requires one or more arguments of the
// form NAME=VALUE and changes the default settings
// of the named environment variables to the given values.
func Write() EnvOpt {
	return func(opts *envOptions) {
		opts.write = ptrhelpers.Bool(true)
	}
}

type envOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	json         *bool
	unset        *bool
	write        *bool
}

func (o *envOptions) toArgs() []string {
	var renderedArgs []string

	if o.json != nil {
		renderedArgs = append(renderedArgs, "-json")
	}

	if o.unset != nil {
		renderedArgs = append(renderedArgs, "-u")
	}

	if o.write != nil {
		renderedArgs = append(renderedArgs, "-w")
	}

	return renderedArgs
}
