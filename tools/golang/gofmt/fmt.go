// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package gofmt

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

// Fmt runs the command 'gofmt -l -w' on the packages named
// by the import paths. It prints the names of the files that are modified.
//
// For more about gofmt, see 'go doc cmd/gofmt'.
// For more about specifying packages, see 'go help packages'.
func Fmt(pkg string, opts ...FmtOpt) *tools.CommandContext {
	options := &fmtOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "fmt")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, pkg)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Fmt runs the command 'gofmt -l -w' on the packages named
// by the import paths. It prints the names of the files that are modified.
//
// For more about gofmt, see 'go doc cmd/gofmt'.
// For more about specifying packages, see 'go help packages'.
func FmtwithPackages(pkgs []string, opts ...FmtOpt) *tools.CommandContext {
	options := &fmtOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "fmt")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, pkgs...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type FmtOpt func(*fmtOptions)

func SetToolProvider(p tools.ToolProvider) FmtOpt {
	return func(opts *fmtOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) FmtOpt {
	return func(opts *fmtOptions) {
		opts.ToolVersion = v
	}
}

// sets which module download mode
// to use: readonly or vendor. See 'go help modules' for more.
func Mod(value string) FmtOpt {
	return func(opts *fmtOptions) {
		opts.mod = &value
	}
}

// print the commands but do not run them.
func PrintCommandsOnly() FmtOpt {
	return func(opts *fmtOptions) {
		opts.printCommandsOnly = ptrhelpers.Bool(true)
	}
}

// print the commands.
func PrintCommands() FmtOpt {
	return func(opts *fmtOptions) {
		opts.printCommands = ptrhelpers.Bool(true)
	}
}

type fmtOptions struct {
	ToolProvider      tools.ToolProvider
	ToolVersion       string
	mod               *string
	printCommandsOnly *bool
	printCommands     *bool
}

func (o *fmtOptions) toArgs() []string {
	var renderedArgs []string

	if o.mod != nil {
		renderedArgs = append(renderedArgs, "-mod", fmt.Sprintf("%s", ptrhelpers.StringValue(o.mod)))
	}

	if o.printCommandsOnly != nil {
		renderedArgs = append(renderedArgs, "-n")
	}

	if o.printCommands != nil {
		renderedArgs = append(renderedArgs, "-x")
	}

	return renderedArgs
}
