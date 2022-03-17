package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"golang.org/x/tools/imports"
)

func ensureDirectoryExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to stat dir '%s': %w", dir, err)
	}

	return nil
}

func (g *toolGenerator) save(f *jen.File, relPath, dest string) error {
	g.generatedFiles = append(g.generatedFiles, dest)

	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return err
	}

	data, err := removeUnnecessaryLocalImportNames(dest, buf.Bytes())
	if err != nil {
		return err
	}

	data, err = imports.Process(dest, data, nil)
	if err != nil {
		return err
	}

	if !fileExists(dest) {
		fmt.Println("    saving new file:", relPath)
		if err := ioutil.WriteFile(dest, data, 0644); err != nil {
			return err
		}

		return nil
	}

	currentData, err := ioutil.ReadFile(dest)
	if err != nil {
		return fmt.Errorf("failed to read existing file for comparison '%s': %w", dest, err)
	}

	if string(data) == string(currentData) {
		return nil
	}

	fmt.Println("    saving changes to file:", relPath)

	if err := ioutil.WriteFile(dest, data, 0644); err != nil {
		return err
	}

	return nil
}

func descriptionComment(desc string) func(*jen.Statement) {
	return func(s *jen.Statement) {
		if desc == "" {
			return
		}
		lines := strings.Split(desc, "\n")
		for i, line := range lines {
			if i != len(lines)-1 {
				s.Comment(line).Line()
			} else {
				s.Comment(line)
			}
		}
	}
}

func fileExists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}

	return true
}

func removeUnnecessaryLocalImportNames(filename string, src []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse source code: %w", err)
	}

	changed := false
	rewriteFunc := func(x ast.Node) bool {
		s, ok := x.(*ast.ImportSpec)
		if !ok {
			return true
		}

		if s.Name == nil {
			return false
		}

		if lastImportPathPart(s.Path.Value) == s.Name.Name {
			changed = true
			s.Name = nil
		}

		return false
	}

	ast.Inspect(file, rewriteFunc)

	if !changed {
		return src, nil
	}

	buf := bytes.Buffer{}
	if err := format.Node(&buf, fset, file); err != nil {
		return nil, fmt.Errorf("failed to format source code: %w", err)
	}

	return buf.Bytes(), nil
}

func lastImportPathPart(p string) string {
	i := strings.LastIndexByte(p, '/')
	if i == -1 {
		return p[1 : len(p)-1]
	}

	return p[i+1 : len(p)-1]
}

func readNLines(path string, n int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	builder := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for i := 0; i < n && scanner.Scan(); i++ {
		builder.WriteString(scanner.Text())

		if i+1 < n {
			builder.WriteRune('\n')
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return builder.String(), nil
}

func isDirectoryEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	// Either not empty or error, suits both cases
	return false, err
}
