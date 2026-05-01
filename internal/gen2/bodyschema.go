package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
	"strings"
)

func unwrapPointer(expr ast.Expr) ast.Expr {
	for {
		star, ok := expr.(*ast.StarExpr)
		if !ok {
			return expr
		}
		expr = star.X
	}
}

func basicTypeToJSON(name string) (string, bool) {
	switch name {
	case "string":
		return "string", true
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return "number", true
	case "bool":
		return "boolean", true
	default:
		return "", false
	}
}

func extractJSONName(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}
	tagStr := strings.Trim(field.Tag.Value, "`")
	tag := reflect.StructTag(tagStr)
	jsonTag := tag.Get("json")
	if jsonTag == "" {
		return ""
	}
	name, _, _ := strings.Cut(jsonTag, ",")
	return name
}

func hasRequiredAnnotation(doc *ast.CommentGroup) bool {
	if doc == nil {
		return false
	}
	for _, c := range doc.List {
		if strings.Contains(c.Text, "Required: true") {
			return true
		}
	}
	return false
}

// extractDescription returns the field's doc comment with go-swagger's known
// schema-annotation lines (Required:, Enum:, Format:, etc.) stripped out,
// so only the human-readable description remains. The original physical line
// breaks are preserved (kept lines are joined with "\n") so callers that
// render the description can lay out multi-point comments readably. If the
// result is a trivial restatement of jsonName (only case/whitespace differs),
// it is dropped as noise.
func extractDescription(doc *ast.CommentGroup, jsonName string) string {
	if doc == nil {
		return ""
	}
	var kept []string
	for line := range strings.SplitSeq(strings.TrimSpace(doc.Text()), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || isAnnotationLine(line) {
			continue
		}
		kept = append(kept, line)
	}
	desc := strings.Join(kept, "\n")
	if isTrivialDescription(desc, jsonName) {
		return ""
	}
	return desc
}

// isTrivialDescription reports whether desc is a restatement of jsonName after
// lowercasing and stripping non-alphanumerics. E.g. jsonName "roleUid" and
// desc "role Uid" both normalize to "roleuid".
func isTrivialDescription(desc, jsonName string) bool {
	if desc == "" {
		return true
	}
	return normalizeForCompare(desc) == normalizeForCompare(jsonName)
}

func normalizeForCompare(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r + ('a' - 'A'))
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		}
	}
	return b.String()
}

