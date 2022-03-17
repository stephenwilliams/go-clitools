package tools

import (
	"os/exec"

	"github.com/Masterminds/semver/v3"
)

type PathToolProvider struct {
}

func (p PathToolProvider) GetPath(tool ToolInfo, version string) (string, error) {
	if version == "" {
		return exec.LookPath(tool.ExecutableName())
	}

	if isVersionConstraints(version) {
		constraints, err := semver.NewConstraint(version)
		if err != nil {
			return "", err
		}

		return p.GetPathWithConstraint(tool, constraints)
	}

	return exec.LookPath(tool.ExecutableName())
}

func (p PathToolProvider) GetPathWithVersion(tool ToolInfo, version *semver.Version) (string, error) {
	return exec.LookPath(tool.ExecutableName())
}

func (p PathToolProvider) GetPathWithConstraint(tool ToolInfo, constraints *semver.Constraints) (string, error) {
	path, err := exec.LookPath(tool.ExecutableName())
	if err != nil {
		return "", err
	}

	currentVersionRaw, err := tool.GetVersionWithPath(path)
	if err != nil {
		return "", ErrFailedToDetermineVersion
	}

	currentVersion, err := semver.NewVersion(currentVersionRaw)
	if err != nil {
		return "", ErrFailedToDetermineVersion
	}

	if !constraints.Check(currentVersion) {
		return "", nil
	}

	return path, nil
}
