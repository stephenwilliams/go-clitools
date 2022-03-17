package tools

import "github.com/Masterminds/semver/v3"

type DefinedPathToolProvider string

func (p DefinedPathToolProvider) GetPath(tool ToolInfo, version string) (string, error) {
	if version == "" {
		return string(p), nil
	} else if isVersionConstraints(version) {
		constraints, err := semver.NewConstraint(version)
		if err != nil {
			return "", err
		}

		return p.GetPathWithConstraint(tool, constraints)
	}

	return string(p), nil
}

func (p DefinedPathToolProvider) GetPathWithConstraint(tool ToolInfo, constraints *semver.Constraints) (string, error) {
	path := string(p)

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

func (p DefinedPathToolProvider) GetPathWithVersion(tool ToolInfo, version *semver.Version) (string, error) {
	return string(p), nil
}
