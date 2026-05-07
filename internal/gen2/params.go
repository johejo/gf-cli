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
	FieldName  string // "TeamID"
	Type       string // "string", "int64", "bool", "[]string", "[]int64"
	IsPtr      bool
	IsBody     bool
	IsRequired bool   // derived from WriteToRequest: true when the Set*Param call is unconditional (no `if o.<Field> != nil` guard)
	ModelType  string // non-empty for body fields, e.g. "models.AddTeamRoleCommand"
	Doc        string // cleaned doc comment from source, e.g. "Search Query"
	Default    string // Go-literal default extracted from the doc comment, e.g. `1000` or `"View"`; empty if absent
	In         string // OpenAPI parameter location: "path", "query", "header", "body", "form", "file"; "" if not derivable
}

// ParseParams finds the *Params struct by type name within the package directory
// and extracts its exported fields. The returned ParamFields carry the
// parameter location ("path", "query", "header", "body", ...) when it can be
// derived from the WriteToRequest method on the same struct.
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
		fields, err := extractParamFields(st)
		if err != nil {
			return nil, err
		}
		inMap := extractParamIn(f, paramsTypeName)
		optMap := extractOptionalSet(f, paramsTypeName)
		for _, pf := range fields {
			if v, ok := inMap[pf.FieldName]; ok {
				pf.In = v
			}
			pf.IsRequired = !optMap[pf.FieldName]
		}
		return fields, nil
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

// extractParamIn locates the WriteToRequest method on *paramsTypeName and
// returns a map from Go field name to OpenAPI parameter location.
//
// go-swagger emits one r.Set{Path,Query,Header,Body,Form,File}Param call per
// field. Required value-type fields appear directly with `o.<Field>` in the
// call's args; optional pointer fields are wired through a temp variable
// inside an enclosing `if o.<Field> != nil { ... }` block. We catch both
// shapes by collecting `o.<Field>` selectors from (a) the call's own arguments
// and (b) the conditions of every enclosing IfStmt at the call site.
//
// Returns an empty map when WriteToRequest is not found (defensive — every
// generated Params struct in grafana-openapi-client-go has one).
func extractParamIn(f *ast.File, paramsTypeName string) map[string]string {
	out := make(map[string]string)
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Recv == nil || fd.Name.Name != "WriteToRequest" || fd.Body == nil {
			continue
		}
		if extractTypeName(fd.Recv.List[0].Type) != paramsTypeName {
			continue
		}
		// stack tracks ancestors of the node currently being visited.
		// ast.Inspect calls f(nil) once per non-nil node it descends into, so
		// the push-on-non-nil / pop-on-nil pairing stays balanced.
		var stack []ast.Node
		ast.Inspect(fd.Body, func(n ast.Node) bool {
			if n == nil {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
				return false
			}
			if call, ok := n.(*ast.CallExpr); ok {
				if in, ok := setParamIn(call); ok {
					seen := make(map[string]bool)
					for _, arg := range call.Args {
						for _, name := range collectOSelectors(arg) {
							seen[name] = true
						}
					}
					for _, anc := range stack {
						if ifs, ok := anc.(*ast.IfStmt); ok && ifs.Cond != nil {
							for _, name := range collectOSelectors(ifs.Cond) {
								seen[name] = true
							}
						}
					}
					for fname := range seen {
						if _, exists := out[fname]; !exists {
							out[fname] = in
						}
					}
				}
			}
			stack = append(stack, n)
			return true
		})
		return out
	}
	return out
}

// extractOptionalSet returns the set of field names that appear inside an
// `if o.<Field> != nil` (or `== nil`) guard anywhere in the WriteToRequest
// method body. go-swagger wraps every optional parameter — both pointer
// primitives and slice-typed query params — in such a guard, so the presence
// of a guard is a more reliable signal of optionality than checking IsPtr
// alone (which misclassifies optional slice query params as required because
// slices are emitted as value-type fields).
func extractOptionalSet(f *ast.File, paramsTypeName string) map[string]bool {
	out := make(map[string]bool)
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Recv == nil || fd.Name.Name != "WriteToRequest" || fd.Body == nil {
			continue
		}
		if extractTypeName(fd.Recv.List[0].Type) != paramsTypeName {
			continue
		}
		ast.Inspect(fd.Body, func(n ast.Node) bool {
			ifs, ok := n.(*ast.IfStmt)
			if !ok || ifs.Cond == nil {
				return true
			}
			be, ok := ifs.Cond.(*ast.BinaryExpr)
			if !ok {
				return true
			}
			if be.Op != token.NEQ && be.Op != token.EQL {
				return true
			}
			var sel *ast.SelectorExpr
			switch {
			case isNilIdent(be.Y):
				sel, _ = be.X.(*ast.SelectorExpr)
			case isNilIdent(be.X):
				sel, _ = be.Y.(*ast.SelectorExpr)
			}
			if sel == nil {
				return true
			}
			id, ok := sel.X.(*ast.Ident)
			if !ok || id.Name != "o" {
				return true
			}
			out[sel.Sel.Name] = true
			return true
		})
		return out
	}
	return out
}

func isNilIdent(e ast.Expr) bool {
	id, ok := e.(*ast.Ident)
	return ok && id.Name == "nil"
}

// setParamIn matches r.Set{Path,Query,Header,Body,Form,File}Param(...) and
// returns the lowercase parameter location.
func setParamIn(call *ast.CallExpr) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}
	rest, ok := strings.CutPrefix(sel.Sel.Name, "Set")
	if !ok {
		return "", false
	}
	rest, ok = strings.CutSuffix(rest, "Param")
	if !ok {
		return "", false
	}
	switch rest {
	case "Path", "Query", "Header", "Body", "Form", "File":
		return strings.ToLower(rest), true
	}
	return "", false
}

// collectOSelectors gathers field names from `o.<Ident>` SelectorExpr nodes
// inside expr. The receiver is named `o` by go-swagger convention.
func collectOSelectors(expr ast.Node) []string {
	var out []string
	ast.Inspect(expr, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		id, ok := sel.X.(*ast.Ident)
		if !ok || id.Name != "o" {
			return true
		}
		out = append(out, sel.Sel.Name)
		return true
	})
	return out
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
