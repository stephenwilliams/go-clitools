// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package edit

import (
	"fmt"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

// Edit provides a command-line interface for editing go.mod,
// for use primarily by tools or scripts. It reads only go.mod;
// it does not look up information about the modules involved.
// By default, edit reads and writes the go.mod file of the main module,
// but a different target file can be specified after the editing flags.
//
// The editing flags specify a sequence of editing operations.
func Edit(opts ...EditOpt) *tools.CommandContext {
	options := &editOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string

	_args = append(_args, options.toArgs()...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Edit provides a command-line interface for editing go.mod,
// for use primarily by tools or scripts. It reads only go.mod;
// it does not look up information about the modules involved.
// By default, edit reads and writes the go.mod file of the main module,
// but a different target file can be specified after the editing flags.
//
// The editing flags specify a sequence of editing operations.
func EditWithGoModPath(goModPath string, opts ...EditOpt) *tools.CommandContext {
	options := &editOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string

	_args = append(_args, options.toArgs()...)
	_args = append(_args, goModPath)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type EditOpt func(*editOptions)

func SetToolProvider(p tools.ToolProvider) EditOpt {
	return func(opts *editOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) EditOpt {
	return func(opts *editOptions) {
		opts.ToolVersion = v
	}
}

// reformats the go.mod file without making other changes.
// This reformatting is also implied by any other modifications that use or
// rewrite the go.mod file. The only time this flag is needed is if no other
// flags are specified, as in 'go mod edit -fmt'.
func Fmt() EditOpt {
	return func(opts *editOptions) {
		opts.fmt = ptrhelpers.Bool(true)
	}
}

// changes the module's path (the go.mod file's module line).
func Module() EditOpt {
	return func(opts *editOptions) {
		opts.module = ptrhelpers.Bool(true)
	}
}

// The -require=path@version and -droprequire=path flags
// add and drop a requirement on the given module path and version.
// Note that -require overrides any existing requirements on path.
// These flags are mainly for tools that understand the module graph.
// Users should prefer 'go get path@version' or 'go get path@none',
// which make other go.mod adjustments as needed to satisfy
// constraints imposed by other modules.
func Require(value string) EditOpt {
	return func(opts *editOptions) {
		opts.require = &value
	}
}

// The -require=path@version and -droprequire=path flags
// add and drop a requirement on the given module path and version.
// Note that -require overrides any existing requirements on path.
// These flags are mainly for tools that understand the module graph.
// Users should prefer 'go get path@version' or 'go get path@none',
// which make other go.mod adjustments as needed to satisfy
// constraints imposed by other modules.
func DropRequire(value string) EditOpt {
	return func(opts *editOptions) {
		opts.dropRequire = &value
	}
}

// The -exclude=path@version and -dropexclude=path@version flags
// add and drop an exclusion for the given module path and version.
// Note that -exclude=path@version is a no-op if that exclusion already exists.
func Exclude(value string) EditOpt {
	return func(opts *editOptions) {
		opts.exclude = &value
	}
}

// The -exclude=path@version and -dropexclude=path@version flags
// add and drop an exclusion for the given module path and version.
// Note that -exclude=path@version is a no-op if that exclusion already exists.
func DropExclude(value string) EditOpt {
	return func(opts *editOptions) {
		opts.dropExclude = &value
	}
}

// The -replace=old[@v]=new[@v] flag adds a replacement of the given
// module path and version pair. If the @v in old@v is omitted, a
// replacement without a version on the left side is added, which applies
// to all versions of the old module path. If the @v in new@v is omitted,
// the new path should be a local module root directory, not a module
// path. Note that -replace overrides any redundant replacements for old[@v],
// so omitting @v will drop existing replacements for specific versions.
func Replace(value string) EditOpt {
	return func(opts *editOptions) {
		opts.replace = &value
	}
}

// The -dropreplace=old[@v] flag drops a replacement of the given
// module path and version pair. If the @v is omitted, a replacement without
// a version on the left side is dropped.
func DropReplace(value string) EditOpt {
	return func(opts *editOptions) {
		opts.dropReplace = &value
	}
}

// sets the expected Go language version.
func GoVersion(value string) EditOpt {
	return func(opts *editOptions) {
		opts.goVersion = &value
	}
}

// prints the final go.mod in its text format instead of
// writing it back to go.mod.
func Print() EditOpt {
	return func(opts *editOptions) {
		opts.print = ptrhelpers.Bool(true)
	}
}

// prints the final go.mod file in JSON format instead of
// writing it back to go.mod. The JSON output corresponds to these Go types:
//
// 	type Module struct {
// 		Path string
// 		Version string
// 	}
//
// 	type GoMod struct {
// 		Module  Module
// 		Go      string
// 		Require []Require
// 		Exclude []Module
// 		Replace []Replace
// 	}
//
// 	type Require struct {
// 		Path string
// 		Version string
// 		Indirect bool
// 	}
//
// 	type Replace struct {
// 		Old Module
// 		New Module
// 	}
//
// Note that this only describes the go.mod file itself, not other modules
// referred to indirectly. For the full set of modules available to a build,
// use 'go list -m -json all'.
//
// For example, a tool can obtain the go.mod as a data structure by
// parsing the output of 'go mod edit -json' and can then make changes
// by invoking 'go mod edit' with -require, -exclude, and so on.
func JSON() EditOpt {
	return func(opts *editOptions) {
		opts.json = ptrhelpers.Bool(true)
	}
}

type editOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	fmt          *bool
	module       *bool
	require      *string
	dropRequire  *string
	exclude      *string
	dropExclude  *string
	replace      *string
	dropReplace  *string
	goVersion    *string
	print        *bool
	json         *bool
}

func (o *editOptions) toArgs() []string {
	var renderedArgs []string

	if o.fmt != nil {
		renderedArgs = append(renderedArgs, "-fmt")
	}

	if o.module != nil {
		renderedArgs = append(renderedArgs, "-module")
	}

	if o.require != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-require=%s", ptrhelpers.StringValue(o.require)))
	}

	if o.dropRequire != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-droprequire=%s", ptrhelpers.StringValue(o.dropRequire)))
	}

	if o.exclude != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-exclude=%s", ptrhelpers.StringValue(o.exclude)))
	}

	if o.dropExclude != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-dropexclude=%s", ptrhelpers.StringValue(o.dropExclude)))
	}

	if o.replace != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-replace=%s", ptrhelpers.StringValue(o.replace)))
	}

	if o.dropReplace != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-dropreplace=%s", ptrhelpers.StringValue(o.dropReplace)))
	}

	if o.goVersion != nil {
		renderedArgs = append(renderedArgs, fmt.Sprintf("-go=%s", ptrhelpers.StringValue(o.goVersion)))
	}

	if o.print != nil {
		renderedArgs = append(renderedArgs, "-print")
	}

	if o.json != nil {
		renderedArgs = append(renderedArgs, "-json")
	}

	return renderedArgs
}