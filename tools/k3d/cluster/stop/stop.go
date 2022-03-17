// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/k3d.json

package stop

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/k3d"
)

// Stop existing k3d cluster(s).
func Stop(opts ...StopOpt) *tools.CommandContext {
	options := &stopOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cluster", "stop")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(k3d.K3dToolInfo, defaults.String(options.ToolVersion, k3d.DefaultToolVersion), options.ToolProvider, k3d.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Stop existing k3d cluster(s).
func StopWithName(name string, opts ...StopOpt) *tools.CommandContext {
	options := &stopOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cluster", "stop")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, name)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(k3d.K3dToolInfo, defaults.String(options.ToolVersion, k3d.DefaultToolVersion), options.ToolProvider, k3d.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Stop existing k3d cluster(s).
func StopWithNames(names []string, opts ...StopOpt) *tools.CommandContext {
	options := &stopOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cluster", "stop")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, names...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(k3d.K3dToolInfo, defaults.String(options.ToolVersion, k3d.DefaultToolVersion), options.ToolProvider, k3d.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type StopOpt func(*stopOptions)

func SetToolProvider(p tools.ToolProvider) StopOpt {
	return func(opts *stopOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) StopOpt {
	return func(opts *stopOptions) {
		opts.ToolVersion = v
	}
}

// Stops all existing clusters
func All() StopOpt {
	return func(opts *stopOptions) {
		opts.all = ptrhelpers.Bool(true)
	}
}

// Enable Log timestamps
func Timestamps() StopOpt {
	return func(opts *stopOptions) {
		opts.timestamps = ptrhelpers.Bool(true)
	}
}

// Enable super verbose output (trace logging)
func Trace() StopOpt {
	return func(opts *stopOptions) {
		opts.trace = ptrhelpers.Bool(true)
	}
}

// Enable verbose output (debug logging)
func Verbose() StopOpt {
	return func(opts *stopOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

type stopOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	all          *bool
	timestamps   *bool
	trace        *bool
	verbose      *bool
}

func (o *stopOptions) toArgs() []string {
	var renderedArgs []string

	if o.all != nil {
		renderedArgs = append(renderedArgs, "--all")
	}

	if o.timestamps != nil {
		renderedArgs = append(renderedArgs, "--timestamps")
	}

	if o.trace != nil {
		renderedArgs = append(renderedArgs, "--trace")
	}

	if o.verbose != nil {
		renderedArgs = append(renderedArgs, "--verbose")
	}

	return renderedArgs
}