// isAnnotationLine reports whether a doc line is a go-swagger schema-annotation
// header that should be stripped from descriptions. Note `Description:` is
// itself a marker: go-swagger emits it on its own line and the lines that
// follow are the field's description body — so we strip the marker but keep
// the body via extractDescription's default path.
func isAnnotationLine(s string) bool {
	for _, p := range []string{
		"Required:", "Enum:", "Format:", "Pattern:",
		"Minimum:", "Maximum:",
		"MinLength:", "MaxLength:", "Min Length:", "Max Length:",
		"MinItems:", "MaxItems:", "Min Items:", "Max Items:",
		"UniqueItems:", "Unique Items:",
		"Example:", "Default:", "MultipleOf:", "Multiple Of:",
		"Description:", "Read Only:", "ReadOnly:",
		"+optional",
	} {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

func extractEnumValues(doc *ast.CommentGroup) []string {
	if doc == nil {
		return nil
	}
	for _, c := range doc.List {
		text := c.Text
		_, after, ok := strings.Cut(text, "Enum: [")
		if !ok {
			continue
		}
		rest := after
		before, _, ok := strings.Cut(rest, "]")
		if !ok {
			continue
		}
		inner := strings.TrimSpace(before)
		// Handle nested brackets like [[expired active pending]]
		inner = strings.Trim(inner, "[]")
		if inner == "" {
			continue
		}
		// go-swagger emits quoted, comma-separated values for string enums
		// (`"a","b","c"`); older versions used whitespace separation (`a b c`).
		if strings.Contains(inner, `"`) {
			var out []string
			for p := range strings.SplitSeq(inner, ",") {
				p = strings.TrimSpace(p)
				if unq, err := strconv.Unquote(p); err == nil {
					out = append(out, unq)
				} else if p != "" {
					out = append(out, p)
				}
			}
			return out
		}
		return strings.Fields(inner)
	}
	return nil
}

// BuildJSONSchemaFromExpr returns a pretty-printed JSON Schema (draft 2020-12)
// document describing an arbitrary Go type expression from the upstream client's
// response package (e.g. *models.TeamDTO, []*models.TeamMemberDTO, map[string]any),
// along with a typed field table covering the same fields and the schema's
// top-level description. Used for response Payload fields, where the declared
// type is not always a named struct in the models/ package.
//
// When frameNote is true, the schema's top-level description carries a warning
// that the response wire format differs from the Go type (models.Frame mismatch);
// the same string is returned separately so callers can surface it above the
// field table without parsing the schema.
func BuildJSONSchemaFromExpr(baseDir string, expr ast.Expr, title string, frameNote bool) (jsonSchema, annotations, rootDesc string, err error) {
	files, err := getModelParsedFiles(baseDir)
	if err != nil {
		return "", "", "", err
	}
	b := &jsonSchemaBuilder{files: files, visited: map[string]bool{}}
	inner := b.fieldSchema(expr)
	if inner == nil {
		return "", "", "", fmt.Errorf("empty schema for %s", title)
	}
	top := newOrderedObj()
	top.set("$schema", "https://json-schema.org/draft/2020-12/schema")
	if title != "" {
		top.set("title", title)
	}
	if frameNote {
		top.set("description", "Response wire format differs from the Go type; consider --raw.")
	}
	for _, k := range inner.keys {
		top.set(k, inner.values[k])
	}
	out, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		return "", "", "", err
	}
	return string(out), annotationsFromSchema(top), rootDescription(top), nil
}

// BuildBodyJSONSchema returns a pretty-printed JSON Schema (draft 2020-12)
// document describing the body struct referenced by modelType (e.g.
// "models.CreateTeamCommand"), along with a "REQUIRED / description / enum"
// annotation table covering the same fields. The schema is derived by walking
// the Go AST of the models package with full-depth recursion; self-referencing
// types are broken off with an opaque {"type": "object"}.
func BuildBodyJSONSchema(baseDir string, modelType string) (string, string, error) {
	typeName := strings.TrimPrefix(modelType, "models.")
	if typeName == modelType {
		return "", "", fmt.Errorf("unsupported model type: %s", modelType)
	}
	files, err := getModelParsedFiles(baseDir)
	if err != nil {
		return "", "", err
	}
	b := &jsonSchemaBuilder{files: files, visited: map[string]bool{}}
	inner := b.structSchema(typeName)
	if inner == nil {
		return "", "", fmt.Errorf("struct %s not found in models package", typeName)
	}
	top := newOrderedObj()
	top.set("$schema", "https://json-schema.org/draft/2020-12/schema")
	top.set("title", typeName)
	for _, k := range inner.keys {
		top.set(k, inner.values[k])
	}
	out, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		return "", "", err
	}
	return string(out), annotationsFromSchema(top), nil
}

// fieldRow buffers the data for a single property before width computation.
// The walker collects rows in source order; annotationsFromSchema then sizes
// the path and type columns to the longest values within the table.
type fieldRow struct {
	path     string
	typ      string
	required bool
	desc     string // description joined with ", enum: ..." when both are present
}

// annotationsFromSchema walks an already-built JSON Schema object and emits a
// scannable left-aligned table listing every property with its type and, when
// present, REQUIRED / description / enum metadata. The walker recurses into
// nested objects (joined with ".") and arrays of objects (parent path suffixed
// with "[]"), but does not recurse into maps — the parent row already carries
// "map<...>" so children would be redundant. Path and type column widths are
// computed per-table from the longest entries (mirroring how cobra/pflag sizes
// flag-help columns), so short tables stay tight and long paths don't overflow.
func annotationsFromSchema(schema *orderedObj) string {
	var rows []fieldRow
	walkSchemaAnnotations(&rows, schema, "")
	if len(rows) == 0 {
		return ""
	}
	pathW, typeW := 0, 0
	for _, r := range rows {
		if len(r.path) > pathW {
			pathW = len(r.path)
		}
		if len(r.typ) > typeW {
			typeW = len(r.typ)
		}
	}
	cont := strings.Repeat(" ", len("  ")+pathW+len("  ")+typeW+len("  "))
	var lines []string
	for _, r := range rows {
		appendRowLines(&lines, r, pathW, typeW, cont)
	}
	return strings.Join(lines, "\n")
}

