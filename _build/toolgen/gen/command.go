package gen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/stephenwilliams/go-clitools/_build/toolgen/tools"
	"github.com/stephenwilliams/go-clitools/ptrhelpers"

	. "github.com/dave/jennifer/jen"
)

type commandGenerator struct {
	toolGen     *toolGenerator
	commandPath string
	cmd         *tools.Command
	toolInfoVar string

	f *File

	optionType string
	optsStruct string

	options []tools.Option
}

func newCommandGenerator(toolGen *toolGenerator, commandPath, toolInfoVar string, cmd *tools.Command) *commandGenerator {
	g := &commandGenerator{
		toolGen:     toolGen,
		commandPath: commandPath,
		toolInfoVar: toolInfoVar,
		cmd:         cmd,
		optionType:  fmt.Sprintf("%sOpt", cmd.ExportedName),
		optsStruct:  fmt.Sprintf("%sOptions", cmd.Name),
	}

	g.options = g.getOptions()

	return g
}

func (g *commandGenerator) Generate() error {
	g.f = NewFilePath(g.commandPath)
	g.f.HeaderComment(g.toolGen.header)

	if len(g.cmd.ArgumentSets) == 0 {
		// Generate the command func with no arguments
		g.generateCommandFunctionForArgSet(&tools.ArgumentSet{})
	} else {
		for _, argSet := range g.cmd.ArgumentSets {
			g.generateCommandFunctionForArgSet(&argSet)
		}
	}

	g.f.Type().Id(g.optionType).Func().Params(Op("*").Id(g.optsStruct))
	g.f.Line()
	if err := g.generateOptionsFuncs(); err != nil {
		return err
	}
	g.f.Line()
	g.f.Type().Id(g.optsStruct).StructFunc(func(structGrp *Group) {
		structGrp.Id("ToolProvider").Qual(pkgTools, "ToolProvider")
		structGrp.Id("ToolVersion").String()
		for _, opt := range g.options {
			structGrp.Id(opt.Name).Add(tools.GetGoType(opt.Type).MustGetOptionalTypeString())
		}
	})
	g.f.Line()

	g.f.Func().Parens(Id("o").Op("*").Id(g.optsStruct)).
		Id("toArgs").Params().Index().String().BlockFunc(g.generateOptionsToArgsFuncBody())

	dest := filepath.Join(g.commandPath, fmt.Sprintf("%s.go", g.cmd.Name))

	return g.toolGen.save(g.f, g.toolGen.GetRelativePath(dest), dest)
}

func (g *commandGenerator) generateCommandFunctionForArgSet(argSet *tools.ArgumentSet) {
	g.f.Do(descriptionComment(ptrhelpers.StringValue(g.cmd.Description)))
	g.f.Func().Id(g.cmd.ExportedName+argSet.ExportedNameSuffix).ParamsFunc(func(grp *Group) {
		for _, arg := range argSet.Args {
			grp.Id(arg.Name).Add(tools.GetGoType(arg.Type).MustGetTypeString())
		}

		grp.Id("opts").Op("...").Id(g.optionType)
	}).Op("*").Qual(pkgTools, "CommandContext").BlockFunc(g.commandFunctionBlock(argSet))
	g.f.Line()
}

