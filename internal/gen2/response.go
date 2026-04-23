package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type ResponseField struct {
	HasPayload    bool
	ContainsFrame bool
}

// ParseResponse finds the response struct by type name within the package directory
// and checks if it has a Payload field. When a Payload is present its type tree is
// walked into the models package to detect whether models.Frame appears anywhere —
// responses containing Frame have a wire-format mismatch with Grafana's actual
// output, so callers use this flag to default --raw to true on those commands.
func ParseResponse(baseDir string, pkgName string, responseTypeName string) (*ResponseField, error) {
	files, err := getParsedFiles(baseDir, pkgName, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		st := findStructType(f, responseTypeName)
		if st == nil {
			continue
		}
		for _, field := range st.Fields.List {
			if len(field.Names) == 0 || field.Names[0].Name != "Payload" {
				continue
			}
			rf := &ResponseField{HasPayload: true}
			modelFiles, err := getModelParsedFiles(baseDir)
			if err != nil {
				return rf, nil
			}
			rf.ContainsFrame = typeExprContainsFrame(modelFiles, field.Type, map[string]bool{})
			return rf, nil
		}
		return &ResponseField{HasPayload: false}, nil
	}
	return nil, fmt.Errorf("response type %s not found in package %s", responseTypeName, pkgName)
}

// typeExprContainsFrame walks a type expression and recursively follows named
// types declared in the models package, returning true if models.Frame appears
// anywhere in the transitive type graph.
func typeExprContainsFrame(files []*ast.File, expr ast.Expr, visited map[string]bool) bool {
	expr = unwrapPointer(expr)
	switch t := expr.(type) {
	case *ast.Ident:
		return walkNamedModelType(files, t.Name, visited)
	case *ast.SelectorExpr:
		if pkgIdent, ok := t.X.(*ast.Ident); ok && pkgIdent.Name == "strfmt" {
			return false
		}
		return walkNamedModelType(files, t.Sel.Name, visited)
	case *ast.ArrayType:
		return typeExprContainsFrame(files, t.Elt, visited)
	case *ast.MapType:
		return typeExprContainsFrame(files, t.Key, visited) ||
			typeExprContainsFrame(files, t.Value, visited)
	case *ast.StarExpr:
		return typeExprContainsFrame(files, t.X, visited)
	}
	return false
}

func walkNamedModelType(files []*ast.File, name string, visited map[string]bool) bool {
	if name == "Frame" {
		return true
	}
	if _, ok := basicTypeToJSON(name); ok {
		return false
	}
	if visited[name] {
		return false
	}
	visited[name] = true
	for _, f := range files {
		ts := findTypeSpec(f, name)
		if ts == nil {
			continue
		}
		switch ut := ts.Type.(type) {
		case *ast.StructType:
			for _, field := range ut.Fields.List {
				if typeExprContainsFrame(files, field.Type, visited) {
					return true
				}
			}
			return false
		default:
			return typeExprContainsFrame(files, ut, visited)
		}
	}
	return false
}

func findTypeSpec(f *ast.File, typeName string) *ast.TypeSpec {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != typeName {
				continue
			}
			return ts
		}
	}
	return nil
}
