package gen

import (
	"path/filepath"

	"github.com/stephenwilliams/go-clitools/ptrhelpers"

	"github.com/stephenwilliams/go-clitools/_build/toolgen/tools"
)

type groupGenerator struct {
	toolGen     *toolGenerator
	group       *tools.CommandGroup
	toolInfoVar string

	groupPath   string
	description string
}

func newGroupGenerator(toolGen *toolGenerator, group *tools.CommandGroup, path, toolInfoVar string) *groupGenerator {
	var groupPath string
	var description string
	if group.IsRoot() {
		groupPath = toolGen.toolPath
		description = toolGen.tool.Description
	} else {
		groupPath = filepath.Join(path, group.Package)
		description = group.Description
	}

	return &groupGenerator{
		toolGen:     toolGen,
		group:       group,
		groupPath:   groupPath,
		description: description,
		toolInfoVar: toolInfoVar,
	}
}

func (g *groupGenerator) Generate() error {
	if err := ensureDirectoryExists(g.groupPath); err != nil {
		return err
	}

	if err := g.toolGen.GeneratePackageDocs(g.groupPath, g.description); err != nil {
		return err
	}

	if len(g.group.Commands) == 1 {
		gen := newCommandGenerator(g.toolGen, g.groupPath, g.toolInfoVar, &g.group.Commands[0])

		if err := gen.Generate(); err != nil {
			return err
		}
	} else {
		for _, cmd := range g.group.Commands {
			commandPath := filepath.Join(g.groupPath, ptrhelpers.StringValueDefault(cmd.Package, cmd.Name))

			if err := ensureDirectoryExists(commandPath); err != nil {
				return err
			}

			gen := newCommandGenerator(g.toolGen, commandPath, g.toolInfoVar, &cmd)

			if err := gen.Generate(); err != nil {
				return err
			}
		}
	}

	for _, subGroup := range g.group.Groups {
		gen := newGroupGenerator(g.toolGen, &subGroup, g.groupPath, g.toolInfoVar)

		if err := gen.Generate(); err != nil {
			return err
		}
	}

	return nil
}
