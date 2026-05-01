package main

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestBuildBodyJSONSchema(t *testing.T) {
	src := `package models

import "github.com/go-openapi/strfmt"

// CreateThing create thing
type CreateThing struct {
	// name
	// Required: true
	Name string ` + "`json:\"name\"`" + `

	// Email
	// address of the user
	Email *string ` + "`json:\"email,omitempty\"`" + `

	// Enum: [active pending expired]
	Status string ` + "`json:\"status\"`" + `

	// Tags
	Tags []string ` + "`json:\"tags,omitempty\"`" + `

	// Meta
	Meta map[string]string ` + "`json:\"meta,omitempty\"`" + `

	// Options
	Options *ThingOptions ` + "`json:\"options,omitempty\"`" + `

	// Permissions
	Permissions []*Permission ` + "`json:\"permissions,omitempty\"`" + `

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime ` + "`json:\"createdAt,omitempty\"`" + `

	// Extra arbitrary blob
	Extra interface{} ` + "`json:\"extra,omitempty\"`" + `

	// Skipped
	Skipped string ` + "`json:\"-\"`" + `
}

type ThingOptions struct {
	Layout string ` + "`json:\"layout,omitempty\"`" + `
}

type Permission struct {
	// Required: true
	Action string ` + "`json:\"action\"`" + `
	Scope  string ` + "`json:\"scope,omitempty\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "models.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	b := &jsonSchemaBuilder{files: []*ast.File{f}, visited: map[string]bool{}}
	inner := b.structSchema("CreateThing")
	if inner == nil {
		t.Fatal("structSchema returned nil")
	}
	top := newOrderedObj()
	top.set("$schema", "https://json-schema.org/draft/2020-12/schema")
	top.set("title", "CreateThing")
	for _, k := range inner.keys {
		top.set(k, inner.values[k])
	}
	out, err := jsonIndent(top)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	got := string(out)
	want := `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "CreateThing",
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    },
    "email": {
      "type": "string",
      "description": "Email\naddress of the user"
    },
    "status": {
      "type": "string",
      "enum": [
        "active",
        "pending",
        "expired"
      ]
    },
    "tags": {
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "meta": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "options": {
      "type": "object",
      "properties": {
        "layout": {
          "type": "string"
        }
      }
    },
    "permissions": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "action": {
            "type": "string"
          },
          "scope": {
            "type": "string"
          }
        },
        "required": [
          "action"
        ]
      }
    },
    "createdAt": {
      "type": "string"
    },
    "extra": {
      "description": "Extra arbitrary blob"
    }
  },
  "required": [
    "name"
  ]
}`
	if got != want {
		t.Fatalf("schema mismatch\n got: %s\nwant: %s", got, want)
	}
}

func TestBuildBodyJSONSchemaCycle(t *testing.T) {
	src := `package models

type Node struct {
	Name     string  ` + "`json:\"name\"`" + `
	Parent   *Node   ` + "`json:\"parent,omitempty\"`" + `
	Children []*Node ` + "`json:\"children,omitempty\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "models.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	b := &jsonSchemaBuilder{files: []*ast.File{f}, visited: map[string]bool{}}
	inner := b.structSchema("Node")
	if inner == nil {
		t.Fatal("structSchema returned nil")
	}
	out, err := jsonIndent(inner)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	got := string(out)
	want := `{
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    },
    "parent": {
      "type": "object"
    },
    "children": {
      "type": "array",
      "items": {
        "type": "object"
      }
    }
  }
}`
	if got != want {
		t.Fatalf("cycle schema mismatch\n got: %s\nwant: %s", got, want)
	}
}

