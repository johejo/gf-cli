package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

// ParseModelSchema parses a model struct from the models/ package and extracts
// field metadata for displaying body schema in CLI help.
func ParseModelSchema(baseDir string, modelType string) (*BodySchemaInfo, error) {
	typeName := strings.TrimPrefix(modelType, "models.")
	if typeName == modelType {
		return nil, fmt.Errorf("unsupported model type: %s", modelType)
	}

	files, err := getModelParsedFiles(baseDir)
	if err != nil {
		return nil, err
	}

	fields, err := parseStructFields(files, typeName, true)
	if err != nil {
		return nil, err
	}
	return &BodySchemaInfo{
		TypeName: typeName,
		Fields:   fields,
	}, nil
}

// parseStructFields extracts fields from a named struct.
// If expand is true, nested model types are expanded 1 level.
func parseStructFields(files []*ast.File, typeName string, expand bool) ([]*ModelField, error) {
	for _, f := range files {
		st := findStructType(f, typeName)
		if st == nil {
			continue
		}
		return extractModelFields(files, st, expand), nil
	}
	return nil, fmt.Errorf("struct %s not found in models package", typeName)
}

func extractModelFields(files []*ast.File, st *ast.StructType, expand bool) []*ModelField {
	var fields []*ModelField
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
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

		mf := &ModelField{
			JSONName:   jsonName,
			GoType:     resolveDisplayType(field.Type),
			IsRequired: hasRequiredAnnotation(field.Doc),
			EnumValues: extractEnumValues(field.Doc),
		}

		classifyField(mf, field.Type, files, expand)
		fields = append(fields, mf)
	}
	return fields
}

// classifyField resolves the JSON type, array/map flags, and nested fields.
func classifyField(mf *ModelField, expr ast.Expr, files []*ast.File, expand bool) {
	expr = unwrapPointer(expr)

	switch t := expr.(type) {
	case *ast.Ident:
		if jt, ok := basicTypeToJSON(t.Name); ok {
			mf.JSONType = jt
		} else {
			// Model type alias or struct in same package.
			mf.JSONType = resolveModelTypeJSON(files, t.Name, mf, expand)
		}
	case *ast.SelectorExpr:
		// e.g. strfmt.DateTime, models.Foo
		selName := t.Sel.Name
		if pkgIdent, ok := t.X.(*ast.Ident); ok && pkgIdent.Name == "strfmt" {
			mf.JSONType = "string"
			return
		}
		mf.JSONType = resolveModelTypeJSON(files, selName, mf, expand)
	case *ast.ArrayType:
		mf.IsArray = true
		inner := &ModelField{}
		classifyField(inner, t.Elt, files, expand)
		mf.JSONType = inner.JSONType
		mf.NestedFields = inner.NestedFields
	case *ast.MapType:
		mf.IsMap = true
		mf.JSONType = "object"
		valField := &ModelField{}
		classifyField(valField, t.Value, files, false)
		mf.MapValueType = valField.JSONType
	case *ast.InterfaceType:
		mf.JSONType = "any"
	default:
		mf.JSONType = "any"
	}
}

func unwrapPointer(expr ast.Expr) ast.Expr {
	for {
		star, ok := expr.(*ast.StarExpr)
		if !ok {
			return expr
		}
		expr = star.X
	}
}

// resolveModelTypeJSON tries to resolve a models-package type name.
// If it's a type alias for a basic type, return the JSON type.
// If it's a struct, optionally expand its fields.
func resolveModelTypeJSON(files []*ast.File, typeName string, mf *ModelField, expand bool) string {
	for _, f := range files {
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
						return jt
					}
				case *ast.StructType:
					if expand {
						mf.NestedFields = extractModelFields(files, ut, false)
					}
					return "object"
				case *ast.InterfaceType:
					return "any"
				case *ast.ArrayType:
					mf.IsArray = true
					inner := &ModelField{}
					classifyField(inner, ut.Elt, files, false)
					mf.JSONType = inner.JSONType
					mf.NestedFields = inner.NestedFields
					return inner.JSONType
				}
			}
		}
	}
	return "object"
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

func resolveDisplayType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return resolveDisplayType(t.X)
	case *ast.ArrayType:
		return "[]" + resolveDisplayType(t.Elt)
	case *ast.SelectorExpr:
		return t.Sel.Name
	case *ast.MapType:
		return "map[" + resolveDisplayType(t.Key) + "]" + resolveDisplayType(t.Value)
	case *ast.InterfaceType:
		return "object"
	default:
		return "any"
	}
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
		return strings.Fields(inner)
	}
	return nil
}

