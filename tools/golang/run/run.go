// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/go.json

package run

import (
	"fmt"
	"strings"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/golang"
)

// Run compiles and runs the named main Go package.
// Typically the package is specified as a list of .go source files from a single directory,
// but it may also be an import path, file system path, or pattern
// matching a single known package, as in 'go run .' or 'go run my/cmd'.
//
// The exit status of Run is not the exit status of the compiled binary.
func Run(pkg string, opts ...RunOpt) *tools.CommandContext {
	options := &runOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "run")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, pkg)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

// Run compiles and runs the named main Go package.
// Typically the package is specified as a list of .go source files from a single directory,
// but it may also be an import path, file system path, or pattern
// matching a single known package, as in 'go run .' or 'go run my/cmd'.
//
// The exit status of Run is not the exit status of the compiled binary.
func RunWithArgs(pkg string, args []string, opts ...RunOpt) *tools.CommandContext {
	options := &runOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "run")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, pkg)

	_args = append(_args, args...)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(golang.GoToolInfo, defaults.String(options.ToolVersion, golang.DefaultToolVersion), options.ToolProvider, golang.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type RunOpt func(*runOptions)

func SetToolProvider(p tools.ToolProvider) RunOpt {
	return func(opts *runOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) RunOpt {
	return func(opts *runOptions) {
		opts.ToolVersion = v
	}
}

// By default, 'go run' runs the compiled binary directly: 'a.out arguments...'.
// If the -exec flag is given, 'go run' invokes the binary using xprog:
// 	'xprog a.out arguments...'.
// If the -exec flag is not given, GOOS or GOARCH is different from the system
// default, and a program named go_$GOOS_$GOARCH_exec can be found
// on the current search path, 'go run' invokes the binary using that program,
// for example 'go_js_wasm_exec a.out arguments...'. This allows execution of
// cross-compiled programs when a simulator or other execution method is
// available.
func Exec(value string) RunOpt {
	return func(opts *runOptions) {
		opts.exec = &value
	}
}

// force rebuilding of packages that are already up-to-date.
func ForceRebuild() RunOpt {
	return func(opts *runOptions) {
		opts.forceRebuild = ptrhelpers.Bool(true)
	}
}

// the number of programs, such as build commands or
// test binaries, that can be run in parallel.
// The default is the number of CPUs available.
func Workers(value int) RunOpt {
	return func(opts *runOptions) {
		opts.workers = &value
	}
}

// enable data race detection.
// Supported only on linux/amd64, freebsd/amd64, darwin/amd64, windows/amd64,
// linux/ppc64le and linux/arm64 (only for 48-bit VMA).
func Race() RunOpt {
	return func(opts *runOptions) {
		opts.race = ptrhelpers.Bool(true)
	}
}

// enable interoperation with memory sanitizer.
// Supported only on linux/amd64, linux/arm64
// and only with Clang/LLVM as the host C compiler.
// On linux/arm64, pie build mode will be used.
func EnableMemorySanitizer() RunOpt {
	return func(opts *runOptions) {
		opts.msan = ptrhelpers.Bool(true)
	}
}

// print the names of packages as they are compiled.
func Verbose() RunOpt {
	return func(opts *runOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

// print the name of the temporary work directory and
// do not delete it when exiting.
func Work() RunOpt {
	return func(opts *runOptions) {
		opts.work = ptrhelpers.Bool(true)
	}
}

// value '[pattern=]arg list'
// arguments to pass on each go tool asm invocation.
func ASMFlags(value string) RunOpt {
	return func(opts *runOptions) {
		opts.asmflags = &value
	}
}

// build mode to use. See 'go help buildmode' for more.
func BuildMode(value string) RunOpt {
	return func(opts *runOptions) {
		opts.buildmode = &value
	}
}

// name of compiler to use, as in runtime.Compiler (gccgo or gc).
func Compiler(value string) RunOpt {
	return func(opts *runOptions) {
		opts.compiler = &value
	}
}

// value '[pattern=]arg list'
// arguments to pass on each gccgo compiler/linker invocation.
func GCCGoFlags(value string) RunOpt {
	return func(opts *runOptions) {
		opts.gccgoflags = &value
	}
}

// value '[pattern=]arg list'
// arguments to pass on each go tool compile invocation.
func GCCFlags(value string) RunOpt {
	return func(opts *runOptions) {
		opts.gccflags = &value
	}
}

// a suffix to use in the name of the package installation directory,
// in order to keep output separate from default builds.
// If using the -race flag, the install suffix is automatically set to race
// or, if set explicitly, has _race appended to it. Likewise for the -msan
// flag. Using a -buildmode option that requires non-default compile flags
// has a similar effect.
func InstallSuffix(value string) RunOpt {
	return func(opts *runOptions) {
		opts.installSuffix = &value
	}
}

// value '[pattern=]arg list'
// arguments to pass on each go tool link invocation.
func LDFlags(value string) RunOpt {
	return func(opts *runOptions) {
		opts.ldflags = &value
	}
}

// build code that will be linked against shared libraries previously
// created with -buildmode=shared.
func LinkShared() RunOpt {
	return func(opts *runOptions) {
		opts.linkShared = ptrhelpers.Bool(true)
	}
}

// module download mode to use: readonly, vendor, or mod.
// See 'go help modules' for more.
func Mod(value string) RunOpt {
	return func(opts *runOptions) {
		opts.mod = &value
	}
}

// leave newly-created directories in the module cache read-write
// instead of making them read-only.
func ModCacheRW() RunOpt {
	return func(opts *runOptions) {
		opts.modcacherw = ptrhelpers.Bool(true)
	}
}

// in module aware mode, read (and possibly write) an alternate go.mod
// file instead of the one in the module root directory. A file named
// "go.mod" must still be present in order to determine the module root
// directory, but it is not accessed. When -modfile is specified, an
// alternate go.sum file is also used: its path is derived from the
// -modfile flag by trimming the ".mod" extension and appending ".sum".
func ModFile(value string) RunOpt {
	return func(opts *runOptions) {
		opts.modfile = &value
	}
}

// install and load all packages from dir instead of the usual locations.
// For example, when building with a non-standard configuration,
// use -pkgdir to keep generated packages in a separate location
func PackageDir(value string) RunOpt {
	return func(opts *runOptions) {
		opts.pkgdir = &value
	}
}

// a comma-separated list of build tags to consider satisfied during the
// build. For more information about build tags, see the description of
// build constraints in the documentation for the go/build package.
// (Earlier versions of Go used a space-separated list, and that form
// is deprecated but still recognized.)
// NOTE: comma separation is handled automatically
func Tags(value ...string) RunOpt {
	return func(opts *runOptions) {
		opts.tags = append(opts.tags, value...)
	}
}

// remove all file system paths from the resulting executable.
// Instead of absolute file system paths, the recorded file names
// will begin with either "go" (for the standard library),
// or a module path@version (when using modules),
// or a plain import path (when using GOPATH).
func TrimPath() RunOpt {
	return func(opts *runOptions) {
		opts.trimpath = ptrhelpers.Bool(true)
	}
}

// a program to use to invoke toolchain programs like vet and asm.
// For example, instead of running asm, the go command will run
// 'cmd args /path/to/asm <arguments for asm>'.
func ToolExec(value string) RunOpt {
	return func(opts *runOptions) {
		opts.toolexec = &value
	}
}

type runOptions struct {
	ToolProvider  tools.ToolProvider
	ToolVersion   string
	exec          *string
	forceRebuild  *bool
	workers       *int
	race          *bool
	msan          *bool
	verbose       *bool
	work          *bool
	asmflags      *string
	buildmode     *string
	compiler      *string
	gccgoflags    *string
	gccflags      *string
	installSuffix *string
	ldflags       *string
	linkShared    *bool
	mod           *string
	modcacherw    *bool
	modfile       *string
	pkgdir        *string
	tags          []string
	trimpath      *bool
	toolexec      *string
}

func (o *runOptions) toArgs() []string {
	var renderedArgs []string

	if o.exec != nil {
		renderedArgs = append(renderedArgs, "-exec", fmt.Sprintf("%s", ptrhelpers.StringValue(o.exec)))
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

	return renderedArgs
}
