// Generated by toolgen. DO NOT EDIT.
// Generated from tool specification:
//   _build/specifications/kustomize.json

package label

import (
	"github.com/stephenwilliams/go-clitools/internal/defaults"
	"github.com/stephenwilliams/go-clitools/tools"
	"github.com/stephenwilliams/go-clitools/tools/kustomize"
)

// Sets one or more commonLabels in kustomization.yaml
func Label(label string, opts ...LabelOpt) *tools.CommandContext {
	options := &labelOptions{}

	for _, optFunc := range opts {
		optFunc(options)
	}

	var _args []string
	_args = append(_args, "edit", "set", "label")

	_args = append(_args, options.toArgs()...)
	_args = append(_args, label)

	return &tools.CommandContext{
		Args: _args,
		Path: tools.MustResolveTool(kustomize.KustomizeToolInfo, defaults.String(options.ToolVersion, kustomize.DefaultToolVersion), options.ToolProvider, kustomize.DefaultToolProvider, tools.DefaultToolProvider),
	}
}

type LabelOpt func(*labelOptions)

func SetToolProvider(p tools.ToolProvider) LabelOpt {
	return func(opts *labelOptions) {
		opts.ToolProvider = p
	}
}

func SetToolVersion(v string) LabelOpt {
	return func(opts *labelOptions) {
		opts.ToolVersion = v
	}
}

type labelOptions struct {
	ToolProvider tools.ToolProvider
	ToolVersion  string
}

func (o *labelOptions) toArgs() []string {
	var renderedArgs []string

	return renderedArgs
}
