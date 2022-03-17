package tools

import (
	"errors"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/stephenwilliams/go-clitools/internal/iohelpers"
)

var DefaultToolProvider ToolProvider = PathToolProvider{}

var (
	ErrFailedToDetermineVersion = errors.New("failed to determine error version")
)

type ToolProvider interface {
	GetPath(tool ToolInfo, version string) (string, error)
	GetPathWithVersion(tool ToolInfo, version *semver.Version) (string, error)
	GetPathWithConstraint(tool ToolInfo, constraints *semver.Constraints) (string, error)
}

func ResolveTool(tool ToolInfo, version string, providers ...ToolProvider) (string, error) {
	for _, provider := range providers {
		if provider != nil {
			return provider.GetPath(tool, version)
		}
	}

	return DefaultToolProvider.GetPath(tool, version)
}

func MustResolveTool(tool ToolInfo, version string, providers ...ToolProvider) string {
	if !iohelpers.DirExists(CacheDirectory) {
		if err := os.MkdirAll(CacheDirectory, 0700); err != nil {
			panic(fmt.Errorf("failed to create cache directory: %w", err))
		}
	}

	path, err := ResolveTool(tool, version, providers...)
	if err != nil {
		panic(fmt.Errorf("unable to resolve a tool with the given providers: %w", err))
	}

	return path
}
