// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package cat

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// [Alpha] Print Resource Config from a local directory.
func Cat(opts ...CatOpt) *tools.CommandContext {
	options := &catOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cfg", "cat")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// [Alpha] Print Resource Config from a local directory.
func CatWithPath(path string, opts ...CatOpt) *tools.CommandContext {
	options := &catOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cfg", "cat")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, path)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type CatOpt func(*catOptions)

func SetToolProvider(p tools.ToolProvider) CatOpt {
	return func(opts *catOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) CatOpt {
	return func(opts *catOptions) {
		opts.ToolVersion = v
	}
}

// annotate resources with their file origins.
func Annotate() CatOpt {
	return func(opts *catOptions) {
		opts.annotate = ptrhelpers.Bool(true)
	}
}

// if specified, write output to a file rather than stdout
func Dest(value string) CatOpt {
	return func(opts *catOptions) {
		opts.dest = &value
	}
}

// if true, exclude non-local-config in the output.
func ExcludeNonLocal() CatOpt {
	return func(opts *catOptions) {
		opts.excludeNonLocal = ptrhelpers.Bool(true)
	}
}

// format resource config yaml before printing. (default true)
func Format(value bool) CatOpt {
	return func(opts *catOptions) {
		opts.format = &value
	}
}

// path to function config to put in ResourceList -- only if wrapped in a ResourceList.
func FunctionConfig(value string) CatOpt {
	return func(opts *catOptions) {
		opts.functionConfig = &value
	}
}

// if true, include local-config in the output.
func IncludeLocal(value bool) CatOpt {
	return func(opts *catOptions) {
		opts.includeLocal = &value
	}
}

// print resources recursively in all the nested subpackages (default true)
func RecurseSubpackages(value bool) CatOpt {
	return func(opts *catOptions) {
		opts.recurseSubpackages = &value
	}
}

// remove comments from yaml.
func StripComments() CatOpt {
	return func(opts *catOptions) {
		opts.stripComments = ptrhelpers.Bool(true)
	}
}

// yaml styles to apply.  may be 'TaggedStyle', 'DoubleQuotedStyle', 'LiteralStyle', 'FoldedStyle', 'FlowStyle'.
func Style(value string) CatOpt {
	return func(opts *catOptions) {
		opts.style = &value
	}
}

// if set, wrap the output in this list type kind.
func WrapKind(value string) CatOpt {
	return func(opts *catOptions) {
		opts.wrapKind = &value
	}
}

// if set, wrap the output in this list type apiVersion.
func WrapVersion(value string) CatOpt {
	return func(opts *catOptions) {
		opts.wrapVersion = &value
	}
}

// print a stack-trace on error
func StackTrace() CatOpt {
	return func(opts *catOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type catOptions struct {
	ToolProvider       tools.ToolProvider
	ToolVersion        string
	annotate           *bool
	dest               *string
	excludeNonLocal    *bool
	format             *bool
	functionConfig     *string
	includeLocal       *bool
	recurseSubpackages *bool
	stripComments      *bool
	style              *string
	wrapKind           *string
	wrapVersion        *string
	stackTrace         *bool
}

func (o *catOptions) toArgs() []string {
	var renderedArgs []string

	if o.annotate != nil {
		renderedArgs = append(renderedArgs, "--annotate")
	}

	if o.dest != nil {
		renderedArgs = append(renderedArgs, "--dest", fmt.Sprintf("%s", ptrhelpers.StringValue(o.dest)))
	}

	if o.excludeNonLocal != nil {
		renderedArgs = append(renderedArgs, "--exclude-non-local")
	}

	if o.format != nil {
		renderedArgs = append(renderedArgs, "--format", fmt.Sprintf("%v", ptrhelpers.BoolValue(o.format)))
	}

	if o.functionConfig != nil {
		renderedArgs = append(renderedArgs, "--function-config", fmt.Sprintf("%s", ptrhelpers.StringValue(o.functionConfig)))
	}

	if o.includeLocal != nil {
		renderedArgs = append(renderedArgs, "--include-local")
	}

	if o.recurseSubpackages != nil {
		renderedArgs = append(renderedArgs, "--recurse-subpackages", fmt.Sprintf("%v", ptrhelpers.BoolValue(o.recurseSubpackages)))
	}

	if o.stripComments != nil {
		renderedArgs = append(renderedArgs, "--strip-comments")
	}

	if o.style != nil {
		renderedArgs = append(renderedArgs, "--style", fmt.Sprintf("%s", ptrhelpers.StringValue(o.style)))
	}

	if o.wrapKind != nil {
		renderedArgs = append(renderedArgs, "--wrap-kind", fmt.Sprintf("%s", ptrhelpers.StringValue(o.wrapKind)))
	}

	if o.wrapVersion != nil {
		renderedArgs = append(renderedArgs, "--wrap-version", fmt.Sprintf("%s", ptrhelpers.StringValue(o.wrapVersion)))
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}