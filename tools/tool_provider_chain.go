package tools

import (
	"errors"
	"github.com/hashicorp/go-multierror"

	"github.com/Masterminds/semver/v3"
)

type ChainToolProvider struct {
	Providers []ToolProvider
}

func (p *ChainToolProvider) GetPath(tool ToolInfo, version string) (string, error) {
	if version == "" {
		return p.getPath(tool, version)
	}

	if isVersionConstraints(version) {
		constraints, err := semver.NewConstraint(version)
		if err != nil {
			return "", err
		}

		return p.GetPathWithConstraint(tool, constraints)
	}

	v, err := semver.NewVersion(version)
	if err != nil || v == nil {
		return p.getPath(tool, version)
	}

	return p.GetPathWithVersion(tool, v)
}

func (p *ChainToolProvider) getPath(tool ToolInfo, version string) (string, error) {
	var resultErr error

	for _, tp := range p.Providers {
		if path, err := tp.GetPath(tool, version); err != nil && !errors.Is(err, ErrFailedToDetermineVersion) {
			err = multierror.Append(resultErr, err)
		} else if path != "" {
			return path, nil
		}
	}

	return "", resultErr
}

func (p *ChainToolProvider) GetPathWithVersion(tool ToolInfo, version *semver.Version) (string, error) {
	var resultErr error

	for _, tp := range p.Providers {
		if path, err := tp.GetPathWithVersion(tool, version); err != nil && !errors.Is(err, ErrFailedToDetermineVersion) {
			err = multierror.Append(resultErr, err)
		} else if path != "" {
			return path, nil
		}
	}

	return "", resultErr
}

func (p *ChainToolProvider) GetPathWithConstraint(tool ToolInfo, constraints *semver.Constraints) (string, error) {
	var resultErr error

	for _, tp := range p.Providers {
		if path, err := tp.GetPathWithConstraint(tool, constraints); err != nil && !errors.Is(err, ErrFailedToDetermineVersion) {
			err = multierror.Append(resultErr, err)
		} else if path != "" {
			return path, nil
		}
	}

	return "", resultErr
}

var _ = &ChainToolProvider{
	Providers: []ToolProvider{},
}
