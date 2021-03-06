// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/lefthook.json

package uninstall

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/lefthook"
)

// Revert install command
func Uninstall(opts ...UninstallOpt) *tools.CommandContext {
	options := &uninstallOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "uninstall")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(lefthook.LefthookToolInfo, defaults.String(options.ToolVersion, lefthook.DefaultToolVersion), options.ToolProvider, lefthook.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type UninstallOpt func(*uninstallOptions)

func SetToolProvider(p tools.ToolProvider) UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.ToolVersion = v
	}
}

// keep configuration files and source directories present
func KeepConfig() UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.keepConfig = ptrhelpers.Bool(true)
	}
}

// remove all hooks from .git/hooks dir and install lefthook hooks
func Aggressive() UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.aggressive = ptrhelpers.Bool(true)
	}
}

// reinstall hooks without checking config version
func Force() UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.force = ptrhelpers.Bool(true)
	}
}

// disable colored output
func NoColors() UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.noColors = ptrhelpers.Bool(true)
	}
}
func Verbose() UninstallOpt {
	return func(opts *uninstallOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

type uninstallOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	keepConfig   *bool
	aggressive   *bool
	force        *bool
	noColors     *bool
	verbose      *bool
}

func (o *uninstallOptions) toArgs() []string {
	var renderedArgs []string

	if o.keepConfig != nil {
		renderedArgs = append(renderedArgs, "--keep-config")
	}

	if o.aggressive != nil {
		renderedArgs = append(renderedArgs, "--aggressive")
	}

	if o.force != nil {
		renderedArgs = append(renderedArgs, "--aggressive")
	}

	if o.noColors != nil {
		renderedArgs = append(renderedArgs, "--no-colors")
	}

	if o.verbose != nil {
		renderedArgs = append(renderedArgs, "--verbose")
	}

	return renderedArgs
}
