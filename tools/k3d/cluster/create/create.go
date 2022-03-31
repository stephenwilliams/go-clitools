// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/k3d.json

package create

import (
	"fmt"
	"strings"
	"time"

	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/k3d"
)

// Create a new k3s cluster with containerized nodes (k3s in docker). Every cluster will consist of one or more containers: - 1 (or more) server node container (k3s) - (optionally) 1 loadbalancer container as the entrypoint to the cluster (nginx) - (optionally) 1 (or more) agent node containers (k3s)
func Create(name string, opts ...CreateOpt) *tools.CommandContext {
	options := &createOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "cluster", "create")

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

// Specify how many agents you want to create
func Agents(value int) CreateOpt {
	return func(opts *createOptions) {
		opts.agents = &value
	}
}

// Memory limit imposed on the agents nodes [From docker]
func AgentsMemory(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.agentsMemory = &value
	}
}

// Specify the Kubernetes API server port exposed on the LoadBalancer (Format: [HOST:]HOSTPORT)
func APIPort(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.apiPort = &value
	}
}

// Path of a config file to use
func Config(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.config = &value
	}
}

// Add environment variables to nodes (Format: KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]
func Env(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.env = value
	}
}

// GPU devices to add to the cluster node containers ('all' to pass all GPUs) [From docker]
func GPUs(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.gpus = &value
	}
}

// Specify k3s image that you want to use for the nodes
func Image(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.image = &value
	}
}

// Additional args passed to the k3s agent command on agent nodes
func K3sAgentArgs(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.k3sAgentArgs = value
	}
}

// Additional args passed to the k3s server command on server nodes
func K3sServerArgs(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.k3sServerArgs = value
	}
}

// Directly switch the default kubeconfig's current-context to the new cluster's context (requires --kubeconfig-update-default) (default true)
func KubeconfigSwitchContext(value bool) CreateOpt {
	return func(opts *createOptions) {
		opts.KubeconfigSwitchContext = &value
	}
}

// Directly update the default kubeconfig with the new cluster's context (default true)
func KubeconfigUpdateDefault(value bool) CreateOpt {
	return func(opts *createOptions) {
		opts.kubeconfigUpdateDefault = &value
	}
}

// Add label to node container (Format: KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]
func Labels(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.labels = value
	}
}

// Join an existing network
func Network(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.network = &value
	}
}

// Disable the automatic injection of the Host IP as 'host.k3d.internal' into the containers and CoreDNS
func NoHostIP() CreateOpt {
	return func(opts *createOptions) {
		opts.noHostIP = ptrhelpers.Bool(true)
	}
}

// Disable the creation of a volume for importing images
func NoImageVolume() CreateOpt {
	return func(opts *createOptions) {
		opts.noImageVolume = ptrhelpers.Bool(true)
	}
}

// Disable the creation of a LoadBalancer in front of the server nodes
func NoLoadBalancer() CreateOpt {
	return func(opts *createOptions) {
		opts.noLB = ptrhelpers.Bool(true)
	}
}

// Disable the automatic rollback actions, if anything goes wrong
func NoRollback() CreateOpt {
	return func(opts *createOptions) {
		opts.noRollback = ptrhelpers.Bool(true)
	}
}

// Map ports from the node containers to the host (Format: [HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER])
func Port(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.port = value
	}
}

// Specify path to an extra registries.yaml file
func RegistryConfig(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.registryConfig = &value
	}
}

// Create a k3d-managed registry and connect it to the cluster
func RegistryCreate() CreateOpt {
	return func(opts *createOptions) {
		opts.registryCreate = ptrhelpers.Bool(true)
	}
}

// Connect to one or more k3d-managed registries running locally
func RegistryUse(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.registryUse = value
	}
}

// Specify how many servers you want to create
func Servers(value int) CreateOpt {
	return func(opts *createOptions) {
		opts.servers = &value
	}
}

// Memory limit imposed on the server nodes [From docker]
func ServersMemory(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.serversMemory = &value
	}
}

// [Experimental: IPAM] Define a subnet for the newly created container network (Example: 172.28.0.0/16)
func Subnet(value string) CreateOpt {
	return func(opts *createOptions) {
		opts.subnet = &value
	}
}

// Rollback changes if cluster couldn't be created in specified duration.
func Timeout(value time.Duration) CreateOpt {
	return func(opts *createOptions) {
		opts.timeout = &value
	}
}

// Mount volumes into the nodes (Format: [SOURCE:]DEST[@NODEFILTER[;NODEFILTER...]]
func Volume(value ...string) CreateOpt {
	return func(opts *createOptions) {
		opts.volume = value
	}
}