// formatBodySchema produces a JSON template + field annotations.
func formatBodySchema(schema *BodySchemaInfo) string {
	if schema == nil || len(schema.Fields) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Body schema (%s):", schema.TypeName)
	b.WriteByte('\n')
	writeJSONTemplate(&b, schema.Fields, 2)

	// Annotations for fields with metadata (required, enum).
	if ann := formatAnnotations(schema.Fields, ""); ann != "" {
		b.WriteByte('\n')
		b.WriteString(ann)
	}

	return b.String()
}

func writeJSONTemplate(b *strings.Builder, fields []*ModelField, indent int) {
	prefix := strings.Repeat(" ", indent)
	b.WriteString("{\n")
	for i, f := range fields {
		b.WriteString(prefix)
		fmt.Fprintf(b, `"%s": `, f.JSONName)
		writeFieldPlaceholder(b, f, indent)
		if i < len(fields)-1 {
			b.WriteByte(',')
		}
		b.WriteByte('\n')
	}
	b.WriteString(strings.Repeat(" ", max(0, indent-2)))
	b.WriteByte('}')
}

func writeFieldPlaceholder(b *strings.Builder, f *ModelField, indent int) {
	if f.IsMap {
		vt := f.MapValueType
		if vt == "" {
			vt = "any"
		}
		fmt.Fprintf(b, `{"key": %s}`, vt)
		return
	}

	if f.IsArray && len(f.NestedFields) > 0 {
		b.WriteString("[\n")
		b.WriteString(strings.Repeat(" ", indent+2))
		writeJSONTemplate(b, f.NestedFields, indent+4)
		b.WriteByte('\n')
		b.WriteString(strings.Repeat(" ", indent))
		b.WriteByte(']')
		return
	}

	if f.IsArray {
		fmt.Fprintf(b, "[%s]", f.JSONType)
		return
	}

	if len(f.NestedFields) > 0 {
		writeJSONTemplate(b, f.NestedFields, indent+2)
		return
	}

	b.WriteString(f.JSONType)
}

func formatAnnotations(fields []*ModelField, prefix string) string {
	var lines []string
	for _, f := range fields {
		var parts []string
		key := prefix + f.JSONName
		if f.IsRequired {
			parts = append(parts, "required")
		}
		if len(f.EnumValues) > 0 {
			parts = append(parts, "enum: "+strings.Join(f.EnumValues, " | "))
		}
		if len(parts) > 0 {
			lines = append(lines, fmt.Sprintf("  %-24s %s", key, strings.Join(parts, ", ")))
		}
		// Recurse into nested fields for annotations.
		if len(f.NestedFields) > 0 {
			nestedPrefix := key
			if f.IsArray {
				nestedPrefix += "[]"
			}
			if ann := formatAnnotations(f.NestedFields, nestedPrefix+"."); ann != "" {
				lines = append(lines, ann)
			}
		}
	}
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\n")
}

// BuildBodyJSONSchema returns a pretty-printed JSON Schema (draft 2020-12)
// document describing the body struct referenced by modelType (e.g. "models.CreateTeamCommand").
// The schema is derived by walking the Go AST of the models package with full-depth
// recursion; self-referencing types are broken off with an opaque {"type": "object"}.
func BuildBodyJSONSchema(baseDir string, modelType string) (string, error) {
	typeName := strings.TrimPrefix(modelType, "models.")
	if typeName == modelType {
		return "", fmt.Errorf("unsupported model type: %s", modelType)
	}
	files, err := getModelParsedFiles(baseDir)
	if err != nil {
		return "", err
	}
	b := &jsonSchemaBuilder{files: files, visited: map[string]bool{}}
	inner := b.structSchema(typeName)
	if inner == nil {
		return "", fmt.Errorf("struct %s not found in models package", typeName)
	}
	top := newOrderedObj()
	top.set("$schema", "https://json-schema.org/draft/2020-12/schema")
	top.set("title", typeName)
	for _, k := range inner.keys {
		top.set(k, inner.values[k])
	}
	out, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
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
		if enums := extractEnumValues(field.Doc); len(enums) > 0 {
			prop.set("enum", stringsToAny(enums))
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
