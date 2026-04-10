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

		pf := &ParamField{
			FieldName: name,
			Doc:       cleanFieldDoc(field.Doc),
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

// cleanFieldDoc extracts a clean help string from a go-swagger field comment.
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
// We strip the first line (field name), "Format:" lines, and trim whitespace.
func cleanFieldDoc(doc *ast.CommentGroup) string {
	if doc == nil {
		return ""
	}
	text := doc.Text() // already strips /* */ and // prefixes
	lines := strings.Split(strings.TrimSpace(text), "\n")
	if len(lines) == 0 {
		return ""
	}

	// First line is typically "FieldName." — skip it
	var kept []string
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i == 0 {
			// Skip the "FieldName." header line
			continue
		}
		// Skip "Format: xxx" lines (type hint, not useful for CLI help)
		if strings.HasPrefix(line, "Format:") {
			continue
		}
		if line != "" {
			kept = append(kept, line)
		}
	}
	return strings.Join(kept, " ")
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
