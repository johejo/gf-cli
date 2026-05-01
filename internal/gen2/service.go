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
// and returns the service-level short description and the list of methods matching our criteria.
func ParseService(baseDir string, pkgName string) (string, []*MethodInfo, error) {
	clientFile := filepath.Join(baseDir, "client", pkgName, pkgName+"_client.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, clientFile, nil, parser.ParseComments)
	if err != nil {
		return "", nil, fmt.Errorf("parsing service file %s: %w", clientFile, err)
	}

	short := extractServiceShort(f)

	// Build doc comment map from implementation methods: func (a *Client) Xxx(...)
	docMap := buildMethodDocMap(f)

	// Find ClientService interface
	iface := findClientServiceInterface(f)
	if iface == nil {
		return "", nil, fmt.Errorf("ClientService interface not found in %s", clientFile)
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
	return short, methods, nil
}

// extractServiceShort reads the doc comment on the Client struct
// (e.g. "Client for access control API") and turns it into a cobra Short
// (e.g. "Access control API"). Returns "" if no usable comment is found.
func extractServiceShort(f *ast.File) string {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != "Client" {
				continue
			}
			doc := ""
			if gd.Doc != nil {
				doc = gd.Doc.Text()
			}
			if doc == "" && ts.Doc != nil {
				doc = ts.Doc.Text()
			}
			return formatServiceShort(doc)
		}
	}
	return ""
}

func formatServiceShort(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	if idx := strings.IndexByte(s, '\n'); idx >= 0 {
		s = strings.TrimSpace(s[:idx])
	}
	s = strings.TrimPrefix(s, "Client for ")
	return capitalizeFirst(s)
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
//
// go-swagger sometimes emits a blank line between the doc block and the func
// declaration, which makes Go's AST parser leave fd.Doc unset. We recover
// those by scanning free-floating comment groups whose first word matches the
// upcoming method name.
func buildMethodDocMap(f *ast.File) map[string]MethodDoc {
	m := make(map[string]MethodDoc)

	attached := make(map[*ast.CommentGroup]bool)
	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Doc != nil {
				attached[d.Doc] = true
				if d.Recv != nil {
					m[d.Name.Name] = splitDoc(d.Doc.Text())
				}
			}
		case *ast.GenDecl:
			if d.Doc != nil {
				attached[d.Doc] = true
			}
		}
	}

	var prev ast.Node
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Recv == nil {
			prev = decl
			continue
		}
		if _, ok := m[fd.Name.Name]; ok {
			prev = decl
			continue
		}
		var lower token.Pos
		if prev != nil {
			lower = prev.End()
		}
		upper := fd.Pos()
		var best *ast.CommentGroup
		for _, cg := range f.Comments {
			if attached[cg] {
				continue
			}
			if cg.Pos() <= lower || cg.End() >= upper {
				continue
			}
			if best == nil || cg.Pos() > best.Pos() {
				best = cg
			}
		}
		if best != nil {
			text := strings.TrimSpace(best.Text())
			if docFirstWord(text) == fd.Name.Name {
				m[fd.Name.Name] = splitDoc(text)
			}
		}
		prev = decl
	}
	return m
}

func docFirstWord(s string) string {
	if idx := strings.IndexAny(s, " \t\n"); idx >= 0 {
		return s[:idx]
	}
	return s
}

// shortMaxLen is the upper bound on a generated Short. Anything longer is
// treated as upstream prose that leaked into the OpenAPI summary slot;
// Short is cleared and the text is promoted into Long instead.
const shortMaxLen = 120

// splitDoc cleans a go-swagger doc comment and returns short and long help text.
// Input example:
//
//	"CreateDashboardSnapshot whens creating a snapshot using the API...\n\nSnapshot public mode should be enabled..."
//
// The first word is typically the method name repeated; we strip it and split on paragraph boundaries.
//
// When the first paragraph spans multiple source lines or exceeds shortMaxLen
// after the strip, the upstream OpenAPI summary almost certainly leaked a
// description into the summary slot. We then suppress Short and promote the
// paragraph into Long so agent listings don't surface a broken one-liner.
func splitDoc(s string) MethodDoc {
	s = strings.TrimSpace(s)
	if s == "" {
		return MethodDoc{}
	}

	paragraphs := normalizeDocParagraphs(s)
	if len(paragraphs) == 0 {
		return MethodDoc{}
	}

	first := capitalizeFirst(paragraphs[0].Text)
	var short string
	longParagraphs := make([]string, 0, len(paragraphs))
	if paragraphs[0].MultiLine || len(first) > shortMaxLen {
		longParagraphs = append(longParagraphs, first)
	} else {
		short = first
	}
	for _, p := range paragraphs[1:] {
		longParagraphs = append(longParagraphs, p.Text)
	}

	return MethodDoc{
		Short: short,
		Long:  strings.Join(longParagraphs, "\n\n"),
	}
}

type docParagraph struct {
	Text      string
	MultiLine bool
}

func normalizeDocParagraphs(s string) []docParagraph {
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

	var paragraphs []docParagraph
	var current []string
	flush := func() {
		if len(current) == 0 {
			return
		}
		paragraphs = append(paragraphs, docParagraph{
			Text:      strings.Join(current, " "),
			MultiLine: len(current) > 1,
		})
		current = nil
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			flush()
			continue
		}
		current = append(current, line)
	}
	flush()
	return paragraphs
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
