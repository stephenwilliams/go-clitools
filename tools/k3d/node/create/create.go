// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/k3d.json

package create

import (
	"fmt"
	"time"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/k3d"
)

// Create a new k3s node in docker
func Create(name string, opts ...CreateOpt) *tools.CommandContext {
	options := &createOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "node", "create")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, name)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(k3d.K3dToolInfo, defaults.String(options.ToolVersion, k3d.DefaultToolVersion), options.ToolProvider, k3d.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type CreateOpt func(*createOptions)

func SetToolProvider(p tools.ToolProvider) CreateOpt {
	return func(opts *createOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) CreateOpt {
	return func(opts *createOptions) {
		opts.ToolVersion = v
	}
}

// Select the cluster that the node shall connect to. (default "k3s-default")
func Cluster(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.cluster = &value
	}
}

// Specify k3s image used for the node(s) (default "docker.io/rancher/k3s:v1.20.0-k3s2")
func Image(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.image = &value
	}
}

// Memory limit imposed on the node [From docker]
func Memory(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.memory = &value
	}
}

// Number of replicas of this node specification. (default 1)
func Replicas(value int) CreateOpt {
	return func(opts *createOptions) {
		opts.replicas = &value
	}
}

// Specify node role [server, agent] (default "agent")
func Role(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.role = &value
	}
}

// Maximum waiting time for '--wait' before canceling/returning.
func Timeout(value time.Duration) CreateOpt {
	return func(opts *createOptions) {
		opts.timeout = &value
	}
}

// Wait for the node(s) to be ready before returning. (default true)
func Wait(value bool) CreateOpt {
	return func(opts *createOptions) {
		opts.wait = &value
	}
}

// Enable Log timestamps
func Timestamps() CreateOpt {
	return func(opts *createOptions) {
		opts.timestamps = ptrhelpers.Bool(true)
	}
}

// Enable super verbose output (trace logging)
func Trace() CreateOpt {
	return func(opts *createOptions) {
		opts.trace = ptrhelpers.Bool(true)
	}
}

// Enable verbose output (debug logging)
func Verbose() CreateOpt {
	return func(opts *createOptions) {
		opts.verbose = ptrhelpers.Bool(true)
	}
}

type createOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
	cluster      *string
	image        *string
	memory       *string
	replicas     *int
	role         *string
	timeout      *time.Duration
	wait         *bool
	timestamps   *bool
	trace        *bool
	verbose      *bool
}

func (o *createOptions) toArgs() []string {
	var renderedArgs []string

	if o.cluster != nil {
		renderedArgs = append(renderedArgs, "--cluster", fmt.Sprintf("%s", ptrhelpers.StringValue(o.cluster)))
	}

	if o.image != nil {
		renderedArgs = append(renderedArgs, "--image", fmt.Sprintf("%s", ptrhelpers.StringValue(o.image)))
	}

	if o.memory != nil {
		renderedArgs = append(renderedArgs, "--memory", fmt.Sprintf("%s", ptrhelpers.StringValue(o.memory)))
	}

	if o.replicas != nil {
		renderedArgs = append(renderedArgs, "--replicas", fmt.Sprintf("%d", ptrhelpers.IntValue(o.replicas)))
	}

	if o.role != nil {
		renderedArgs = append(renderedArgs, "--role", fmt.Sprintf("%s", ptrhelpers.StringValue(o.role)))
	}

	if o.timeout != nil {
		renderedArgs = append(renderedArgs, "--timeout", fmt.Sprintf("%v", o.timeout))
	}

	if o.wait != nil {
		renderedArgs = append(renderedArgs, "--wait", fmt.Sprintf("%t", ptrhelpers.BoolValue(o.wait)))
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