// Wait for the server(s) to be ready before returning. Use '--timeout DURATION' to not wait forever. (default true)
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
	ToolProvider            tools.ToolProvider
	ToolVersion             string
	agents                  *int
	agentsMemory            *string
	apiPort                 *string
	config                  *string
	env                     []string
	gpus                    *string
	image                   *string
	k3sAgentArgs            []string
	k3sServerArgs           []string
	KubeconfigSwitchContext *bool
	kubeconfigUpdateDefault *bool
	labels                  []string
	network                 *string
	noHostIP                *bool
	noImageVolume           *bool
	noLB                    *bool
	noRollback              *bool
	port                    []string
	registryConfig          *string
	registryCreate          *bool
	registryUse             []string
	servers                 *int
	serversMemory           *string
	subnet                  *string
	timeout                 *time.Duration
	volume                  []string
	wait                    *bool
	timestamps              *bool
	trace                   *bool
	verbose                 *bool
}

func (o *createOptions) toArgs() []string {
	var renderedArgs []string

	if o.agents != nil {
		renderedArgs = append(renderedArgs, "--agents", fmt.Sprintf("%d", ptrhelpers.IntValue(o.agents)))
	}

	if o.agentsMemory != nil {
		renderedArgs = append(renderedArgs, "--agents-memory", fmt.Sprintf("%s", ptrhelpers.StringValue(o.agentsMemory)))
	}

	if o.apiPort != nil {
		renderedArgs = append(renderedArgs, "--api-port", fmt.Sprintf("%s", ptrhelpers.StringValue(o.apiPort)))
	}

	if o.config != nil {
		renderedArgs = append(renderedArgs, "--config", fmt.Sprintf("%s", ptrhelpers.StringValue(o.config)))
	}

	if o.env != nil {
		for _, v := range o.env {
			renderedArgs = append(renderedArgs, "--env", fmt.Sprintf("%s", v))
		}
	}

	if o.gpus != nil {
		renderedArgs = append(renderedArgs, "--gpus", fmt.Sprintf("%s", ptrhelpers.StringValue(o.gpus)))
	}

	if o.image != nil {
		renderedArgs = append(renderedArgs, "--image", fmt.Sprintf("%s", ptrhelpers.StringValue(o.image)))
	}

	if o.k3sAgentArgs != nil {
		for _, v := range o.k3sAgentArgs {
			renderedArgs = append(renderedArgs, "--k3s-agent-arg", fmt.Sprintf("%s", v))
		}
	}

	if o.k3sServerArgs != nil {
		for _, v := range o.k3sServerArgs {
			renderedArgs = append(renderedArgs, "--k3s-server-arg", fmt.Sprintf("%s", v))
		}
	}

	if o.KubeconfigSwitchContext != nil {
		renderedArgs = append(renderedArgs, "--kubeconfig-switch-context", fmt.Sprintf("%t", ptrhelpers.BoolValue(o.KubeconfigSwitchContext)))
	}

	if o.kubeconfigUpdateDefault != nil {
		renderedArgs = append(renderedArgs, "--kubeconfig-update-default", fmt.Sprintf("%t", ptrhelpers.BoolValue(o.kubeconfigUpdateDefault)))
	}

	if o.labels != nil {
		for _, v := range o.labels {
			renderedArgs = append(renderedArgs, "--label", fmt.Sprintf("%s", v))
		}
	}

	if o.network != nil {
		renderedArgs = append(renderedArgs, "--network", fmt.Sprintf("%s", ptrhelpers.StringValue(o.network)))
	}

	if o.noHostIP != nil {
		renderedArgs = append(renderedArgs, "--no-hostip")
	}

	if o.noImageVolume != nil {
		renderedArgs = append(renderedArgs, "--no-image-volume")
	}

	if o.noLB != nil {
		renderedArgs = append(renderedArgs, "--no-lb")
	}

	if o.noRollback != nil {
		renderedArgs = append(renderedArgs, "--no-rollback")
	}

	if o.port != nil {
		for _, v := range o.port {
			renderedArgs = append(renderedArgs, "--port", fmt.Sprintf("%s", v))
		}
	}

	if o.registryConfig != nil {
		renderedArgs = append(renderedArgs, "--registry-config", fmt.Sprintf("%s", ptrhelpers.StringValue(o.registryConfig)))
	}

	if o.registryCreate != nil {
		renderedArgs = append(renderedArgs, "--registry-create")
	}

	if o.registryUse != nil {
		renderedArgs = append(renderedArgs, "--registry-use", fmt.Sprintf("%s", strings.Join(o.registryUse, ",")))
	}

	if o.servers != nil {
		renderedArgs = append(renderedArgs, "--servers", fmt.Sprintf("%d", ptrhelpers.IntValue(o.servers)))
	}

	if o.serversMemory != nil {
		renderedArgs = append(renderedArgs, "--servers-memory", fmt.Sprintf("%s", ptrhelpers.StringValue(o.serversMemory)))
	}

	if o.subnet != nil {
		renderedArgs = append(renderedArgs, "--subnet", fmt.Sprintf("%s", ptrhelpers.StringValue(o.subnet)))
	}

	if o.timeout != nil {
		renderedArgs = append(renderedArgs, "--timeout", fmt.Sprintf("%v", o.timeout))
	}

	if o.volume != nil {
		for _, v := range o.volume {
			renderedArgs = append(renderedArgs, "--volume", fmt.Sprintf("%s", v))
		}
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
