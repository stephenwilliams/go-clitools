// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package namesuffix

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Sets the value of the nameSuffix field in the kustomization file.
func NameSuffix(suffix string, opts ...NameSuffixOpt) *tools.CommandContext {
	options := &namesuffixOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "set", "namesuffix")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, suffix)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type NameSuffixOpt func(*namesuffixOptions)

func SetToolProvider(p tools.ToolProvider) NameSuffixOpt {
	return func(opts *namesuffixOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) NameSuffixOpt {
	return func(opts *namesuffixOptions) {
		opts.ToolVersion = v
	}
}

type namesuffixOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
}

func (o *namesuffixOptions) toArgs() []string {
	var renderedArgs []string

	return renderedArgs
}