// rootDescription returns the schema's top-level "description" field, or "".
// Used to surface the response Frame-mismatch advisory above the field table
// without baking it into the table itself.
func rootDescription(schema *orderedObj) string {
	if schema == nil {
		return ""
	}
	if v, ok := schema.values["description"].(string); ok {
		return v
	}
	return ""
}

func walkSchemaAnnotations(rows *[]fieldRow, schema *orderedObj, prefix string) {
	if schema == nil {
		return
	}
	propsAny, ok := schema.values["properties"]
	if !ok {
		return
	}
	props, ok := propsAny.(*orderedObj)
	if !ok {
		return
	}
	required := requiredSet(schema)
	for _, name := range props.keys {
		prop, ok := props.values[name].(*orderedObj)
		if !ok {
			continue
		}
		path := prefix + name
		*rows = append(*rows, collectFieldRow(path, prop, required[name]))
		// Recurse into a nested object schema.
		if _, ok := prop.values["properties"]; ok {
			walkSchemaAnnotations(rows, prop, path+".")
			continue
		}
		// Recurse into an array whose items are an object schema.
		if itemsAny, ok := prop.values["items"]; ok {
			if items, ok := itemsAny.(*orderedObj); ok {
				if _, ok := items.values["properties"]; ok {
					walkSchemaAnnotations(rows, items, path+"[].")
				}
			}
		}
	}
}

// renderType returns a short type label for a property schema, used in the
// table's type column. Composite shapes are flattened so the column width
// stays bounded: arrays as "array<inner>", maps as "map<string, value>", and
// nested object structs as plain "object" (their children appear as their own
// rows). Unknown shapes fall back to "object".
func renderType(prop *orderedObj) string {
	if prop == nil {
		return "object"
	}
	t, _ := prop.values["type"].(string)
	switch t {
	case "string", "number", "boolean":
		return t
	case "array":
		if itemsAny, ok := prop.values["items"]; ok {
			if items, ok := itemsAny.(*orderedObj); ok {
				return "array<" + renderType(items) + ">"
			}
		}
		return "array"
	case "object":
		if addAny, ok := prop.values["additionalProperties"]; ok {
			if add, ok := addAny.(*orderedObj); ok {
				return "map<string, " + renderType(add) + ">"
			}
			return "map<string, object>"
		}
		return "object"
	}
	return "object"
}

func requiredSet(schema *orderedObj) map[string]bool {
	out := map[string]bool{}
	reqAny, ok := schema.values["required"]
	if !ok {
		return out
	}
	reqs, ok := reqAny.([]any)
	if !ok {
		return out
	}
	for _, r := range reqs {
		if s, ok := r.(string); ok {
			out[s] = true
		}
	}
	return out
}

func collectFieldRow(path string, prop *orderedObj, required bool) fieldRow {
	var descParts []string
	if descAny, ok := prop.values["description"]; ok {
		if s, ok := descAny.(string); ok && s != "" {
			descParts = append(descParts, s)
		}
	}
	if enumAny, ok := prop.values["enum"]; ok {
		if vals, ok := enumAny.([]any); ok && len(vals) > 0 {
			ss := make([]string, 0, len(vals))
			for _, v := range vals {
				ss = append(ss, formatEnumValue(v))
			}
			descParts = append(descParts, "enum: "+strings.Join(ss, " | "))
		}
	}
	return fieldRow{
		path:     path,
		typ:      renderType(prop),
		required: required,
		desc:     strings.Join(descParts, ", "),
	}
}

