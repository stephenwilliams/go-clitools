// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package graph

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

// Graph prints the module requirement graph (with replacements applied)
// in text form. Each line in the output has two space-separated fields: a module
// and one of its requirements. Each module is identified as a string of the form
// path@version, except for the main module, which has no @version suffix.
func Graph(opts ...GraphOpt) *tools.CommandContext {
	options := &graphOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "mod", "graph")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type GraphOpt func(*graphOptions)

func SetToolProvider(p tools.ToolProvider) GraphOpt {
	return func(opts *graphOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) GraphOpt {
	return func(opts *graphOptions) {
		opts.ToolVersion = v
	}
}

type graphOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
}

func (o *graphOptions) toArgs() []string {
	var renderedArgs []string

	return renderedArgs
}
