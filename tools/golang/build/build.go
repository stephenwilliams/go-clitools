// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package build

import (
	"fmt"
	"strings"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

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
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

func BuildWithPackage(pkg string, opts ...BuildOpt) *tools.CommandContext {
	options := &buildOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "build")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, pkg)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

func BuildWithPackages(packages []string, opts ...BuildOpt) *tools.CommandContext {
	options := &buildOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "build")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, packages...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
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

// forces build to write the resulting executable or object
// to the named output file or directory, instead of the default behavior described
// in the last two paragraphs. If the named output is a directory that exists,
// then any resulting executables will be written to that directory.
func Output(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.output = &value
	}
}

// force rebuilding of packages that are already up-to-date.
func ForceRebuild() BuildOpt {
	return func(opts *buildOptions) {
		opts.forceRebuild = ptrhelpers.Bool(true)
	}
}

// the number of programs, such as build commands or
// test binaries, that can be run in parallel.
// The default is the number of CPUs available.
func Workers(value int) BuildOpt {
	return func(opts *buildOptions) {
		opts.workers = &value
	}
}

// enable data race detection.
// Supported only on linux/amd64, freebsd/amd64, darwin/amd64, windows/amd64,
// linux/ppc64le and linux/arm64 (only for 48-bit VMA).
func Race() BuildOpt {
	return func(opts *buildOptions) {
		opts.race = ptrhelpers.Bool(true)
	}
}

// enable interoperation with memory sanitizer.
// Supported only on linux/amd64, linux/arm64
// and only with Clang/LLVM as the host C compiler.
// On linux/arm64, pie build mode will be used.
func EnableMemorySanitizer() BuildOpt {
	return func(opts *buildOptions) {
		opts.msan = ptrhelpers.Bool(true)
	}
}

// print the names of packages as they are compiled.
func Verbose() BuildOpt {
	return func(opts *buildOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

// print the name of the temporary work directory and
// do not delete it when exiting.
func Work() BuildOpt {
	return func(opts *buildOptions) {
		opts.work = ptrhelpers.Bool(true)
	}
}

// value '[pattern=]arg list'
// arguments to pass on each go tool asm invocation.
func ASMFlags(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.asmflags = &value
	}
}

// build mode to use. See 'go help buildmode' for more.
func BuildMode(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.buildmode = &value
	}
}

// name of compiler to use, as in runtime.Compiler (gccgo or gc).
func Compiler(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.compiler = &value
	}
}

// value '[pattern=]arg list'
// arguments to pass on each gccgo compiler/linker invocation.
func GCCGoFlags(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.gccgoflags = &value
	}
}

// value '[pattern=]arg list'
// arguments to pass on each go tool compile invocation.
func GCCFlags(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.gccflags = &value
	}
}

// a suffix to use in the name of the package installation directory,
// in order to keep output separate from default builds.
// If using the -race flag, the install suffix is automatically set to race
// or, if set explicitly, has _race appended to it. Likewise for the -msan
// flag. Using a -buildmode option that requires non-default compile flags
// has a similar effect.
func InstallSuffix(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.installSuffix = &value
	}
}

// value '[pattern=]arg list'
// arguments to pass on each go tool link invocation.
func LDFlags(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.ldflags = &value
	}
}

// build code that will be linked against shared libraries previously
// created with -buildmode=shared.
func LinkShared() BuildOpt {
	return func(opts *buildOptions) {
		opts.linkShared = ptrhelpers.Bool(true)
	}
}

// module download mode to use: readonly, vendor, or mod.
// See 'go help modules' for more.
func Mod(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.mod = &value
	}
}

// leave newly-created directories in the module cache read-write
// instead of making them read-only.
func ModCacheRW() BuildOpt {
	return func(opts *buildOptions) {
		opts.modcacherw = ptrhelpers.Bool(true)
	}
}

// in module aware mode, read (and possibly write) an alternate go.mod
// file instead of the one in the module root directory. A file named
// "go.mod" must still be present in order to determine the module root
// directory, but it is not accessed. When -modfile is specified, an
// alternate go.sum file is also used: its path is derived from the
// -modfile flag by trimming the ".mod" extension and appending ".sum".
func ModFile(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.modfile = &value
	}
}

// install and load all packages from dir instead of the usual locations.
// For example, when building with a non-standard configuration,
// use -pkgdir to keep generated packages in a separate location
func PackageDir(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.pkgdir = &value
	}
}

// a comma-separated list of build tags to consider satisfied during the
// build. For more information about build tags, see the description of
// build constraints in the documentation for the go/build package.
// (Earlier versions of Go used a space-separated list, and that form
// is deprecated but still recognized.)
// NOTE: comma separation is handled automatically
func Tags(value ...string) BuildOpt {
	return func(opts *buildOptions) {
		opts.tags = value
	}
}

// remove all file system paths from the resulting executable.
// Instead of absolute file system paths, the recorded file names
// will begin with either "go" (for the standard library),
// or a module path@version (when using modules),
// or a plain import path (when using GOPATH).
func TrimPath() BuildOpt {
	return func(opts *buildOptions) {
		opts.trimpath = ptrhelpers.Bool(true)
	}
}