func jsonIndent(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func TestBuildBodyJSONSchemaNamedAliases(t *testing.T) {
	src := `package models

// Status is a string alias used as an enum-carrying type.
type Status string

// Counts is a map alias.
type Counts map[string]int

type Config struct {
	Status Status ` + "`json:\"status,omitempty\"`" + `
	Counts Counts ` + "`json:\"counts,omitempty\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "models.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	b := &jsonSchemaBuilder{files: []*ast.File{f}, visited: map[string]bool{}}
	out, err := jsonIndent(b.structSchema("Config"))
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	got := string(out)
	want := `{
  "type": "object",
  "properties": {
    "status": {
      "type": "string"
    },
    "counts": {
      "type": "object",
      "additionalProperties": {
        "type": "number"
      }
    }
  }
}`
	if got != want {
		t.Fatalf("schema mismatch\n got: %s\nwant: %s", got, want)
	}
}

func TestBuildBodyJSONSchemaInvalidModelType(t *testing.T) {
	if _, _, err := BuildBodyJSONSchema("/does/not/matter", "NotAModel"); err == nil {
		t.Fatal("expected error for non-models.* modelType")
	}
}

func TestAnnotationsFromSchema(t *testing.T) {
	src := `package models

type Permission struct {
	// Required: true
	Action string ` + "`json:\"action\"`" + `
	Scope  string ` + "`json:\"scope,omitempty\"`" + `
}

// CreateThing create thing
type CreateThing struct {
	// Required: true
	Name string ` + "`json:\"name\"`" + `

	// Email
	// address of the user
	Email *string ` + "`json:\"email,omitempty\"`" + `

	// Enum: [active pending expired]
	Status string ` + "`json:\"status\"`" + `

	// Permissions
	Permissions []*Permission ` + "`json:\"permissions,omitempty\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "models.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	b := &jsonSchemaBuilder{files: []*ast.File{f}, visited: map[string]bool{}}
	inner := b.structSchema("CreateThing")
	if inner == nil {
		t.Fatal("structSchema returned nil")
	}
	got := annotationsFromSchema(inner)
	want := `  name                  string         REQUIRED
  email                 string         Email
                                       address of the user
  status                string         enum: active | pending | expired
  permissions           array<object>
  permissions[].action  string         REQUIRED
  permissions[].scope   string`
	if got != want {
		t.Fatalf("annotation mismatch\n got: %q\nwant: %q", got, want)
	}
}

func TestExtractEnumValues(t *testing.T) {
	t.Parallel()

	src := `package models

type T struct {
	// Enum: [active pending expired]
	A string ` + "`json:\"a\"`" + `

	// Enum: ["None","Viewer","Editor","Admin"]
	B string ` + "`json:\"b\"`" + `

	// Enum: [[expired active pending]]
	C string ` + "`json:\"c\"`" + `

	// Enum: [1]
	D int64 ` + "`json:\"d\"`" + `

	// Enum: ["regex","logfmt"]
	E string ` + "`json:\"e\"`" + `

	// no enum
	F string ` + "`json:\"f\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "t.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	st := findStructType(f, "T")
	if st == nil {
		t.Fatal("struct T not found")
	}
	got := make(map[string][]string)
	for _, field := range st.Fields.List {
		got[field.Names[0].Name] = extractEnumValues(field.Doc)
	}
	want := map[string][]string{
		"A": {"active", "pending", "expired"},
		"B": {"None", "Viewer", "Editor", "Admin"},
		"C": {"expired", "active", "pending"},
		"D": {"1"},
		"E": {"regex", "logfmt"},
		"F": nil,
	}
	for k, w := range want {
		g := got[k]
		if len(g) != len(w) {
			t.Errorf("%s: got %v, want %v", k, g, w)
			continue
		}
		for i := range w {
			if g[i] != w[i] {
				t.Errorf("%s[%d]: got %q, want %q", k, i, g[i], w[i])
			}
		}
	}
}

func TestBuildBodyJSONSchemaEmbeddedAllOf(t *testing.T) {
	src := `package models

type AllOf0 struct {
	Dashboard string ` + "`json:\"dashboard,omitempty\"`" + `
}

type AllOf1 struct {
	Meta string ` + "`json:\"meta,omitempty\"`" + `

	// Required: true
	Title string ` + "`json:\"title\"`" + `
}

type Combined struct {
	AllOf0
	AllOf1
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "models.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	b := &jsonSchemaBuilder{files: []*ast.File{f}, visited: map[string]bool{}}
	out, err := jsonIndent(b.structSchema("Combined"))
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	want := `{
  "type": "object",
  "properties": {
    "dashboard": {
      "type": "string"
    },
    "meta": {
      "type": "string"
    },
    "title": {
      "type": "string"
    }
  },
  "required": [
    "title"
  ]
}`
	if string(out) != want {
		t.Fatalf("schema mismatch\n got: %s\nwant: %s", out, want)
	}
}

func TestBuildBodyJSONSchemaTypedEnum(t *testing.T) {
	src := `package models

type T struct {
	// Enum: [1]
	Kind int64 ` + "`json:\"kind,omitempty\"`" + `

	// Enum: ["a","b"]
	Mode string ` + "`json:\"mode,omitempty\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "models.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	b := &jsonSchemaBuilder{files: []*ast.File{f}, visited: map[string]bool{}}
	out, err := jsonIndent(b.structSchema("T"))
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	want := `{
  "type": "object",
  "properties": {
    "kind": {
      "type": "number",
      "enum": [
        1
      ]
    },
    "mode": {
      "type": "string",
      "enum": [
        "a",
        "b"
      ]
    }
  }
}`
	if string(out) != want {
		t.Fatalf("schema mismatch\n got: %s\nwant: %s", out, want)
	}
}

func TestRenderType(t *testing.T) {
	t.Parallel()

	scalarString := newOrderedObj()
	scalarString.set("type", "string")

	scalarBool := newOrderedObj()
	scalarBool.set("type", "boolean")

	arrayOfString := newOrderedObj()
	arrayOfString.set("type", "array")
	arrayOfString.set("items", scalarString)

	objWithProps := newOrderedObj()
	objWithProps.set("type", "object")
	objWithProps.set("properties", newOrderedObj())

	arrayOfObject := newOrderedObj()
	arrayOfObject.set("type", "array")
	arrayOfObject.set("items", objWithProps)

	mapOfString := newOrderedObj()
	mapOfString.set("type", "object")
	mapOfString.set("additionalProperties", scalarString)

	bareObject := newOrderedObj()
	bareObject.set("type", "object")

	bareArray := newOrderedObj()
	bareArray.set("type", "array")

	missingType := newOrderedObj()

	tests := []struct {
		name string
		prop *orderedObj
		want string
	}{
		{"scalar string", scalarString, "string"},
		{"scalar boolean", scalarBool, "boolean"},
		{"array of string", arrayOfString, "array<string>"},
		{"array of object", arrayOfObject, "array<object>"},
		{"object with properties", objWithProps, "object"},
		{"map of string", mapOfString, "map<string, string>"},
		{"bare object", bareObject, "object"},
		{"bare array", bareArray, "array"},
		{"missing type", missingType, "object"},
		{"nil prop", nil, "object"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := renderType(tt.prop); got != tt.want {
				t.Fatalf("renderType(%s) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestRootDescription(t *testing.T) {
	t.Parallel()

	withDesc := newOrderedObj()
	withDesc.set("description", "Response wire format differs from the Go type; consider --raw.")

	withoutDesc := newOrderedObj()
	withoutDesc.set("type", "object")

	if got := rootDescription(withDesc); got != "Response wire format differs from the Go type; consider --raw." {
		t.Errorf("rootDescription(withDesc) = %q, want advisory string", got)
	}
	if got := rootDescription(withoutDesc); got != "" {
		t.Errorf("rootDescription(withoutDesc) = %q, want empty", got)
	}
	if got := rootDescription(nil); got != "" {
		t.Errorf("rootDescription(nil) = %q, want empty", got)
	}
}

func TestExtractDescriptionSkipsAnnotations(t *testing.T) {
	src := `package models

type T struct {
	// first line
	// continues here
	// Required: true
	// Format: date-time
	// Enum: [a b c]
	// Pattern: ^.*$
	// Example: foo
	A string ` + "`json:\"a\"`" + `

	B string ` + "`json:\"b\"`" + `

	// Required: true
	C string ` + "`json:\"c\"`" + `

	// role Uid
	RoleUID string ` + "`json:\"roleUid\"`" + `
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "t.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	st := findStructType(f, "T")
	if st == nil {
		t.Fatal("struct T not found")
	}
	got := make(map[string]string)
	for _, field := range st.Fields.List {
		got[field.Names[0].Name] = extractDescription(field.Doc, extractJSONName(field))
	}
	want := map[string]string{
		"A":       "first line\ncontinues here",
		"B":       "",
		"C":       "",
		"RoleUID": "", // "role Uid" normalizes to the same as "roleUid" -> dropped
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s: got %q, want %q", k, got[k], v)
		}
	}
}
