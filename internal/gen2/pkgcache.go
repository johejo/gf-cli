package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

// pkgCache caches parsed AST packages to avoid re-parsing the same directory.
var pkgCache = make(map[string]map[string]*ast.Package)

func getParsedPkg(baseDir string, pkgName string, mode parser.Mode) (map[string]*ast.Package, error) {
	key := pkgName + "|" + fmt.Sprint(mode)
	if cached, ok := pkgCache[key]; ok {
		return cached, nil
	}
	pkgDir := filepath.Join(baseDir, "client", pkgName)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, pkgDir, nil, mode)
	if err != nil {
		return nil, fmt.Errorf("parsing package dir %s: %w", pkgDir, err)
	}
	pkgCache[key] = pkgs
	return pkgs, nil
}
