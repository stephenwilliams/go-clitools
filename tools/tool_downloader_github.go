package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/stephenwilliams/go-clitools/internal/oshelpers"

	"github.com/Masterminds/semver/v3"
	"github.com/stephenwilliams/go-clitools/internal/archives"
	"github.com/stephenwilliams/go-clitools/internal/github"
	"github.com/stephenwilliams/go-clitools/internal/hashing"
	"github.com/stephenwilliams/go-clitools/internal/templatehelpers"
)

type GithubReleaseDownloader struct {
	Tool           string
	ToolExecutable string
	Owner          string
	Repository     string
	TagPrefix      string

	AssetSelector       string
	AssetSelectorRegexp bool
	Archived            bool
	ArchivePath         string

	MultipleReleases          bool
	MultipleReleasesTagPrefix string
}

var _ Downloader = &GithubReleaseDownloader{}

func (d *GithubReleaseDownloader) Download(version string) (string, error) {
	if err := cleanCacheDirectory(); err != nil {
		return "", err
	}

	release, err := d.getRelease(version)
	if err != nil {
		return "", err
	}

	return d.downloadRelease(release)
}

func (d *GithubReleaseDownloader) DownloadWithConstraints(constraints *semver.Constraints) (string, error) {
	if err := cleanCacheDirectory(); err != nil {
		return "", err
	}

	release, err := d.findRelease(constraints)
	if err != nil {
		return "", err
	}

	return d.downloadRelease(release)
}

func (d *GithubReleaseDownloader) DownloadWithVersion(version *semver.Version) (string, error) {
	return d.Download(version.String())
}

func (d *GithubReleaseDownloader) downloadRelease(release *github.Release) (string, error) {
	dest := pathForTool(d.ToolExecutable, d.formatVersion(release))

	if found, err := fileExists(dest); err != nil {
		return "", err
	} else if found {
		return dest, nil
	}

	tmpDir, err := ioutil.TempDir(os.TempDir(), "go-clitools_github_release_downloader_")
	if err != nil {
		return "", err
	}

	assetPath, err := d.downloadAsset(tmpDir, release)
	if err != nil {
		return "", err
	}

	if err := ensureDirectory(BinDirectory); err != nil {
		return "", err
	}

	if !d.Archived {
		if err := os.Rename(assetPath, dest); err != nil {
			return "", err
		}

		stat, err := os.Stat(dest)
		if err != nil {
			return "", fmt.Errorf("failed to stat file: %w", err)
		}

		if !oshelpers.IsExecAny(stat.Mode()) {
			if err := os.Chmod(dest, 0700); err != nil {
				return "", fmt.Errorf("failed to make file executable: %w", err)
			}
		}

		return dest, nil
	}

	extracted, err := d.extractTool(tmpDir, assetPath, release)
	if err != nil {
		return "", err
	}

	if err := os.Rename(extracted, dest); err != nil {
		return "", err
	}

	stat, err := os.Stat(dest)
	if err != nil {
		return "", fmt.Errorf("failed to stat file: %w", err)
	}

	if !oshelpers.IsExecAny(stat.Mode()) {
		if err := os.Chmod(dest, 0700); err != nil {
			return "", fmt.Errorf("failed to make file executable: %w", err)
		}
	}

	return dest, nil
}

func (d *GithubReleaseDownloader) findRelease(constraints *semver.Constraints) (*github.Release, error) {
	release := &github.Release{}
	key := fmt.Sprintf("%s_%s", d.Owner, d.Repository)
	versionKey := fmt.Sprintf("constraint_%s", hashing.SHA1([]byte(constraints.String())))
	if ok, err := getCache("github-release", key, versionKey, release); err != nil {
		return nil, err
	} else if ok {
		return release, nil
	}

	renderedTagPrefix := d.MultipleReleasesTagPrefix + d.TagPrefix

	selector := func(releases []*github.Release) (*github.Release, error) {
		for _, r := range releases {
			if r.Draft || r.Prerelease {
				continue
			}

			if d.MultipleReleases && !strings.HasPrefix(r.TagName, renderedTagPrefix) {
				continue
			}

			v, err := semver.NewVersion(d.formatVersion(r))
			if err != nil {
				continue
			}

			if constraints.Check(v) {
				return r, nil
			}
		}

		return nil, nil
	}

	release, err := github.FindRelease(d.Owner, d.Repository, selector)
	if err != nil {
		return nil, err
	} else if release == nil {
		return nil, errors.New("github release not found")
	}

	if err := setCache("github-release", key, versionKey, release); err != nil {
		return nil, err
	}

	return release, nil
}

