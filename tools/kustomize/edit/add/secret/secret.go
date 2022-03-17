// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package secret

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Adds a secret to the kustomization file.
func Secret(name string, opts ...SecretOpt) *tools.CommandContext {
	options := &secretOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "add", "secret")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, name)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type SecretOpt func(*secretOptions)

func SetToolProvider(p tools.ToolProvider) SecretOpt {
	return func(opts *secretOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) SecretOpt {
	return func(opts *secretOptions) {
		opts.ToolVersion = v
	}
}

// Disable the name suffix for the secret
func DisableNameSuffixHash() SecretOpt {
	return func(opts *secretOptions) {
		opts.disableNameSuffixHash = ptrhelpers.Bool(true)
	}
}

// Specify the path to a file to read lines of key=val pairs to create a secret (i.e. a Docker .env file).
func FromEnvFile(value string) SecretOpt {
	return func(opts *secretOptions) {
		opts.fromEnvFile = &value
	}
}

// Key file can be specified using its file path, in which case file basename will be used as secret key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid secret key.
func FromFile(value string) SecretOpt {
	return func(opts *secretOptions) {
		opts.fromFile = &value
	}
}

// Specify a key and literal value to insert in secret (i.e. mykey=somevalue)
func FromLiteral(value ...string) SecretOpt {
	return func(opts *secretOptions) {
		opts.fromLiteral = value
	}
}

// Specify the namespace of the secret
func Namespace(value string) SecretOpt {
	return func(opts *secretOptions) {
		opts.namespace = &value
	}
}

// Specify the secret type this can be 'Opaque' (default), or 'kubernetes.io/tls' (default "Opaque")
func Type(value string) SecretOpt {
	return func(opts *secretOptions) {
		opts.typeName = &value
	}
}

// print a stack-trace on error
func StackTrace() SecretOpt {
	return func(opts *secretOptions) {
		opts.stackTrace = ptrhelpers.Bool(true)
	}
}

type secretOptions struct {
	ToolProvider          tools.ToolProvider
	ToolVersion           string
	disableNameSuffixHash *bool
	fromEnvFile           *string
	fromFile              *string
	fromLiteral           []string
	namespace             *string
	typeName              *string
	stackTrace            *bool
}

func (o *secretOptions) toArgs() []string {
	var renderedArgs []string

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
		renderedArgs = append(renderedArgs, "--from-literal", fmt.Sprintf("%s", o.fromLiteral))
	}

	if o.namespace != nil {
		renderedArgs = append(renderedArgs, "--namespace", fmt.Sprintf("%s", ptrhelpers.StringValue(o.namespace)))
	}

	if o.typeName != nil {
		renderedArgs = append(renderedArgs, "--type", fmt.Sprintf("%s", ptrhelpers.StringValue(o.typeName)))
	}

	if o.stackTrace != nil {
		renderedArgs = append(renderedArgs, "--stack-trace")
	}

	return renderedArgs
}
