package gen

import (
	"errors"
	"fmt"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/stephenwilliams/go-clitools/_build/toolgen/tools"
)

type toolProviderGenerator struct {
	toolGen             *toolGenerator
	defaultToolProvider *tools.ToolProvider
	toolInfoVar         string

	f *File
}

func newToolProviderGenerator(toolGen *toolGenerator, toolInfoVar string) *toolProviderGenerator {
	return &toolProviderGenerator{
		toolGen:             toolGen,
		toolInfoVar:         toolInfoVar,
		defaultToolProvider: toolGen.tool.DefaultToolProvider,
	}
}

func (g *toolProviderGenerator) Generate() error {
	dest := filepath.Join(g.toolGen.toolPath, fmt.Sprintf("%s_tool_provider.go", g.toolGen.tool.Package))

	g.f = NewFilePath(g.toolGen.toolPath)
	g.f.HeaderComment(g.toolGen.header)

	g.f.Var().Id("DefaultToolVersion").String()

	if g.defaultToolProvider == nil {
		g.f.Var().Id("DefaultToolProvider").
			Qual(pkgTools, "ToolProvider").Op("=").
			Qual(pkgTools, "PathToolProvider").Values()
	} else {
		tp, err := g.generateToolProvider(g.defaultToolProvider)
		if err != nil {
			return err
		}

		g.f.Var().Id("DefaultToolProvider").
			Qual(pkgTools, "ToolProvider").Op("=").Add(tp)
	}

	g.generateEnsureToolProvider()

	return g.toolGen.save(g.f, g.toolGen.GetRelativePath(dest), dest)
}

func (g *toolProviderGenerator) generateEnsureToolProvider() {
	g.f.ImportName(pkgDefaults, "defaults")

	version := Qual(pkgDefaults, "String").Call(List(Id("v"), Id("DefaultToolVersion")))

	g.f.Func().Id("EnsureToolVersion").Params(Id("v").String()).Error().Block(
		If(List(Id("path"), Id("err")).Op(":=").Id("DefaultToolProvider").
			Dot("GetPath").Call(List(Id(g.toolInfoVar), version)).
			Op(";").Id("err").Op("!=").Nil()).Block(Return(Id("err"))).Else().If(Id("path").Op("==").Lit("")).Block(
			Return(Qual("fmt", "Errorf").Call(List(Lit("tool not found with version '%s'"), Id("v")))),
		),
		Line(),
		Return(Nil()),
	)
}

func (g *toolProviderGenerator) generateToolProvider(tp *tools.ToolProvider) (Code, error) {
	if tp.Chain != nil {
		return g.generateToolProviderChain(tp.Chain)
	} else if tp.DefinedToolPath != nil {
		return g.generateToolProviderDefinedPath(tp.DefinedToolPath), nil
	} else if tp.Downloader != nil {
		return g.generateToolProviderDownloader(tp.Downloader), nil
	} else if tp.Path != nil {
		return g.generateToolProviderPath(), nil
	} else {
		return nil, errors.New("no tool provider defined")
	}
}

func (g *toolProviderGenerator) generateToolProviderChain(tp *tools.ToolProviderChain) (Code, error) {
	providers := make([]Code, len(tp.Providers))
	for i, p := range tp.Providers {
		code, err := g.generateToolProvider(p)
		if err != nil {
			return nil, err
		}

		providers[i] = Add(Line(), code)
	}

	return Op("&").Qual(pkgTools, "ChainToolProvider").Values(Dict{
		Id("Providers"): Index().Qual(pkgTools, "ToolProvider").Values(providers...),
	}), nil
}

func (g *toolProviderGenerator) generateToolProviderDefinedPath(tp *tools.ToolProviderDefinedToolPath) Code {
	return Qual(pkgTools, "DefinedPathToolProvider").Call(Lit(tp.Path))
}

func (g *toolProviderGenerator) generateToolProviderDownloader(tp *tools.ToolProviderDownloader) Code {
	return Op("&").Qual(pkgTools, "DownloaderToolProvider").Values(Dict{
		Id("Downloader"): Id(tp.Downloader),
	})
}

func (g *toolProviderGenerator) generateToolProviderPath() Code {
	return Qual(pkgTools, "PathToolProvider").Values()
}