func appendRowLines(lines *[]string, r fieldRow, pathW, typeW int, cont string) {
	switch {
	case r.required:
		*lines = append(*lines, formatRow(r.path, r.typ, "REQUIRED", pathW, typeW))
		if r.desc != "" {
			for c := range strings.SplitSeq(r.desc, "\n") {
				*lines = append(*lines, cont+c)
			}
		}
	case r.desc != "":
		first := true
		for line := range strings.SplitSeq(r.desc, "\n") {
			if first {
				*lines = append(*lines, formatRow(r.path, r.typ, line, pathW, typeW))
				first = false
				continue
			}
			*lines = append(*lines, cont+line)
		}
	default:
		*lines = append(*lines, formatRow(r.path, r.typ, "", pathW, typeW))
	}
}

// formatRow lays out a single field row. Columns are separated by two spaces
// so the boundary stays visible even when path is at the table's max width.
// The metadata cell is omitted entirely when empty so plain rows do not carry
// trailing whitespace; pathW/typeW pad the path and type columns so all rows
// in the same table align.
func formatRow(path, typ, meta string, pathW, typeW int) string {
	if meta == "" {
		return fmt.Sprintf("  %-*s  %s", pathW, path, typ)
	}
	return fmt.Sprintf("  %-*s  %-*s  %s", pathW, path, typeW, typ, meta)
}

type jsonSchemaBuilder struct {
	files   []*ast.File
	visited map[string]bool
}

func (b *jsonSchemaBuilder) structSchema(typeName string) *orderedObj {
	st := b.findStruct(typeName)
	if st == nil {
		return nil
	}
	if b.visited[typeName] {
		obj := newOrderedObj()
		obj.set("type", "object")
		return obj
	}
	b.visited[typeName] = true
	defer delete(b.visited, typeName)

	props := newOrderedObj()
	var required []string
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			// Anonymous embedded field: go-swagger renders OpenAPI `allOf`
			// composition as embedded structs; inline the embedded type's
			// properties/required into this schema.
			if name := embeddedTypeName(field.Type); name != "" {
				if inner := b.structSchema(name); inner != nil {
					mergeStructInto(props, &required, inner)
				}
			}
			continue
		}
		name := field.Names[0].Name
		if !ast.IsExported(name) {
			continue
		}
		jsonName := extractJSONName(field)
		if jsonName == "" || jsonName == "-" {
			continue
		}
		prop := b.fieldSchema(field.Type)
		if desc := extractDescription(field.Doc, jsonName); desc != "" {
			prop.set("description", desc)
		}
		if enums := extractEnumValues(field.Doc); len(enums) > 0 {
			prop.set("enum", coerceEnumValues(prop, enums))
		}
		props.set(jsonName, prop)
		if hasRequiredAnnotation(field.Doc) {
			required = append(required, jsonName)
		}
	}
	out := newOrderedObj()
	out.set("type", "object")
	out.set("properties", props)
	if len(required) > 0 {
		out.set("required", stringsToAny(required))
	}
	return out
}

func embeddedTypeName(expr ast.Expr) string {
	switch t := unwrapPointer(expr).(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return t.Sel.Name
	}
	return ""
}

func formatEnumValue(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprint(v)
}

// coerceEnumValues types the enum values to match the property's JSON Schema
// "type" so validators don't reject e.g. {"type":"number","enum":["1"]}.
// Unparseable numeric values are dropped rather than emitted as strings,
// which would produce a heterogeneous (silently invalid) enum array.
func coerceEnumValues(prop *orderedObj, vals []string) []any {
	typ, _ := prop.values["type"].(string)
	if typ == "number" || typ == "integer" {
		out := make([]any, 0, len(vals))
		for _, v := range vals {
			if n, err := strconv.ParseFloat(v, 64); err == nil {
				out = append(out, n)
			}
		}
		return out
	}
	return stringsToAny(vals)
}

