package svu

import "github.com/stephenwilliams/go-clitools/tools"

var GithubReleaseDownloader = &tools.GithubReleaseDownloader{
	ArchivePath:         "svu",
	Archived:            true,
	AssetSelector:       "svu_{{trimPrefix \"v\" .Version}}_{{.OS}}_({{.Arch}}|all).tar.gz",
	AssetSelectorRegexp: true,
	Owner:               "caarlos0",
	Repository:          "svu",
	TagPrefix:           "v",
	Tool:                "Semantic Version Util",
	ToolExecutable:      "svu",
}