func (g *commandGenerator) commandFunctionBlock(argSet *tools.ArgumentSet) func(grp *Group) {
	return func(grp *Group) {
		grp.Id("options").Op(":=").Op("&").Id(g.optsStruct).Values()
		grp.Line()
		grp.For(List(Id("_"), Id("optFunc"))).Op(":=").Range().Id("opts").Block(
			Id("optFunc").Call(Id("options")),
		)
		grp.Line()

		grp.Var().Id("_args").Id("[]string")
		if len(g.cmd.SubCommandPath) > 0 {
			grp.Id("_args").Op("=").AppendFunc(func(appendGrp *Group) {
				appendGrp.Id("_args")
				for _, part := range g.cmd.SubCommandPath {
					appendGrp.Lit(part)
				}
			})
		}
		grp.Line()

		if !ptrhelpers.BoolValue(g.cmd.ArgsFirst) {
			grp.Id("_args").Op("=").Append(Id("_args"), Id("options").Dot("toArgs").Call().Op("..."))
		}

		for _, arg := range argSet.Args {
			goType := tools.GetGoType(arg.Type)

			if ptrhelpers.BoolValue(arg.ExcludeDefaultValue) {
				grp.If(goType.MustGetDefaultValueCondition(arg.Name)).
					Block(goType.MustGetCommandArgsAppend(
						"_args", arg.Name, ptrhelpers.StringValue(arg.Format)))
			} else {
				grp.Add(goType.MustGetCommandArgsAppend(
					"_args", arg.Name, ptrhelpers.StringValue(arg.Format))).Line()
			}
		}

		if ptrhelpers.BoolValue(g.cmd.ArgsFirst) {
			grp.Id("_args").Op("=").Append(Id("_args"), Id("options").Dot("toArgs").Call().Op("..."))
		}

		grp.Line()

		g.f.ImportName(pkgDefaults, "defaults")

		grp.Return(Op("&").Qual(pkgTools, "CommandContext")).Values(Dict{
			Id("Path"): Qual(pkgTools, "MustResolveTool").Call(
				Qual(g.getPathForToolRoot(), g.toolInfoVar),
				Qual(pkgDefaults, "String").Call(List(Id("options").Dot("ToolVersion"), Qual(g.getPathForToolRoot(), "DefaultToolVersion"))),
				Id("options").Dot("ToolProvider"),
				Qual(g.getPathForToolRoot(), "DefaultToolProvider"),
				Qual(pkgTools, "DefaultToolProvider"),
			),
			Id("Args"): Id("_args"),
		})
	}
}

func (g *commandGenerator) generateOptionsFuncs() error {
	g.f.Func().Id("SetToolProvider").Params(Id("p").Qual(pkgTools, "ToolProvider")).Id(g.optionType).Block(
		Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
			Id("opts").Dot("ToolProvider").Op("=").Id("p"),
		)),
	)

	g.f.Line()

	g.f.Func().Id("SetToolVersion").Params(Id("v").String()).Id(g.optionType).Block(
		Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
			Id("opts").Dot("ToolVersion").Op("=").Id("v"),
		)),
	)

	for _, opt := range g.options {
		switch tools.GetGoType(opt.Type) {
		case tools.GoTypeBoolean:
			g.f.Do(descriptionComment(ptrhelpers.StringValue(opt.Description)))
			if ptrhelpers.BoolValue(opt.NoValue) {
				g.f.Func().Id(opt.ExportedName).Params().Id(g.optionType).Block(
					Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
						Id("opts").Dot(opt.Name).Op("=").Qual(pkgPTRHelpers, "Bool").Call(True()),
					)),
				)
			} else {
				g.f.Func().Id(opt.ExportedName).Params(Id("value").Bool()).Id(g.optionType).Block(
					Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
						Id("opts").Dot(opt.Name).Op("=").Id("&value"),
					)),
				)
			}
		case tools.GoTypeDuration:
			g.f.Do(descriptionComment(ptrhelpers.StringValue(opt.Description)))
			g.f.Func().Id(opt.ExportedName).Params(Id("value").Qual("time", "Duration")).
				Id(g.optionType).Block(Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
				Id("opts").Dot(opt.Name).Op("=").Id("&value"),
			)),
			)
		case tools.GoTypeInt:
			g.f.Do(descriptionComment(ptrhelpers.StringValue(opt.Description)))
			g.f.Func().Id(opt.ExportedName).Params(Id("value").Int()).
				Id(g.optionType).Block(Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
				Id("opts").Dot(opt.Name).Op("=").Id("&value"),
			)),
			)
		case tools.GoTypeString:
			g.f.Do(descriptionComment(ptrhelpers.StringValue(opt.Description)))
			g.f.Func().Id(opt.ExportedName).Params(Id("value").String()).Id(g.optionType).Block(
				Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
					Id("opts").Dot(opt.Name).Op("=").Id("&value"),
				)),
			)
		case tools.GoTypeStringSlice:
			g.f.Do(descriptionComment(ptrhelpers.StringValue(opt.Description)))
			g.f.Func().Id(opt.ExportedName).Params(Id("value").Id("...string")).Id(g.optionType).Block(
				Return(Func().Params(Id("opts").Op("*").Id(g.optsStruct)).Block(
					Id("opts").Dot(opt.Name).Op("=").Id("value"),
				)),
			)
		default:
			return fmt.Errorf("unknown type '%s'", opt.Type)
		}
	}

	return nil
}