// a program to use to invoke toolchain programs like vet and asm.
// For example, instead of running asm, the go command will run
// 'cmd args /path/to/asm <arguments for asm>'.
func ToolExec(value string) BuildOpt {
	return func(opts *buildOptions) {
		opts.toolexec = &value
	}
}

// print the commands but do not run them.
func PrintCommandsOnly() BuildOpt {
	return func(opts *buildOptions) {
		opts.printCommandsOnly = ptrhelpers.Bool(true)
	}
}

// print the commands.
func PrintCommands() BuildOpt {
	return func(opts *buildOptions) {
		opts.printCommands = ptrhelpers.Bool(true)
	}
}

type buildOptions struct {
	ToolProvider      tools.ToolProvider
	ToolVersion       string
	output            *string
	forceRebuild      *bool
	workers           *int
	race              *bool
	msan              *bool
	verbose           *bool
	work              *bool
	asmflags          *string
	buildmode         *string
	compiler          *string
	gccgoflags        *string
	gccflags          *string
	installSuffix     *string
	ldflags           *string
	linkShared        *bool
	mod               *string
	modcacherw        *bool
	modfile           *string
	pkgdir            *string
	tags              []string
	trimpath          *bool
	toolexec          *string
	printCommandsOnly *bool
	printCommands     *bool
}

func (o *buildOptions) toArgs() []string {
	var renderedArgs []string

	if o.output != nil {
		renderedArgs = append(renderedArgs, "-o", fmt.Sprintf("%s", ptrhelpers.StringValue(o.output)))
	}

	if o.forceRebuild != nil {
		renderedArgs = append(renderedArgs, "-a")
	}

	if o.workers != nil {
		renderedArgs = append(renderedArgs, "-p", fmt.Sprintf("%d", ptrhelpers.IntValue(o.workers)))
	}

	if o.race != nil {
		renderedArgs = append(renderedArgs, "-race")
	}

	if o.msan != nil {
		renderedArgs = append(renderedArgs, "-msan")
	}

	if o.verbose != nil {
		renderedArgs = append(renderedArgs, "-v")
	}

	if o.work != nil {
		renderedArgs = append(renderedArgs, "-work")
	}

	if o.asmflags != nil {
		renderedArgs = append(renderedArgs, "-asmflags", fmt.Sprintf("%s", ptrhelpers.StringValue(o.asmflags)))
	}

	if o.buildmode != nil {
		renderedArgs = append(renderedArgs, "-buildmode", fmt.Sprintf("%s", ptrhelpers.StringValue(o.buildmode)))
	}

	if o.compiler != nil {
		renderedArgs = append(renderedArgs, "-compiler", fmt.Sprintf("%s", ptrhelpers.StringValue(o.compiler)))
	}

	if o.gccgoflags != nil {
		renderedArgs = append(renderedArgs, "-gccgoflags", fmt.Sprintf("%s", ptrhelpers.StringValue(o.gccgoflags)))
	}

	if o.gccflags != nil {
		renderedArgs = append(renderedArgs, "-gccflags", fmt.Sprintf("%s", ptrhelpers.StringValue(o.gccflags)))
	}

	if o.installSuffix != nil {
		renderedArgs = append(renderedArgs, "-installsuffix", fmt.Sprintf("%s", ptrhelpers.StringValue(o.installSuffix)))
	}

	if o.ldflags != nil {
		renderedArgs = append(renderedArgs, "-ldflags", fmt.Sprintf("%s", ptrhelpers.StringValue(o.ldflags)))
	}

	if o.linkShared != nil {
		renderedArgs = append(renderedArgs, "-linkshared")
	}

	if o.mod != nil {
		renderedArgs = append(renderedArgs, "-mod", fmt.Sprintf("%s", ptrhelpers.StringValue(o.mod)))
	}

	if o.modcacherw != nil {
		renderedArgs = append(renderedArgs, "-modcacherw")
	}

	if o.modfile != nil {
		renderedArgs = append(renderedArgs, "-modfile", fmt.Sprintf("%s", ptrhelpers.StringValue(o.modfile)))
	}

	if o.pkgdir != nil {
		renderedArgs = append(renderedArgs, "-pkgdir", fmt.Sprintf("%s", ptrhelpers.StringValue(o.pkgdir)))
	}

	if o.tags != nil {
		renderedArgs = append(renderedArgs, "-tags", fmt.Sprintf("%s", strings.Join(o.tags, ",")))
	}

	if o.trimpath != nil {
		renderedArgs = append(renderedArgs, "-trimpath")
	}

	if o.toolexec != nil {
		renderedArgs = append(renderedArgs, "-toolexec", fmt.Sprintf("%s", ptrhelpers.StringValue(o.toolexec)))
	}

	if o.printCommandsOnly != nil {
		renderedArgs = append(renderedArgs, "-n")
	}

	if o.printCommands != nil {
		renderedArgs = append(renderedArgs, "-x")
	}

	return renderedArgs
}