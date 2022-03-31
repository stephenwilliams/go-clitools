package kubeconform

import "github.com/stephenwilliams/go-clitools/tools"

var GithubReleaseDownloader = &tools.GithubReleaseDownloader{
	ArchivePath:    "kubeconform",
	Archived:       true,
	AssetSelector:  "kubeconform-{{.OS}}-{{.Arch}}.tar.gz",
	Owner:          "yannh",
	Repository:     "kubeconform",
	TagPrefix:      "v",
	Tool:           "kubeconform",
	ToolExecutable: "kubeconform",
}
