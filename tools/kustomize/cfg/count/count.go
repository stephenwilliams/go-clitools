// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package count

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// [Alpha] Count Resources Config from a local directory.
func Count(opts ...CountOpt) *tools.CommandContext {
	options := &countOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cfg", "count")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// [Alpha] Count Resources Config from a local directory.
func CountWithPath(path string, opts ...CountOpt) *tools.CommandContext {
	options := &countOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cfg", "count")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, path)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type CountOpt func(*countOptions)

func SetToolProvider(p tools.ToolProvider) CountOpt {
	return func(opts *countOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) CountOpt {
	return func(opts *countOptions) {
		opts.ToolVersion = v
	}
}

// print resources recursively in all the nested subpackages (default true)
func RecurseSubpackages(value bool) CountOpt {
	return func(opts *countOptions) {
		opts.recurseSubpackages = &value
	}
}

// count resources by kind. (default true)
func Kind(value bool) CountOpt {
	return func(opts *countOptions) {
		opts.kind = &value
	}
}

// print a stack-trace on error
func StackTrace() CountOpt {
	return func(opts *countOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type countOptions struct {
	ToolProvider       tools.ToolProvider
	ToolVersion        string
	recurseSubpackages *bool
	kind               *bool
	stackTrace         *bool
}

func (o *countOptions) toArgs() []string {
	var renderedArgs []string

	if o.recurseSubpackages != nil {
		renderedArgs = append(renderedArgs, "--recurse-subpackages", fmt.Sprintf("%v", ptrhelpers.BoolValue(o.recurseSubpackages)))
	}

	if o.kind != nil {
		renderedArgs = append(renderedArgs, "--kind", fmt.Sprintf("%v", ptrhelpers.BoolValue(o.kind)))
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}