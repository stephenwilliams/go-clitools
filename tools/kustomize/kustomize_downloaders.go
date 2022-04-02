package kustomize

import "github.com/stephenwilliams/go-clitools/tools"

var GithubReleaseDownloader = &tools.GithubReleaseDownloader{
	ArchivePath:               "kustomize",
	Archived:                  true,
	AssetSelector:             "kustomize_v{{trimPrefix \"v\" .Version}}_{{.OS}}_{{.Arch}}.tar.gz",
	MultipleReleases:          true,
	MultipleReleasesTagPrefix: "kustomize/",
	Owner:                     "kubernetes-sigs",
	Repository:                "kustomize",
	TagPrefix:                 "v",
	Tool:                      "kustomize",
	ToolExecutable:            "kustomize",
}
