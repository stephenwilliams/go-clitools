package gen

import (
	"fmt"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/stephenwilliams/go-clitools/_build/toolgen/tools"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
)

type downloadersGenerator struct {
	toolGen     *toolGenerator
	downloaders *tools.Downloaders

	f *File
}

func newDownloadersGenerator(toolGen *toolGenerator) *downloadersGenerator {
	return &downloadersGenerator{
		toolGen:     toolGen,
		downloaders: toolGen.tool.Downloaders,
	}
}

func (g *downloadersGenerator) Generate() error {
	if g.downloaders == nil || g.downloaders.GithubRelease == nil {
		// Nothing to do
		return nil
	}

	dest := filepath.Join(g.toolGen.toolPath, fmt.Sprintf("%s_downloaders.go", g.toolGen.tool.Package))

	g.f = NewFilePath(g.toolGen.toolPath)

	if g.downloaders.GithubRelease != nil {
		g.f.Var().Id("GithubReleaseDownloader").Op("=").
			Op("&").Qual(pkgTools, "GithubReleaseDownloader").
			Values(DictFunc(g.githubReleaseStructValues()))
	}

	return g.toolGen.save(g.f, g.toolGen.GetRelativePath(dest), dest)
}

func (g *downloadersGenerator) githubReleaseStructValues() func(dict Dict) {
	return func(d Dict) {
		release := g.downloaders.GithubRelease

		d[Id("Tool")] = Lit(g.toolGen.tool.Name)
		d[Id("ToolExecutable")] = Lit(g.toolGen.tool.ExecutableName)
		d[Id("Owner")] = Lit(release.Owner)
		d[Id("Repository")] = Lit(release.Repository)
		d[Id("AssetSelector")] = Lit(release.AssetSelector)

		if release.TagPrefix != nil {
			d[Id("TagPrefix")] = Lit(ptrhelpers.StringValue(release.TagPrefix))
		}

		if ptrhelpers.BoolValue(release.Archived) {
			d[Id("Archived")] = True()
			d[Id("ArchivePath")] = Lit(ptrhelpers.StringValue(release.ArchivePath))
		}
	}
}
