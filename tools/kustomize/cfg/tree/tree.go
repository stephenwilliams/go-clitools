// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package tree

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// [Alpha] Display Resource structure from a directory or stdin.
//
// kustomize cfg tree may be used to print Resources in a directory or cluster, preserving structure
//
// Args:
//
//   DIR:
//     Path to local directory directory.
//
// Resource fields may be printed as part of the Resources by specifying the fields as flags.
//
// kustomize cfg tree has build-in support for printing common fields, such as replicas, container images,
// container names, etc.
//
// kustomize cfg tree supports printing arbitrary fields using the '--field' flag.
//
// By default, kustomize cfg tree uses Resource graph structure if any relationships between resources (ownerReferences)
// are detected, as is typically the case when printing from a cluster. Otherwise, directory graph structure is used. The
// graph structure can also be selected explicitly using the '--graph-structure' flag.
func Tree(opts ...TreeOpt) *tools.CommandContext {
	options := &treeOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cfg", "tree")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// [Alpha] Display Resource structure from a directory or stdin.
//
// kustomize cfg tree may be used to print Resources in a directory or cluster, preserving structure
//
// Args:
//
//   DIR:
//     Path to local directory directory.
//
// Resource fields may be printed as part of the Resources by specifying the fields as flags.
//
// kustomize cfg tree has build-in support for printing common fields, such as replicas, container images,
// container names, etc.
//
// kustomize cfg tree supports printing arbitrary fields using the '--field' flag.
//
// By default, kustomize cfg tree uses Resource graph structure if any relationships between resources (ownerReferences)
// are detected, as is typically the case when printing from a cluster. Otherwise, directory graph structure is used. The
// graph structure can also be selected explicitly using the '--graph-structure' flag.
func TreeWithPath(path string, opts ...TreeOpt) *tools.CommandContext {
	options := &treeOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cfg", "tree")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, path)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type TreeOpt func(*treeOptions)

func SetToolProvider(p tools.ToolProvider) TreeOpt {
	return func(opts *treeOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) TreeOpt {
	return func(opts *treeOptions) {
		opts.ToolVersion = v
	}
}

// print all field infos
func All() TreeOpt {
	return func(opts *treeOptions) {
		opts.all = ptrhelpers.Bool(true)
	}
}

// print args field
func Args() TreeOpt {
	return func(opts *treeOptions) {
		opts.args = ptrhelpers.Bool(true)
	}
}

// print command field
func Command() TreeOpt {
	return func(opts *treeOptions) {
		opts.command = ptrhelpers.Bool(true)
	}
}

// print env field
func Env() TreeOpt {
	return func(opts *treeOptions) {
		opts.env = ptrhelpers.Bool(true)
	}
}

// if true, exclude non-local-config in the output.
func ExcludeNonLocal() TreeOpt {
	return func(opts *treeOptions) {
		opts.excludeNonLocal = ptrhelpers.Bool(true)
	}
}

// print field
func Field(value ...string) TreeOpt {
	return func(opts *treeOptions) {
		opts.field = value
	}
}

// Graph structure to use for printing the tree.  may be any of: owners,directory
func GraphStructure(value string) TreeOpt {
	return func(opts *treeOptions) {
		opts.graphStructure = &value
	}
}

// print image field
func Image() TreeOpt {
	return func(opts *treeOptions) {
		opts.image = ptrhelpers.Bool(true)
	}
}

// if true, include local-config in the output.
func IncludeLocal(value bool) TreeOpt {
	return func(opts *treeOptions) {
		opts.includeLocal = &value
	}
}

// print name field
func Name(value bool) TreeOpt {
	return func(opts *treeOptions) {
		opts.name = &value
	}
}

// print ports field
func Ports(value bool) TreeOpt {
	return func(opts *treeOptions) {
		opts.ports = &value
	}
}

// print replicas field
func Replicas(value bool) TreeOpt {
	return func(opts *treeOptions) {
		opts.replicas = &value
	}
}

// print resources field
func Resources(value bool) TreeOpt {
	return func(opts *treeOptions) {
		opts.resources = &value
	}
}

// print a stack-trace on error
func StackTrace() TreeOpt {
	return func(opts *treeOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type treeOptions struct {
	ToolProvider    tools.ToolProvider
	ToolVersion     string
	all             *bool
	args            *bool
	command         *bool
	env             *bool
	excludeNonLocal *bool
	field           []string
	graphStructure  *string
	image           *bool
	includeLocal    *bool
	name            *bool
	ports           *bool
	replicas        *bool
	resources       *bool
	stackTrace      *bool
}

func (o *treeOptions) toArgs() []string {
	var renderedArgs []string

	if o.all != nil {
		renderedArgs = append(renderedArgs, "--all")
	}

	if o.args != nil {
		renderedArgs = append(renderedArgs, "--args")
	}

	if o.command != nil {
		renderedArgs = append(renderedArgs, "--command")
	}

	if o.env != nil {
		renderedArgs = append(renderedArgs, "--env")
	}

	if o.excludeNonLocal != nil {
		renderedArgs = append(renderedArgs, "--exclude-non-local")
	}

	if o.field != nil {
		renderedArgs = append(renderedArgs, "--field", fmt.Sprintf("%s", o.field))
	}

	if o.graphStructure != nil {
		renderedArgs = append(renderedArgs, "--graph-structure", fmt.Sprintf("%s", ptrhelpers.StringValue(o.graphStructure)))
	}

	if o.image != nil {
		renderedArgs = append(renderedArgs, "--image")
	}

	if o.includeLocal != nil {
		renderedArgs = append(renderedArgs, "--include-local")
	}

	if o.name != nil {
		renderedArgs = append(renderedArgs, "--name")
	}

	if o.ports != nil {
		renderedArgs = append(renderedArgs, "--ports")
	}

	if o.replicas != nil {
		renderedArgs = append(renderedArgs, "--replicas")
	}

	if o.resources != nil {
		renderedArgs = append(renderedArgs, "--resources")
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}