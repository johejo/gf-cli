package main

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"testing"
)

func TestParseModelSchema(t *testing.T) {
	baseDir := os.Getenv("GRAFANA_CLIENT_DIR")
	if baseDir == "" {
		t.Skip("GRAFANA_CLIENT_DIR not set")
	}
	schema, err := ParseModelSchema(baseDir, "models.CreateTeamCommand")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TypeName: %s", schema.TypeName)
	for _, f := range schema.Fields {
		t.Logf("  %s %s required=%v description=%q", f.JSONName, f.GoType, f.IsRequired, f.Description)
	}
	if len(schema.Fields) == 0 {
		t.Error("expected fields")
	}
}

func TestFormatBodySchema(t *testing.T) {
	schema := &BodySchemaInfo{
		TypeName: "CreateTeamCommand",
		Fields: []*ModelField{
			{JSONName: "name", GoType: "string", IsRequired: true},
			{JSONName: "email", GoType: "string"},
		},
	}
	got := formatBodySchema(schema)
	t.Log(got)
	if got == "" {
		t.Error("expected non-empty")
	}
}

func TestFormatBodySchemaNestedFields(t *testing.T) {
	schema := &BodySchemaInfo{
		TypeName: "CreateReport",
		Fields: []*ModelField{
			{
				JSONName: "options",
				JSONType: "object",
				NestedFields: []*ModelField{
					{JSONName: "layout", JSONType: "string"},
					{JSONName: "orientation", JSONType: "string"},
				},
			},
			{
				JSONName: "rules",
				JSONType: "object",
				IsArray:  true,
				NestedFields: []*ModelField{
					{JSONName: "alert", JSONType: "string"},
					{JSONName: "expr", JSONType: "string"},
				},
			},
		},
	}

	got := formatBodySchema(schema)
	want := `Body schema (CreateReport):
{
  "options": {
    "layout": string,
    "orientation": string
  },
  "rules": [
    {
      "alert": string,
      "expr": string
    }
  ]
}`
	if got != want {
		t.Fatalf("formatBodySchema() =\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatBodySchemaMultiLineDescription(t *testing.T) {
	schema := &BodySchemaInfo{
		TypeName: "MetricRequest",
		Fields: []*ModelField{
			{
				JSONName:   "queries",
				JSONType:   "any",
				IsArray:    true,
				IsRequired: true,
				Description: "queries.refId – Specifies an identifier of the query.\n" +
					"queries.datasourceId – Specifies the data source to be queried.\n" +
					"queries.maxDataPoints - Species maximum amount of data points.",
			},
		},
	}

	got := formatBodySchema(schema)
	want := `Body schema (MetricRequest):
{
  "queries": [any]
}
  queries                  REQUIRED
                           queries.refId – Specifies an identifier of the query.
                           queries.datasourceId – Specifies the data source to be queried.
                           queries.maxDataPoints - Species maximum amount of data points.`
	if got != want {
		t.Fatalf("formatBodySchema() =\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatBodySchemaGoTypeHint(t *testing.T) {
	schema := &BodySchemaInfo{
		TypeName: "CreateDashboardSnapshotCommand",
		Fields: []*ModelField{
			{JSONName: "dashboard", JSONType: "any", GoType: "Dashboard", IsRequired: true},
			{JSONName: "relativeTimeRange", JSONType: "object", GoType: "RelativeTimeRange"},
			{JSONName: "extra", JSONType: "any", GoType: ""},
			{JSONName: "panels", JSONType: "any", IsArray: true, GoType: "[]Dashboard"},
			{JSONName: "name", JSONType: "string", GoType: "string"},
			{JSONName: "blob", JSONType: "any", GoType: "object"},
			{JSONName: "freeform", JSONType: "any", GoType: "any"},
			{JSONName: "labels", JSONType: "object", IsMap: true, MapValueType: "string", GoType: "map[string]Dashboard"},
			{
				JSONName: "nested",
				JSONType: "object",
				GoType:   "ExpandedStruct",
				NestedFields: []*ModelField{
					{JSONName: "inner", JSONType: "string"},
				},
			},
		},
	}

	got := formatBodySchema(schema)
	want := `Body schema (CreateDashboardSnapshotCommand):
{
  "dashboard": any,  // models.Dashboard
  "relativeTimeRange": object,  // models.RelativeTimeRange
  "extra": any,
  "panels": [any],  // []models.Dashboard
  "name": string,
  "blob": any,
  "freeform": any,
  "labels": {"key": string},
  "nested": {
    "inner": string
  }
}
  dashboard                REQUIRED`
	if got != want {
		t.Fatalf("formatBodySchema() =\n%s\nwant:\n%s", got, want)
	}
}

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
	if _, err := BuildBodyJSONSchema("/does/not/matter", "NotAModel"); err == nil {
		t.Fatal("expected error for non-models.* modelType")
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
