package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"unicode"
)

type ParamField struct {
	FieldName string // "TeamID"
	Type      string // "string", "int64", "bool", "[]string", "[]int64"
	IsPtr     bool
	IsBody    bool
	ModelType string // non-empty for body fields, e.g. "models.AddTeamRoleCommand"
	Doc       string // cleaned doc comment from source, e.g. "Search Query"
	Default   string // Go-literal default extracted from the doc comment, e.g. `1000` or `"View"`; empty if absent
}

// ParseParams finds the *Params struct by type name within the package directory
// and extracts its exported fields.
func ParseParams(baseDir string, pkgName string, paramsTypeName string) ([]*ParamField, error) {
	files, err := getParsedFiles(baseDir, pkgName, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		st := findStructType(f, paramsTypeName)
		if st == nil {
			continue
		}
		return extractParamFields(st)
	}
	return nil, fmt.Errorf("struct %s not found in package %s", paramsTypeName, pkgName)
}

func findStructType(f *ast.File, typeName string) *ast.StructType {
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
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			return st
		}
	}
	return nil
}

func extractParamFields(st *ast.StructType) ([]*ParamField, error) {
	var fields []*ParamField
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		name := field.Names[0].Name
		// Skip unexported fields
		if !unicode.IsUpper(rune(name[0])) {
			continue
		}
		// Skip infrastructure fields
		switch name {
		case "HTTPClient", "Context":
			continue
		}

		doc, def := cleanFieldDocAndDefault(field.Doc)
		pf := &ParamField{
			FieldName: name,
			Doc:       doc,
			Default:   def,
		}

		if name == "Body" {
			pf.IsBody = true
			pf.ModelType = resolveModelType(field.Type)
			pf.Type = "string"
		} else {
			pf.Type, pf.IsPtr = resolveFieldType(field.Type)
		}

		if pf.Type != "" {
			fields = append(fields, pf)
		}
	}
	return fields, nil
}

// cleanFieldDocAndDefault extracts a clean help string and the literal default
// value from a go-swagger field comment.
//
// Input format (/* block comment */):
//
//	/* FieldName.
//
//	   Description text here
//	   possibly multi-line
//
//	   Default: "value"
//	   Format: int64
//	*/
//
// We strip the first line (field name) and "Format:"/"Default:" lines from the
// returned text, and capture the raw token after "Default: " separately. The
// returned default literal is already in Go-source form for primitive types
// (ints bare, strings quoted, bools/floats bare), since go-swagger emits it
// that way. Empty string means no Default: line was present.
func cleanFieldDocAndDefault(doc *ast.CommentGroup) (text string, defaultLit string) {
	if doc == nil {
		return "", ""
	}
	src := doc.Text() // already strips /* */ and // prefixes
	lines := strings.Split(strings.TrimSpace(src), "\n")
	if len(lines) == 0 {
		return "", ""
	}

	var kept []string
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i == 0 {
			// Skip the "FieldName." header line
			continue
		}
		if strings.HasPrefix(line, "Format:") {
			// Type hint, not useful for CLI help
			continue
		}
		if rest, ok := strings.CutPrefix(line, "Default:"); ok {
			// Capture the literal; later occurrences win (matches go-swagger's
			// "last one wins" if multiple were ever emitted).
			defaultLit = strings.TrimSpace(rest)
			continue
		}
		if line != "" {
			kept = append(kept, line)
		}
	}
	return strings.Join(kept, " "), defaultLit
}

// resolveFieldType resolves an AST type expression to a Go type string and pointer status.
func resolveFieldType(expr ast.Expr) (typ string, isPtr bool) {
	switch t := expr.(type) {
	case *ast.Ident:
		return mapBasicType(t.Name), false
	case *ast.StarExpr:
		inner, _ := resolveFieldType(t.X)
		return inner, true
	case *ast.ArrayType:
		elem, _ := resolveFieldType(t.Elt)
		if elem != "" {
			return "[]" + elem, false
		}
		return "", false
	default:
		return "", false
	}
}

func mapBasicType(name string) string {
	switch name {
	case "string", "int64", "int", "bool", "float64":
		return name
	default:
		return ""
	}
}

// resolveModelType extracts the model type string from a body field type.
// e.g., *models.AddTeamRoleCommand -> "models.AddTeamRoleCommand"
// e.g., interface{} -> "interface{}"
func resolveModelType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return resolveModelType(t.X)
	case *ast.SelectorExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name + "." + t.Sel.Name
		}
	case *ast.ArrayType:
		elem := resolveModelType(t.Elt)
		if elem != "" {
			return "[]" + elem
		}
	case *ast.InterfaceType:
		return "interface{}"
	}
	return ""
}
