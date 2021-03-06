// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/k3d.json

package image

import (
	"fmt"
	"strings"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/k3d"
)

// Import image(s) from docker into k3d cluster(s).
//
// If an IMAGE starts with the prefix ‘docker.io/’, then this prefix is stripped internally. That is, ‘docker.io/rancher/k3d-tools:latest’ is treated as ‘rancher/k3d-tools:latest’.
//
// If an IMAGE starts with the prefix ‘library/’ (or ‘docker.io/library/’), then this prefix is stripped internally. That is, ‘library/busybox:latest’ (or ‘docker.io/library/busybox:latest’) are treated as ‘busybox:latest’.
//
// If an IMAGE does not have a version tag, then ‘:latest’ is assumed. That is, ‘rancher/k3d-tools’ is treated as ‘rancher/k3d-tools:latest’.
//
// A file ARCHIVE always takes precedence. So if a file ‘./rancher/k3d-tools’ exists, k3d will try to import it instead of the IMAGE of the same name.
func Import(imageOrArchive string, opts ...ImportOpt) *tools.CommandContext {
	options := &importOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "image", "import")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, imageOrArchive)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(k3d.K3dToolInfo, defaults.String(options.ToolVersion, k3d.DefaultToolVersion), options.ToolProvider, k3d.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Import image(s) from docker into k3d cluster(s).
//
// If an IMAGE starts with the prefix ‘docker.io/’, then this prefix is stripped internally. That is, ‘docker.io/rancher/k3d-tools:latest’ is treated as ‘rancher/k3d-tools:latest’.
//
// If an IMAGE starts with the prefix ‘library/’ (or ‘docker.io/library/’), then this prefix is stripped internally. That is, ‘library/busybox:latest’ (or ‘docker.io/library/busybox:latest’) are treated as ‘busybox:latest’.
//
// If an IMAGE does not have a version tag, then ‘:latest’ is assumed. That is, ‘rancher/k3d-tools’ is treated as ‘rancher/k3d-tools:latest’.
//
// A file ARCHIVE always takes precedence. So if a file ‘./rancher/k3d-tools’ exists, k3d will try to import it instead of the IMAGE of the same name.
func ImportWithImagesOrArchives(imagesOrArchives []string, opts ...ImportOpt) *tools.CommandContext {
	options := &importOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "image", "import")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, imagesOrArchives...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(k3d.K3dToolInfo, defaults.String(options.ToolVersion, k3d.DefaultToolVersion), options.ToolProvider, k3d.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type ImportOpt func(*importOptions)

func SetToolProvider(p tools.ToolProvider) ImportOpt {
	return func(opts *importOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) ImportOpt {
	return func(opts *importOptions) {
		opts.ToolVersion = v
	}
}

// Select clusters to load the image to. (default [k3s-default])
func Cluster(value ...string) ImportOpt {
	return func(opts *importOptions) {
		opts.cluster = append(opts.cluster, value...)
	}
}

// Do not delete the tarball containing the saved images from the shared volume
func KeepTarball() ImportOpt {
	return func(opts *importOptions) {
		opts.keepTarball = ptrhelpers.Bool(true)
	}
}

// Do not delete the tools node after import
func KeepTools() ImportOpt {
	return func(opts *importOptions) {
		opts.keepTools = ptrhelpers.Bool(true)
	}
}

// Enable Log timestamps
func Timestamps() ImportOpt {
	return func(opts *importOptions) {
		opts.timestamps = ptrhelpers.Bool(true)
	}
}

// Enable super verbose output (trace logging)
func Trace() ImportOpt {
	return func(opts *importOptions) {
		opts.trace = ptrhelpers.Bool(true)
	}
}

// Enable verbose output (debug logging)
func Verbose() ImportOpt {
	return func(opts *importOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

type importOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	cluster      []string
	keepTarball  *bool
	keepTools    *bool
	timestamps   *bool
	trace        *bool
	verbose      *bool
}

func (o *importOptions) toArgs() []string {
	var renderedArgs []string

	if o.cluster != nil {
		renderedArgs = append(renderedArgs, "--cluster", fmt.Sprintf("%s", strings.Join(o.cluster, ",")))
	}

	if o.keepTarball != nil {
		renderedArgs = append(renderedArgs, "--keep-tarball")
	}

	if o.keepTools != nil {
		renderedArgs = append(renderedArgs, "--keep-tools")
	}

	if o.timestamps != nil {
		renderedArgs = append(renderedArgs, "--timestamps")
	}

	if o.trace != nil {
		renderedArgs = append(renderedArgs, "--trace")
	}

	if o.verbose != nil {
		renderedArgs = append(renderedArgs, "--verbose")
	}

	return renderedArgs
}
