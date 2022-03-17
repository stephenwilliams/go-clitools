package tools

import "github.com/Masterminds/semver/v3"

type DownloaderToolProvider struct {
	Downloader Downloader
}

func (p *DownloaderToolProvider) GetPath(tool ToolInfo, version string) (string, error) {
	if version == "" {
		return p.Downloader.Download(version)
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
		return p.Downloader.Download(version)
	}

	return p.GetPathWithVersion(tool, v)
}

func (p *DownloaderToolProvider) GetPathWithVersion(tool ToolInfo, version *semver.Version) (string, error) {
	return p.Downloader.DownloadWithVersion(version)
}

func (p *DownloaderToolProvider) GetPathWithConstraint(tool ToolInfo, constraints *semver.Constraints) (string, error) {
	return p.Downloader.DownloadWithConstraints(constraints)
}
