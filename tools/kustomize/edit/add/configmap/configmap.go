// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package configmap

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Adds a configmap to the kustomization file.
func ConfigMap(name string, opts ...ConfigMapOpt) *tools.CommandContext {
	options := &configmapOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "add", "configmap")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, name)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type ConfigMapOpt func(*configmapOptions)

func SetToolProvider(p tools.ToolProvider) ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.ToolVersion = v
	}
}

// Specify the behavior for config map generation, i.e whether to create a new configmap (the default),  to merge with a previously defined one, or to replace an existing one. Merge and replace should be used only  when overriding an existing configmap defined in a base
func Behavior(value string) ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.behavior = &value
	}
}

// Disable the name suffix for the configmap
func DisableNameSuffixHash() ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.disableNameSuffixHash = ptrhelpers.Bool(true)
	}
}

// Specify the path to a file to read lines of key=val pairs to create a configmap (i.e. a Docker .env file).
func FromEnvFile(value string) ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.fromEnvFile = &value
	}
}

// Key file can be specified using its file path, in which case file basename will be used as configmap key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid configmap key.
func FromFile(value string) ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.fromFile = &value
	}
}

// Specify a key and literal value to insert in configmap (i.e. mykey=somevalue)
func FromLiteral(value ...string) ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.fromLiteral = append(opts.fromLiteral, value...)
	}
}

// print a stack-trace on error
func StackTrace() ConfigMapOpt {
	return func(opts *configmapOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type configmapOptions struct {
	ToolProvider          tools.ToolProvider
	ToolVersion           string
	behavior              *string
	disableNameSuffixHash *bool
	fromEnvFile           *string
	fromFile              *string
	fromLiteral           []string
	stackTrace            *bool
}

func (o *configmapOptions) toArgs() []string {
	var renderedArgs []string

	if o.behavior != nil {
		renderedArgs = append(renderedArgs, "--behavior", fmt.Sprintf("%s", ptrhelpers.StringValue(o.behavior)))
	}

	if o.disableNameSuffixHash != nil {
		renderedArgs = append(renderedArgs, "--disableNameSuffixHash")
	}

	if o.fromEnvFile != nil {
		renderedArgs = append(renderedArgs, "--from-env-file", fmt.Sprintf("%s", ptrhelpers.StringValue(o.fromEnvFile)))
	}

	if o.fromFile != nil {
		renderedArgs = append(renderedArgs, "--from-file", fmt.Sprintf("%s", ptrhelpers.StringValue(o.fromFile)))
	}

	if o.fromLiteral != nil {
		for _, v := range o.fromLiteral {
			renderedArgs = append(renderedArgs, "--from-literal", fmt.Sprintf("%s", v))
		}
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}
