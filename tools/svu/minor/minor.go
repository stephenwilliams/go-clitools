// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/svu.json

package minor

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/svu"
)

// new minor version
func Minor(opts ...MinorOpt) *tools.CommandContext {
	options := &minorOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "minor")

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(svu.SVUToolInfo, defaults.String(options.ToolVersion, svu.DefaultToolVersion), options.ToolProvider, svu.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type MinorOpt func(*minorOptions)

func SetToolProvider(p tools.ToolProvider) MinorOpt {
	return func(opts *minorOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) MinorOpt {
	return func(opts *minorOptions) {
		opts.ToolVersion = v
	}
}

// discards pre-release and build metadata if set to false
func Metadata(value bool) MinorOpt {
	return func(opts *minorOptions) {
		opts.metadata = &value
	}
}

// discards pre-release metadata if set to false
func PreRelease(value bool) MinorOpt {
	return func(opts *minorOptions) {
		opts.preRelease = &value
	}
}

// discards build metadata if set to false
func Build(value bool) MinorOpt {
	return func(opts *minorOptions) {
		opts.build = &value
	}
}

// determines if latest tag of the current or all branches will be used
func TagMode(value string) MinorOpt {
	return func(opts *minorOptions) {
		opts.tagMode = &value
	}
}

type minorOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	metadata     *bool
	preRelease   *bool
	build        *bool
	tagMode      *string
}

func (o *minorOptions) toArgs() []string {
	var renderedArgs []string

	if o.metadata != nil {
		renderedArgs = append(renderedArgs, "--metadata", fmt.Sprintf("%t", ptrhelpers.BoolValue(o.metadata)))
	}

	if o.preRelease != nil {
		renderedArgs = append(renderedArgs, "--pre-release", fmt.Sprintf("%t", ptrhelpers.BoolValue(o.preRelease)))
	}

	if o.build != nil {
		renderedArgs = append(renderedArgs, "--build", fmt.Sprintf("%t", ptrhelpers.BoolValue(o.build)))
	}

	if o.tagMode != nil {
		renderedArgs = append(renderedArgs, "--tag-mode", fmt.Sprintf("%s", ptrhelpers.StringValue(o.tagMode)))
	}

	return renderedArgs
}
