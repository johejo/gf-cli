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
	Short          string // single-line summary extracted from the non-WithParams implementation
	Long           string // detailed help extracted from the non-WithParams implementation
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
	iface := findClientServiceInterface(f)
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
		doc, ok := docMap[baseName]
		if !ok {
			doc = docMap[name]
		}

		methods = append(methods, &MethodInfo{
			Name:           name,
			ParamsTypeName: paramsTypeName,
			ReturnTypeName: returnTypeName,
			NumReturns:     numReturns,
			Short:          doc.Short,
			Long:           doc.Long,
		})
	}
	return methods, nil
}

type MethodDoc struct {
	Short string
	Long  string
}

func findClientServiceInterface(f *ast.File) *ast.InterfaceType {
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
			iface, ok := ts.Type.(*ast.InterfaceType)
			if ok {
				return iface
			}
		}
	}
	return nil
}

// buildMethodDocMap extracts doc comments from method implementations
// (func (a *Client) MethodName(...)) and returns a map of method name -> parsed doc parts.
func buildMethodDocMap(f *ast.File) map[string]MethodDoc {
	m := make(map[string]MethodDoc)
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Doc == nil || fd.Recv == nil {
			continue
		}
		m[fd.Name.Name] = splitDoc(fd.Doc.Text())
	}
	return m
}

// splitDoc cleans a go-swagger doc comment and returns short and long help text.
// Input example:
//
//	"CreateDashboardSnapshot whens creating a snapshot using the API...\n\nSnapshot public mode should be enabled..."
//
// The first word is typically the method name repeated; we strip it and split on paragraph boundaries.
func splitDoc(s string) MethodDoc {
	s = strings.TrimSpace(s)
	if s == "" {
		return MethodDoc{}
	}

	paragraphs := normalizeDocParagraphs(s)
	if len(paragraphs) == 0 {
		return MethodDoc{}
	}

	short := firstLine(paragraphs[0])
	longParagraphs := make([]string, 0, len(paragraphs))
	if paragraphs[0] != short {
		longParagraphs = append(longParagraphs, paragraphs[0])
	}
	if len(paragraphs) > 1 {
		longParagraphs = append(longParagraphs, paragraphs[1:]...)
	}

	return MethodDoc{
		Short: capitalizeFirst(short),
		Long:  strings.Join(longParagraphs, "\n\n"),
	}
}

func normalizeDocParagraphs(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(strings.TrimSpace(s), "\n")
	if len(lines) == 0 {
		return nil
	}

	if idx := strings.IndexByte(lines[0], ' '); idx >= 0 {
		lines[0] = strings.TrimSpace(lines[0][idx+1:])
	} else {
		lines[0] = ""
	}

	var paragraphs []string
	var current []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if len(current) > 0 {
				paragraphs = append(paragraphs, strings.Join(current, " "))
				current = nil
			}
			continue
		}
		current = append(current, line)
	}
	if len(current) > 0 {
		paragraphs = append(paragraphs, strings.Join(current, " "))
	}
	return paragraphs
}

func firstLine(s string) string {
	s = strings.TrimSpace(s)
	if idx := strings.IndexByte(s, '\n'); idx >= 0 {
		s = s[:idx]
	}
	return strings.TrimSpace(s)
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
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
