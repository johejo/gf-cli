package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

type MethodInfo struct {
	Name           string // "AddTeamRoleWithParams"
	ParamsTypeName string // "AddTeamRoleParams"
	ReturnTypeName string // "AddTeamRoleOK" (first return type, empty if error-only)
	NumReturns     int    // number of return values
	Doc            string // doc comment extracted from the non-WithParams implementation
}

// ParseService parses a ClientService interface from <baseDir>/client/<pkgName>/<pkgName>_client.go
// and returns the list of methods matching our criteria.
func ParseService(baseDir string, pkgName string) ([]*MethodInfo, error) {
	clientFile := filepath.Join(baseDir, "client", pkgName, pkgName+"_client.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, clientFile, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing service file %s: %w", clientFile, err)
	}

	// Build doc comment map from implementation methods: func (a *Client) Xxx(...)
	docMap := buildMethodDocMap(f)

	// Find ClientService interface
	var iface *ast.InterfaceType
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != "ClientService" {
				continue
			}
			iface, _ = ts.Type.(*ast.InterfaceType)
		}
	}
	if iface == nil {
		return nil, fmt.Errorf("ClientService interface not found in %s", clientFile)
	}

	var methods []*MethodInfo
	for _, method := range iface.Methods.List {
		if len(method.Names) == 0 {
			continue
		}
		name := method.Names[0].Name

		ft, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		// Skip deprecated ByID methods
		if strings.Contains(name, "ByID") {
			continue
		}
		// Skip datasource proxy methods
		if strings.Contains(name, "DatasourceProxy") {
			continue
		}

		// Only process methods whose first param type ends with "Params"
		if ft.Params == nil || len(ft.Params.List) == 0 {
			continue
		}
		paramsTypeName := extractTypeName(ft.Params.List[0].Type)
		if paramsTypeName == "" || !strings.HasSuffix(paramsTypeName, "Params") {
			continue
		}

		// Extract return types
		numReturns := 0
		var returnTypeName string
		if ft.Results != nil {
			numReturns = len(ft.Results.List)
			if numReturns >= 2 {
				returnTypeName = extractTypeName(ft.Results.List[0].Type)
			}
		}

		// Look up doc comment from the non-WithParams sibling method
		baseName := strings.TrimSuffix(name, "WithParams")
		doc := docMap[baseName]
		if doc == "" {
			doc = docMap[name]
		}

		methods = append(methods, &MethodInfo{
			Name:           name,
			ParamsTypeName: paramsTypeName,
			ReturnTypeName: returnTypeName,
			NumReturns:     numReturns,
			Doc:            doc,
		})
	}
	return methods, nil
}

// buildMethodDocMap extracts doc comments from method implementations
// (func (a *Client) MethodName(...)) and returns a map of method name -> cleaned doc string.
func buildMethodDocMap(f *ast.File) map[string]string {
	m := make(map[string]string)
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Doc == nil || fd.Recv == nil {
			continue
		}
		m[fd.Name.Name] = cleanDoc(fd.Doc.Text())
	}
	return m
}

// cleanDoc cleans a go-swagger doc comment.
// Input example:
//
//	"CreateDashboardSnapshot whens creating a snapshot using the API...\n\nSnapshot public mode should be enabled..."
//
// The first word is typically the method name repeated; we strip it and capitalize the rest.
func cleanDoc(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	// go-swagger comments start with "MethodName description...", strip the method name prefix
	// by finding the first space and taking the rest.
	if idx := strings.IndexByte(s, ' '); idx >= 0 {
		s = s[idx+1:]
	}
	s = strings.TrimSpace(s)
	// Take only the first sentence/line for Short help
	if idx := strings.IndexByte(s, '\n'); idx >= 0 {
		s = s[:idx]
	}
	s = strings.TrimSpace(s)
	// Capitalize first letter
	if len(s) > 0 {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

// extractTypeName extracts the type name from an AST expression.
// For *Foo it returns "Foo", for Foo it returns "Foo", for pkg.Foo it returns "Foo".
func extractTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return extractTypeName(t.X)
	case *ast.SelectorExpr:
		return t.Sel.Name
	default:
		return ""
	}
}
