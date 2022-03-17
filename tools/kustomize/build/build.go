// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package build

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Build a set of KRM resources using a 'kustomization.yaml' file.
func Build(opts ...BuildOpt) *tools.CommandContext {
	options := &buildOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "build")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Build a set of KRM resources using a 'kustomization.yaml' file.
func BuildWithPath(path string, opts ...BuildOpt) *tools.CommandContext {
	options := &buildOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "build")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, path)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type BuildOpt func(*buildOptions)

func SetToolProvider(p tools.ToolProvider) BuildOpt {
	return func(opts *buildOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) BuildOpt {
	return func(opts *buildOptions) {
		opts.ToolVersion = v
	}
}

// use the uid and gid of the command executor to run the function in the container
func AsCurrentUser() BuildOpt {
	return func(opts *buildOptions) {
		opts.asCurrentUser = ptrhelpers.Bool(true)
	}
}

// enable kustomize plugins
func EnableAlphaPlugins() BuildOpt {
	return func(opts *buildOptions) {
		opts.enableAlphaPlugins = ptrhelpers.Bool(true)
	}
}

// enable support for exec functions (raw executables); do not use for untrusted configs! (Alpha)
func EnableExec() BuildOpt {
	return func(opts *buildOptions) {
		opts.enableExec = ptrhelpers.Bool(true)
	}
}

// Enable use of the Helm chart inflator generator.
func EnableHelm() BuildOpt {
	return func(opts *buildOptions) {
		opts.enableHelm = ptrhelpers.Bool(true)
	}
}

// enable adding app.kubernetes.io/managed-by
func EnableManagedByLabel() BuildOpt {
	return func(opts *buildOptions) {
		opts.enableManagedByLabel = ptrhelpers.Bool(true)
	}
}

// enable support for starlark functions. (Alpha)
func EnableStar() BuildOpt {
	return func(opts *buildOptions) {
		opts.enableStar = ptrhelpers.Bool(true)
	}
}

// a list of environment variables to be used by functions
func Env(value ...string) BuildOpt {
	return func(opts *buildOptions) {
		opts.env = value
	}
}

// helm command (path to executable) (default "helm")
func HelmCommand(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.helmCommand = &value
	}
}

// if set to 'LoadRestrictionsNone', local kustomizations may load files from outside their root. This does, however, break the relocatability of the kustomization. (default "LoadRestrictionsRootOnly")
func LoadRestrictor(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.loadRestrictor = &value
	}
}

// a list of storage options read from the filesystem
func Mount(value ...string) BuildOpt {
	return func(opts *buildOptions) {
		opts.mount = value
	}
}

// enable network access for functions that declare it
func Network() BuildOpt {
	return func(opts *buildOptions) {
		opts.network = ptrhelpers.Bool(true)
	}
}

// the docker network to run the container in (default "bridge")
func NetworkName(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.networkName = &value
	}
}

// If specified, write output to this path.
func Output(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.output = &value
	}
}

// Reorder the resources just before output. Use 'legacy' to apply a legacy reordering (Namespaces first, Webhooks last, etc). Use 'none' to suppress a final reordering. (default "legacy")
func Reorder(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.reorder = &value
	}
}

// print a stack-trace on error
func StackTrace() BuildOpt {
	return func(opts *buildOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type buildOptions struct {
	ToolProvider         tools.ToolProvider
	ToolVersion          string
	asCurrentUser        *bool
	enableAlphaPlugins   *bool
	enableExec           *bool
	enableHelm           *bool
	enableManagedByLabel *bool
	enableStar           *bool
	env                  []string
	helmCommand          *string
	loadRestrictor       *string
	mount                []string
	network              *bool
	networkName          *string
	output               *string
	reorder              *string
	stackTrace           *bool
}

func (o *buildOptions) toArgs() []string {
	var renderedArgs []string

	if o.asCurrentUser != nil {
		renderedArgs = append(renderedArgs, "--as-current-user")
	}

	if o.enableAlphaPlugins != nil {
		renderedArgs = append(renderedArgs, "--enable-alpha-plugins")
	}

	if o.enableExec != nil {
		renderedArgs = append(renderedArgs, "--enable-exec")
	}

	if o.enableHelm != nil {
		renderedArgs = append(renderedArgs, "--enable-helm")
	}

	if o.enableManagedByLabel != nil {
		renderedArgs = append(renderedArgs, "--enable-managed-by-label")
	}

	if o.enableStar != nil {
		renderedArgs = append(renderedArgs, "--enable-star")
	}

	if o.env != nil {
		renderedArgs = append(renderedArgs, "--env", fmt.Sprintf("%s", o.env))
	}

	if o.helmCommand != nil {
		renderedArgs = append(renderedArgs, "--helm-command", fmt.Sprintf("%s", ptrhelpers.StringValue(o.helmCommand)))
	}

	if o.loadRestrictor != nil {
		renderedArgs = append(renderedArgs, "--load-restrictor", fmt.Sprintf("%s", ptrhelpers.StringValue(o.loadRestrictor)))
	}

	if o.mount != nil {
		renderedArgs = append(renderedArgs, "--mount", fmt.Sprintf("%s", o.mount))
	}

	if o.network != nil {
		renderedArgs = append(renderedArgs, "--network")
	}

	if o.networkName != nil {
		renderedArgs = append(renderedArgs, "--network-name", fmt.Sprintf("%s", ptrhelpers.StringValue(o.networkName)))
	}

	if o.output != nil {
		renderedArgs = append(renderedArgs, "--output", fmt.Sprintf("%s", ptrhelpers.StringValue(o.output)))
	}

	if o.reorder != nil {
		renderedArgs = append(renderedArgs, "--reorder", fmt.Sprintf("%s", ptrhelpers.StringValue(o.reorder)))
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}