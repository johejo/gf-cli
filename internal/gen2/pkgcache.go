package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// pkgCache caches parsed AST files to avoid re-parsing the same directory.
var pkgCache = make(map[string][]*ast.File)

func getModelParsedFiles(baseDir string) ([]*ast.File, error) {
	const key = "models|comments"
	if cached, ok := pkgCache[key]; ok {
		return cached, nil
	}
	pkgDir := filepath.Join(baseDir, "models")
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		return nil, fmt.Errorf("reading models dir %s: %w", pkgDir, err)
	}
	fset := token.NewFileSet()
	var files []*ast.File
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(pkgDir, e.Name()), nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", e.Name(), err)
		}
		files = append(files, f)
	}
	pkgCache[key] = files
	return files, nil
}

func getParsedFiles(baseDir string, pkgName string, mode parser.Mode) ([]*ast.File, error) {
	key := pkgName + "|" + fmt.Sprint(mode)
	if cached, ok := pkgCache[key]; ok {
		return cached, nil
	}
	pkgDir := filepath.Join(baseDir, "client", pkgName)
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		return nil, fmt.Errorf("reading package dir %s: %w", pkgDir, err)
	}
	fset := token.NewFileSet()
	var files []*ast.File
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(pkgDir, e.Name()), nil, mode)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", e.Name(), err)
		}
		files = append(files, f)
	}
	pkgCache[key] = files
	return files, nil
}
