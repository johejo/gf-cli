package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"unicode"
)

type ServiceEntry struct {
	FieldName string // "AccessControl"
	PkgName   string // "access_control"
}

// ParseFacade parses the GrafanaHTTPAPI struct from the facade file
// and returns the list of service entries.
func ParseFacade(baseDir string) ([]*ServiceEntry, error) {
	facadePath := filepath.Join(baseDir, "client", "grafana_http_api_client.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, facadePath, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parsing facade file: %w", err)
	}

	// Build import alias -> path map
	importMap := make(map[string]string) // alias/name -> import path
	for _, imp := range f.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		var name string
		if imp.Name != nil {
			name = imp.Name.Name
		} else {
			// last segment of path
			parts := strings.Split(path, "/")
			name = parts[len(parts)-1]
		}
		importMap[name] = path
	}

	// Find GrafanaHTTPAPI struct
	var entries []*ServiceEntry
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != "GrafanaHTTPAPI" {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, field := range st.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				name := field.Names[0].Name
				if !unicode.IsUpper(rune(name[0])) {
					continue
				}
				// Type should be a selector expr like access_control.ClientService
				sel, ok := field.Type.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				if sel.Sel.Name != "ClientService" {
					continue
				}
				pkgIdent, ok := sel.X.(*ast.Ident)
				if !ok {
					continue
				}
				pkgAlias := pkgIdent.Name
				if _, ok := importMap[pkgAlias]; !ok {
					continue
				}
				entries = append(entries, &ServiceEntry{
					FieldName: name,
					PkgName:   pkgAlias,
				})
			}
		}
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("no services found in GrafanaHTTPAPI struct")
	}
	return entries, nil
}