// findLatestInMultipleRelease finds the latest release when a repository releases multiple applications out of the same
// repository with different tag prefixes
func (d *GithubReleaseDownloader) findLatestInMultipleRelease() (*github.Release, error) {
	release := &github.Release{}
	key := fmt.Sprintf("%s_%s", d.Owner, d.Repository)
	if ok, err := getCache("github-release", key, "latest_multiple_release", release); err != nil {
		return nil, err
	} else if ok {
		return release, nil
	}

	renderedTagPrefix := d.MultipleReleasesTagPrefix + d.TagPrefix

	selector := func(releases []*github.Release) (*github.Release, error) {
		for _, r := range releases {
			if r.Draft || r.Prerelease {
				continue
			}

			if strings.HasPrefix(r.TagName, renderedTagPrefix) {
				return r, nil
			}
		}

		return nil, nil
	}

	release, err := github.FindRelease(d.Owner, d.Repository, selector)
	if err != nil {
		return nil, err
	} else if release == nil {
		return nil, errors.New("github release not found")
	}

	if err := setCache("github-release", key, "latest_multiple_release", release); err != nil {
		return nil, err
	}

	return release, nil
}

func (d *GithubReleaseDownloader) getRelease(version string) (*github.Release, error) {
	if version == "" {
		version = "latest"
	}

	release := &github.Release{}
	if ok, err := getCache("github-release", fmt.Sprintf("%s_%s", d.Owner, d.Repository), version, release); err != nil {
		return nil, err
	} else if ok {
		return release, nil
	}

	if version == "latest" {
		if d.MultipleReleases {
			var err error
			release, err = d.findLatestInMultipleRelease()
			if err != nil {
				return nil, err
			}
		} else {
			var err error
			release, err = github.GetLatestRelease(d.Owner, d.Repository)
			if err != nil {
				return nil, err
			}
		}
	} else {
		v := version
		if d.TagPrefix != "" && !strings.HasPrefix(v, d.TagPrefix) {
			v = d.TagPrefix + v
		}

		if d.MultipleReleases {
			v = d.MultipleReleasesTagPrefix + v
		}

		var err error
		release, err = github.GetReleaseByTag(d.Owner, d.Repository, v)
		if err != nil {
			return nil, err
		}
	}

	if release == nil {
		return nil, errors.New("github release was nil")
	}

	if err := setCache("github-release", fmt.Sprintf("%s_%s", d.Owner, d.Repository), version, release); err != nil {
		return nil, err
	}

	return release, nil
}

func (d *GithubReleaseDownloader) downloadAsset(dir string, release *github.Release) (string, error) {
	var assetName string
	if strings.Contains(d.AssetSelector, "{") {
		t, err := template.New("").Funcs(templatehelpers.Funcs).Parse(d.AssetSelector)
		if err != nil {
			return "", nil
		}

		sb := strings.Builder{}
		if err := t.Execute(&sb, getDownloaderTemplateData(d.formatVersion(release))); err != nil {
			return "", err
		}

		assetName = strings.TrimSpace(sb.String())
	} else {
		assetName = d.AssetSelector
	}

	var assetSelectorRegexp *regexp.Regexp
	if d.AssetSelectorRegexp {
		var err error
		assetSelectorRegexp, err = regexp.Compile(assetName)
		if err != nil {
			return "", fmt.Errorf("failed to compile asset selector regexp: %w", err)
		}
	}

	var asset *github.Asset
	for _, asset = range release.Assets {
		if assetSelectorRegexp != nil {
			if assetSelectorRegexp.MatchString(asset.Name) {
				break
			}
		} else {
			if assetName == asset.Name {
				break
			}
		}
	}

	if asset == nil {
		return "", fmt.Errorf("unable to find asset '%s'", assetName)
	}

	dest := filepath.Join(dir, asset.Name)

	if err := downloadFile(github.HTTP, asset.BrowserDownloadURL, dest); err != nil {
		return "", err
	}

	return dest, nil
}

func (d *GithubReleaseDownloader) extractTool(dir, archive string, release *github.Release) (string, error) {
	unarchivedDir := filepath.Join(dir, "_unarchived")

	if err := archives.Unarchive(archive, unarchivedDir); err != nil {
		return "", err
	}

	var archivePath string
	if strings.Contains(d.ArchivePath, "{") {
		t, err := template.New("").Funcs(templatehelpers.Funcs).Parse(d.ArchivePath)
		if err != nil {
			return "", nil
		}

		sb := strings.Builder{}
		if err := t.Execute(&sb, getDownloaderTemplateData(d.formatVersion(release))); err != nil {
			return "", err
		}

		archivePath = strings.TrimSpace(sb.String())
	} else {
		archivePath = d.ArchivePath
	}

	path := filepath.Join(unarchivedDir, archivePath)

	if ok, err := fileExists(path); err != nil {
		return "", err
	} else if !ok {
		return "", fmt.Errorf("file does not exist '%s'", path)
	}

	return path, nil
}

func (d *GithubReleaseDownloader) formatVersion(r *github.Release) string {
	if d.MultipleReleases {
		return strings.TrimPrefix(r.TagName, d.MultipleReleasesTagPrefix)
	}

	return r.TagName
}
