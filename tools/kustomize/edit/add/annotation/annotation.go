// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package annotation

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Adds one or more commonAnnotations to kustomization.yaml
func Annotation(annotation string, opts ...AnnotationOpt) *tools.CommandContext {
	options := &annotationOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "add", "annotation")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, annotation)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Adds one or more commonAnnotations to kustomization.yaml
func AnnotationWithAnnotations(annotations []string, opts ...AnnotationOpt) *tools.CommandContext {
	options := &annotationOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "add", "annotation")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, annotations...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type AnnotationOpt func(*annotationOptions)

func SetToolProvider(p tools.ToolProvider) AnnotationOpt {
	return func(opts *annotationOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) AnnotationOpt {
	return func(opts *annotationOptions) {
		opts.ToolVersion = v
	}
}

// overwrite commonAnnotation if it already exists
func Force() AnnotationOpt {
	return func(opts *annotationOptions) {
		opts.force = ptrhelpers.Bool(true)
	}
}

// print a stack-trace on error
func StackTrace() AnnotationOpt {
	return func(opts *annotationOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type annotationOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	force        *bool
	stackTrace   *bool
}

func (o *annotationOptions) toArgs() []string {
	var renderedArgs []string

	if o.force != nil {
		renderedArgs = append(renderedArgs, "--force")
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}