// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package tidy

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

// Tidy makes sure go.mod matches the source code in the module.
// It adds any missing modules necessary to build the current module's
// packages and dependencies, and it removes unused modules that
// don't provide any relevant packages. It also adds any missing entries
// to go.sum and removes any unnecessary ones.
func Tidy(opts ...TidyOpt) *tools.CommandContext {
	options := &tidyOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "mod", "tidy")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type TidyOpt func(*tidyOptions)

func SetToolProvider(p tools.ToolProvider) TidyOpt {
	return func(opts *tidyOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) TidyOpt {
	return func(opts *tidyOptions) {
		opts.ToolVersion = v
	}
}

// Prints information about removed modules to standard error.
func Verbose() TidyOpt {
	return func(opts *tidyOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

type tidyOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	verbose      *bool
}

func (o *tidyOptions) toArgs() []string {
	var renderedArgs []string

	if o.verbose != nil {
		renderedArgs = append(renderedArgs, "-v")
	}

	return renderedArgs
}
