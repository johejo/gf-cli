package main

import (
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
		t.Logf("  %s %s required=%v", f.JSONName, f.GoType, f.IsRequired)
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
