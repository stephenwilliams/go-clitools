package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stephenwilliams/go-clitools/_build/toolgen/tools"
)

const (
	headerFirstLine = "// Generated by toolgen. DO NOT EDIT."
)

type ToolGenerator interface {
	Generate() error
}

type toolGenerator struct {
	header string

	basePath       string
	toolPath       string
	toolBaseImport string
	tool           *tools.Tool

	generatedFiles []string
}

func NewToolGenerator(basePath, baseImport, specPath string) (ToolGenerator, error) {
	data, err := ioutil.ReadFile(specPath)
	if err != nil {
		return nil, err
	}

	tool := &tools.Tool{}
	if err := json.Unmarshal(data, tool); err != nil {
		return nil, err
	}

	absSpecPath, err := filepath.Abs(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of specPath: %w", err)
	}

	absBasePath, err := filepath.Abs(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of basePath: %w", err)
	}

	absBasePath += "/"

	relPath := strings.Replace(absSpecPath, absBasePath, "", 1)
	header := fmt.Sprintf(`%s
// Generated from tool specification: 
//   %s`, headerFirstLine, relPath)

	return &toolGenerator{
		header:         header,
		basePath:       absBasePath,
		toolPath:       filepath.Join(basePath, "tools", tool.Package),
		toolBaseImport: path.Join(baseImport, "tools", tool.Package),
		tool:           tool,
	}, nil
}

func (g *toolGenerator) Generate() error {
	fmt.Println("  running generator for tool:", g.tool.Name)
	if err := ensureDirectoryExists(g.toolPath); err != nil {
		return err
	}

	versionProviderGen := newToolInfoGenerator(g)
	toolInfoVar, err := versionProviderGen.Generate()
	if err != nil {
		return err
	}

	toolProviderGen := newToolProviderGenerator(g, toolInfoVar)
	if err := toolProviderGen.Generate(); err != nil {
		return err
	}

	downloadersGen := newDownloadersGenerator(g)
	if err := downloadersGen.Generate(); err != nil {
		return err
	}

	for _, grp := range g.tool.Groups {
		gen := newGroupGenerator(g, &grp, g.toolPath, toolInfoVar)

		if err := gen.Generate(); err != nil {
			return err
		}
	}

	cleanFunc := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if empty, err := isDirectoryEmpty(path); err != nil {
				return err
			} else if empty {
				fmt.Println("    deleted empty dir:", g.GetRelativePath(path))
				if err := os.Remove(path); err != nil {
					return err
				}
			}

			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		fileHeader, err := readNLines(path, 1)
		if err != nil {
			return err
		}

		if fileHeader == headerFirstLine {
			for _, generatedFile := range g.generatedFiles {
				if generatedFile == path {
					return nil
				}
			}

			// file was not found in generated files, so delete it
			fmt.Println("    deleted file:", g.GetRelativePath(path))
			if err := os.Remove(path); err != nil {
				return err
			}

			dir := filepath.Dir(path)
			if empty, err := isDirectoryEmpty(dir); err != nil {
				return err
			} else if empty {
				fmt.Println("    deleted empty dir:", g.GetRelativePath(dir))
				if err := os.Remove(dir); err != nil {
					return err
				}
			}
		}

		return nil
	}

	if err := filepath.Walk(filepath.Join(g.toolPath), cleanFunc); err != nil {
		return err
	}

	return nil
}

func (g *toolGenerator) GeneratePackageDocs(dir, description string) error {
	dest := filepath.Join(dir, "pkg_doc.go")
	f := jen.NewFilePath(dir)
	f.HeaderComment(g.header)
	f.PackageComment(description)
	return g.save(f, g.GetRelativePath(dest), dest)
}

func (g *toolGenerator) GetRelativePath(p string) string {
	return strings.Replace(p, g.basePath, "", 1)
}