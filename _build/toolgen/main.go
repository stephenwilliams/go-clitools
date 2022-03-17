package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"

	"github.com/stephenwilliams/go-clitools/_build/toolgen/gen"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	fmt.Println("Starting Tool Generation")

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	baseDir, err := filepath.Abs(filepath.Join(wd, "..", ".."))
	if err != nil {
		panic(err)
	}

	baseImportPath, err := getModuleImport(baseDir)
	if err != nil {
		panic(err)
	}

	schema := "file://" + filepath.Join(baseDir, "_build", "tool.schema.json")

	var files []string
	err = filepath.Walk(filepath.Join(baseDir, "_build", "specifications"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	schemaLoader := gojsonschema.NewReferenceLoader(schema)
	allValid := true
	for _, f := range files {
		doc := gojsonschema.NewReferenceLoader("file://" + f)

		result, err := gojsonschema.Validate(schemaLoader, doc)
		if err != nil {
			panic(err)
		}

		if !result.Valid() {
			fmt.Printf("The document '%s' is not valid. see errors :\n", f)
			allValid = false
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}
	}
	if !allValid {
		os.Exit(1)
	}

	for _, f := range files {
		g, err := gen.NewToolGenerator(baseDir, baseImportPath, f)
		if err != nil {
			panic(err)
		}

		err = g.Generate()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Complete")
}

func getModuleImport(baseDir string) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(baseDir, "go.mod"))
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod file: %w", err)
	}

	importPath := modfile.ModulePath(data)

	if importPath == "" {
		return "", errors.New("failed to determine module path")
	}

	return importPath, nil
}