func (g *commandGenerator) generateOptionsToArgsFuncBody() func(grp *Group) {
	return func(grp *Group) {
		grp.Var().Id("renderedArgs").Index().String()
		grp.Line()
		for _, opt := range g.options {
			grp.If(Id("o").Dot(opt.Name).Op("!=").Nil()).Block(Id("renderedArgs").Op("=").AppendFunc(func(appendGrp *Group) {
				appendGrp.Id("renderedArgs")
				for _, v := range opt.Format {
					if err := g.valueHandling(appendGrp, v, &opt); err != nil {
						panic(err)
					}
				}
			}))
			grp.Line()
		}
		grp.Return(Id("renderedArgs"))
	}
}

func (g *commandGenerator) valueHandling(grp *Group, format string, opt *tools.Option) error {
	if !strings.Contains(format, "%") {
		grp.Lit(format)
		return nil
	}

	g.f.ImportName("fmt", "fmt")

	vp, err := g.getValueProvider(opt)
	if err != nil {
		return err
	}

	grp.Id("fmt").Dot("Sprintf").Call(Lit(format), vp)
	return nil
}

func (g *commandGenerator) getValueProvider(opt *tools.Option) (Code, error) {
	switch tools.GetGoType(opt.Type) {
	case tools.GoTypeBoolean:
		return Qual(pkgPTRHelpers, "BoolValue").Call(Id("o").Dot(opt.Name)), nil
	case tools.GoTypeDuration:
		return Id("o").Dot(opt.Name), nil
	case tools.GoTypeInt:
		return Qual(pkgPTRHelpers, "IntValue").Call(Id("o").Dot(opt.Name)), nil
	case tools.GoTypeString:
		return Qual(pkgPTRHelpers, "StringValue").Call(Id("o").Dot(opt.Name)), nil
	case tools.GoTypeStringSlice:
		if opt.ValueJoin != nil {
			g.f.ImportName("strings", "strings")
			return Qual("strings", "Join").Call(Id("o").Dot(opt.Name), Lit(ptrhelpers.StringValue(opt.ValueJoin))), nil
		}

		return Id("o").Dot(opt.Name), nil
	default:
		return nil, fmt.Errorf("unknown type '%s'", opt.Type)
	}
}

func (g *commandGenerator) getPathForToolRoot() string {
	if g.commandPath == g.toolGen.toolPath {
		return g.commandPath
	}

	return g.toolGen.toolBaseImport
}

func (g *commandGenerator) getOptions() []tools.Option {
	if len(g.cmd.OptionSets) == 0 {
		return g.cmd.Options
	}

	length := len(g.cmd.Options)
	sets := make([][]tools.Option, len(g.cmd.OptionSets)+1)
	sets[0] = g.cmd.Options

	for _, setRef := range g.cmd.OptionSets {
		var set tools.OptionSet
		found := false
		for _, v := range g.toolGen.tool.OptionSets {
			if setRef.Name == v.Name {
				set = v
				found = true
				break
			}
		}

		if !found {
			panic(fmt.Errorf("command '%s' optionSet ref '%s' was not found in the tool optionSets", g.cmd.Name, setRef.Name))
		}

		length += len(set.Options)
		sets = append(sets, set.Options)
	}

	options := make([]tools.Option, length)

	i := 0
	for _, set := range sets {
		for _, opt := range set {
			options[i] = opt
			i++
		}
	}

	return options
}