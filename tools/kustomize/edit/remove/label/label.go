// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package label

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Removes one or more commonLabels from kustomization.yaml
func Label(label string, opts ...LabelOpt) *tools.CommandContext {
	options := &labelOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "remove", "label")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, label)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type LabelOpt func(*labelOptions)

func SetToolProvider(p tools.ToolProvider) LabelOpt {
	return func(opts *labelOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) LabelOpt {
	return func(opts *labelOptions) {
		opts.ToolVersion = v
	}
}

// ignore error if the given label doesn't exist
func IgnoreNonExistence() LabelOpt {
	return func(opts *labelOptions) {
		opts.ignoreNonExistence = ptrhelpers.Bool(true)
	}
}

// print a stack-trace on error
func StackTrace() LabelOpt {
	return func(opts *labelOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type labelOptions struct {
	ToolProvider       tools.ToolProvider
	ToolVersion        string
	ignoreNonExistence *bool
	stackTrace         *bool
}

func (o *labelOptions) toArgs() []string {
	var renderedArgs []string

	if o.ignoreNonExistence != nil {
		renderedArgs = append(renderedArgs, "--ignore-non-existence")
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}
