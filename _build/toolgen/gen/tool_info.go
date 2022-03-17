package gen

import (
	"fmt"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/stephenwilliams/go-clitools/_build/toolgen/tools"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"
)

type toolInfoGenerator struct {
	toolGen *toolGenerator

	// Version Info
	args          []string
	selector      *tools.VersionSelector
	captureStderr bool

	f *File
}

func newToolInfoGenerator(toolGen *toolGenerator) *toolInfoGenerator {
	return &toolInfoGenerator{
		toolGen: toolGen,

		args:          toolGen.tool.VersionProvider.CommandArgs,
		selector:      &toolGen.tool.VersionProvider.Selector,
		captureStderr: toolGen.tool.VersionProvider.CaptureStderr,
	}
}

func (g *toolInfoGenerator) Generate() (string, error) {
	dest := filepath.Join(g.toolGen.toolPath, fmt.Sprintf("%s_info.go", g.toolGen.tool.Package))

	g.f = NewFilePath(g.toolGen.toolPath)
	g.f.HeaderComment(g.toolGen.header)

	typeName := fmt.Sprintf("%sToolInfo", g.toolGen.tool.Package)
	exportedVarName := fmt.Sprintf("%sToolInfo", g.toolGen.tool.ExportedName)

	g.f.Var().Id(exportedVarName).Op("=").Id(typeName).Values()
	g.f.Line()
	g.f.Type().Id(typeName).Struct()

	g.f.Line()

	g.f.Func().Params(Id(typeName)).Id("Name").Params().String().
		Block(Return(Lit(g.toolGen.tool.Name)))

	g.f.Line()

	g.f.Func().Params(Id(typeName)).Id("ExecutableName").Params().String().
		Block(Return(Lit(g.toolGen.tool.ExecutableName)))

	g.f.Line()

	g.f.Comment(`// GetVersion returns the version for the tool found with the provided tool provider.
// If tp is nil, the default tool provider is used.`)
	g.f.Func().Params(Id("i").Id(typeName)).Id("GetVersion").Params(Id("tp").Qual(pkgTools, "ToolProvider")).
		Parens(List(String(), Error())).BlockFunc(g.generateGetVersionFuncBody(exportedVarName))

	g.f.Line()

	g.f.Comment(`// GetVersionWithPath returns the version for the tool found with the provided path.
// to the tool`)
	g.f.Func().Params(Id(typeName)).Id("GetVersionWithPath").Params(Id("path").String()).
		Parens(List(String(), Error())).BlockFunc(g.generateGetVersionWithPathFuncBody(exportedVarName))

	g.f.Line()

	return exportedVarName, g.toolGen.save(g.f, g.toolGen.GetRelativePath(dest), dest)
}

func (g *toolInfoGenerator) generateGetVersionFuncBody(toolInfoVar string) func(*Group) {
	return func(grp *Group) {
		grp.If(Id("tp").Op("!=").Nil()).Block(Comment("// do nothing")).Else().
			If(Id("DefaultToolProvider").Op("!=").Nil()).
			Block(Id("tp").Op("=").Id("DefaultToolProvider")).Else().
			If(Qual(pkgTools, "DefaultToolProvider").Op("!=").Nil()).
			Block(Id("tp").Op("=").Qual(pkgTools, "DefaultToolProvider")).
			Else().Block(Return(List(Lit(""), Qual("errors", "New").
			Call(Lit("no tool provider provided and defaults are nil")))))

		grp.Line()

		grp.List(Id("path"), Err()).Op(":=").Qual(pkgTools, "ResolveTool").
			Call(List(Id(toolInfoVar), Id("DefaultToolVersion"), Id("tp")))
		grp.If(Err().Op("!=").Nil()).Block(Return(Lit(""), Err()))

		grp.Line()

		grp.Return(Id("i").Dot("GetVersionWithPath").Call(Id("path")))
	}
}

func (g *toolInfoGenerator) generateGetVersionWithPathFuncBody(toolInfoVar string) func(*Group) {
	return func(grp *Group) {
		grp.Id("cmd").Op(":=").Op("&").Qual(pkgTools, "CommandContext").Values(Dict{
			Id("Path"): Id("path"),
			Id("Args"): Index().String().ValuesFunc(func(values *Group) {
				for _, a := range g.args {
					values.Lit(a)
				}
			}),
		})

		grp.Line()

		var stdErr Code
		if g.captureStderr {
			stdErr = Op("&").Id("b")
		} else {
			stdErr = Qual("os", "Stderr")
		}

		grp.Var().Id("b").Qual("strings", "Builder")
		grp.List(Id("_"), Id("_"), Id("err")).Op(":=").
			Id("cmd").Dot("Exec").Call(List(Op("&").Id("b"), stdErr))
		grp.If(Err().Op("!=").Nil()).Block(Return(Lit(""), Err()))

		grp.Line()

		grp.Id("v").Op(":=").Id("b").Dot("String").Call()

		grp.Line()

		if g.selector.SplitLine != nil {
			grp.Id("v").Op("=").Qual(pkgVersionProviders, "SplitLine").
				Call(List(Id("v"), Lit(g.selector.SplitLine.Index)))

			grp.Line()
		}

		if g.selector.SplitString != nil {
			grp.Id("v").Op("=").Qual(pkgVersionProviders, "SplitString").
				Call(List(Id("v"), Lit(g.selector.SplitString.Separator), Lit(g.selector.SplitString.Index)))

			grp.Line()
		}

		for _, sr := range g.selector.StringReplace {
			n := ptrhelpers.IntValueDefault(sr.NumberOfReplacements, -1)

			grp.Id("v").Op("=").Qual("strings", "Replace").Call(List(
				Id("v"),
				Lit(sr.Old),
				Lit(sr.New),
				Lit(n),
			))

			grp.Line()
		}

		if g.selector.TrimPrefix != nil {
			grp.Id("v").Op("=").Qual("strings", "TrimPrefix").Call(List(
				Id("v"),
				Lit(ptrhelpers.StringValue(g.selector.TrimPrefix)),
			))

			grp.Line()
		}

		if g.selector.TrimSuffix != nil {
			grp.Id("v").Op("=").Qual("strings", "TrimSuffix").Call(List(
				Id("v"),
				Lit(ptrhelpers.StringValue(g.selector.TrimSuffix)),
			))

			grp.Line()
		}

		if er := g.selector.EqualsReplace; er != nil {
			grp.Id("v").Op("=").Qual(pkgVersionProviders, "EqualsReplace").
				Call(List(Id("v"), Lit(er.Old), Lit(er.New)))

			grp.Line()
		}

		grp.Line()

		grp.Id("v").Op("=").Qual("strings", "TrimSpace").Call(Id("v"))

		grp.Line()

		grp.Return(List(Id("v"), Nil()))
	}
}
