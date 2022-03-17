package k3d

import "github.com/stephenwilliams/go-clitools/tools"

var GithubReleaseDownloader = &tools.GithubReleaseDownloader{
	AssetSelector:  "k3d-{{.OS}}-{{.Arch}}",
	Owner:          "rancher",
	Repository:     "k3d",
	TagPrefix:      "v",
	Tool:           "k3d",
	ToolExecutable: "k3d",
}
