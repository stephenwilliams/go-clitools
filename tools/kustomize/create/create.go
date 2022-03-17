// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package create

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Create a new kustomization in the current directory
func Create(opts ...CreateOpt) *tools.CommandContext {
	options := &createOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type CreateOpt func(*createOptions)

func SetToolProvider(p tools.ToolProvider) CreateOpt {
	return func(opts *createOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) CreateOpt {
	return func(opts *createOptions) {
		opts.ToolVersion = v
	}
}

// Add one or more common annotations.
func Annotations(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.annotations = &value
	}
}

// Search for kubernetes resources in the current directory to be added to the kustomization file.
func Autodetect() CreateOpt {
	return func(opts *createOptions) {
		opts.autodetect = ptrhelpers.Bool(true)
	}
}

// Add one or more common labels
func Labels(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.labels = &value
	}
}

// Sets the value of the namePrefix field in the kustomization file.
func NamePrefix(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.nameprefix = &value
	}
}

// Set the value of the namespace field in the customization file.
func Namespace(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.namespace = &value
	}
}

// Sets the value of the nameSuffix field in the kustomization file.
func NameSuffix(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.namesuffix = &value
	}
}

// Enable recursive directory searching for resource auto-detection.
func Recursive(value bool) CreateOpt {
	return func(opts *createOptions) {
		opts.recursive = &value
	}
}

// Name of a file containing a file to add to the kustomization file.
func Resources(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.resources = &value
	}
}

// print a stack-trace on error
func StackTrace() CreateOpt {
	return func(opts *createOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type createOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	annotations  *string
	autodetect   *bool
	labels       *string
	nameprefix   *string
	namespace    *string
	namesuffix   *string
	recursive    *bool
	resources    *string
	stackTrace   *bool
}

func (o *createOptions) toArgs() []string {
	var renderedArgs []string

	if o.annotations != nil {
		renderedArgs = append(renderedArgs, "--annotations", fmt.Sprintf("%s", ptrhelpers.StringValue(o.annotations)))
	}

	if o.autodetect != nil {
		renderedArgs = append(renderedArgs, "--autodetect")
	}

	if o.labels != nil {
		renderedArgs = append(renderedArgs, "--labels", fmt.Sprintf("%s", ptrhelpers.StringValue(o.labels)))
	}

	if o.nameprefix != nil {
		renderedArgs = append(renderedArgs, "--nameprefix", fmt.Sprintf("%s", ptrhelpers.StringValue(o.nameprefix)))
	}

	if o.namespace != nil {
		renderedArgs = append(renderedArgs, "--namespace", fmt.Sprintf("%s", ptrhelpers.StringValue(o.namespace)))
	}

	if o.namesuffix != nil {
		renderedArgs = append(renderedArgs, "--namesuffix", fmt.Sprintf("%s", ptrhelpers.StringValue(o.namesuffix)))
	}

	if o.recursive != nil {
		renderedArgs = append(renderedArgs, "--recursive")
	}

	if o.resources != nil {
		renderedArgs = append(renderedArgs, "--resources", fmt.Sprintf("%s", ptrhelpers.StringValue(o.resources)))
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}