func mergeStructInto(props *orderedObj, required *[]string, inner *orderedObj) {
	innerProps, ok := inner.values["properties"].(*orderedObj)
	if ok {
		for _, k := range innerProps.keys {
			if _, exists := props.values[k]; !exists {
				props.set(k, innerProps.values[k])
			}
		}
	}
	innerReq, ok := inner.values["required"].([]any)
	if ok {
		seen := map[string]bool{}
		for _, r := range *required {
			seen[r] = true
		}
		for _, r := range innerReq {
			if s, ok := r.(string); ok && !seen[s] {
				*required = append(*required, s)
				seen[s] = true
			}
		}
	}
}

func (b *jsonSchemaBuilder) fieldSchema(expr ast.Expr) *orderedObj {
	expr = unwrapPointer(expr)
	switch t := expr.(type) {
	case *ast.Ident:
		if jt, ok := basicTypeToJSON(t.Name); ok {
			obj := newOrderedObj()
			obj.set("type", jt)
			return obj
		}
		return b.namedTypeSchema(t.Name)
	case *ast.SelectorExpr:
		if pkg, ok := t.X.(*ast.Ident); ok && pkg.Name == "strfmt" {
			obj := newOrderedObj()
			obj.set("type", "string")
			return obj
		}
		return b.namedTypeSchema(t.Sel.Name)
	case *ast.ArrayType:
		obj := newOrderedObj()
		obj.set("type", "array")
		obj.set("items", b.fieldSchema(t.Elt))
		return obj
	case *ast.MapType:
		obj := newOrderedObj()
		obj.set("type", "object")
		obj.set("additionalProperties", b.fieldSchema(t.Value))
		return obj
	case *ast.InterfaceType:
		return newOrderedObj()
	}
	return newOrderedObj()
}

func (b *jsonSchemaBuilder) namedTypeSchema(typeName string) *orderedObj {
	for _, f := range b.files {
		for _, decl := range f.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok || ts.Name.Name != typeName {
					continue
				}
				switch ut := ts.Type.(type) {
				case *ast.Ident:
					if jt, ok := basicTypeToJSON(ut.Name); ok {
						obj := newOrderedObj()
						obj.set("type", jt)
						return obj
					}
					return b.namedTypeSchema(ut.Name)
				case *ast.SelectorExpr:
					if pkg, ok := ut.X.(*ast.Ident); ok && pkg.Name == "strfmt" {
						obj := newOrderedObj()
						obj.set("type", "string")
						return obj
					}
					return b.namedTypeSchema(ut.Sel.Name)
				case *ast.StructType:
					if s := b.structSchema(typeName); s != nil {
						return s
					}
					obj := newOrderedObj()
					obj.set("type", "object")
					return obj
				case *ast.ArrayType:
					obj := newOrderedObj()
					obj.set("type", "array")
					obj.set("items", b.fieldSchema(ut.Elt))
					return obj
				case *ast.MapType:
					obj := newOrderedObj()
					obj.set("type", "object")
					obj.set("additionalProperties", b.fieldSchema(ut.Value))
					return obj
				case *ast.InterfaceType:
					return newOrderedObj()
				}
			}
		}
	}
	obj := newOrderedObj()
	obj.set("type", "object")
	return obj
}

func (b *jsonSchemaBuilder) findStruct(typeName string) *ast.StructType {
	for _, f := range b.files {
		if st := findStructType(f, typeName); st != nil {
			return st
		}
	}
	return nil
}

// orderedObj is a JSON object whose keys are serialized in insertion order.
// The standard library's json.Marshal sorts map[string]any keys alphabetically,
// which makes generated schemas less readable and unstable compared to the
// source struct's field order.
type orderedObj struct {
	keys   []string
	values map[string]any
}

func newOrderedObj() *orderedObj {
	return &orderedObj{values: map[string]any{}}
}

func (o *orderedObj) set(k string, v any) {
	if _, ok := o.values[k]; !ok {
		o.keys = append(o.keys, k)
	}
	o.values[k] = v
}

func (o *orderedObj) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, k := range o.keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		kb, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		buf.Write(kb)
		buf.WriteByte(':')
		vb, err := json.Marshal(o.values[k])
		if err != nil {
			return nil, err
		}
		buf.Write(vb)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func stringsToAny(ss []string) []any {
	out := make([]any, len(ss))
	for i, s := range ss {
		out[i] = s
	}
	return out
}
