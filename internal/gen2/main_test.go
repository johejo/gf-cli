package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestFlagHelp(t *testing.T) {
	const bodyDoc = `Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{\"foo\": \"bar\"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.`

	tests := []struct {
		name       string
		fieldName  string
		doc        string
		isRequired bool
		want       string
	}{
		{
			name:       "doc with required",
			fieldName:  "UID",
			doc:        "Alert rule UID",
			isRequired: true,
			want:       `"Alert rule UID [required]"`,
		},
		{
			name:       "doc without required",
			fieldName:  "UID",
			doc:        "Alert rule UID",
			isRequired: false,
			want:       `"Alert rule UID"`,
		},
		{
			name:       "empty doc falls back to fieldName, required",
			fieldName:  "Name",
			doc:        "",
			isRequired: true,
			want:       `"Name [required]"`,
		},
		{
			name:       "empty doc falls back to fieldName, optional",
			fieldName:  "Name",
			doc:        "",
			isRequired: false,
			want:       `"Name"`,
		},
		{
			name:       "Body field always required",
			fieldName:  "Body",
			doc:        "ignored",
			isRequired: true,
			want:       `"` + bodyDoc + ` [required]"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flagHelp(tt.fieldName, tt.doc, tt.isRequired)
			if got != tt.want {
				t.Errorf("flagHelp(%q, %q, %v)\n  got:  %s\n  want: %s", tt.fieldName, tt.doc, tt.isRequired, got, tt.want)
			}
		})
	}
}

func TestDefaultValue(t *testing.T) {
	tests := []struct {
		name      string
		typ       string
		fieldName string
		parsed    string
		want      string
	}{
		// Body always overrides — even with a parsed default.
		{name: "Body wins over parsed", typ: "string", fieldName: "Body", parsed: `"ignored"`, want: `""`},

		// Parsed default takes precedence for primitive types.
		{name: "int parsed", typ: "int64", fieldName: "Page", parsed: "1", want: "1"},
		{name: "int64 parsed", typ: "int64", fieldName: "Limit", parsed: "1000", want: "1000"},
		{name: "string parsed (already quoted)", typ: "string", fieldName: "Permission", parsed: `"View"`, want: `"View"`},
		{name: "bool parsed", typ: "bool", fieldName: "Enabled", parsed: "true", want: "true"},
		{name: "float parsed", typ: "float64", fieldName: "Threshold", parsed: "0.5", want: "0.5"},

		// Slice types ignore parsed (we cannot trust the literal form there).
		{name: "[]string ignores parsed", typ: "[]string", fieldName: "UID", parsed: "[a b]", want: "[]string{}"},
		{name: "[]int64 ignores parsed", typ: "[]int64", fieldName: "IDs", parsed: "[1 2]", want: "[]int64{}"},

		// No parsed: zero-value fallback per type.
		{name: "string zero", typ: "string", fieldName: "Q", parsed: "", want: `""`},
		{name: "int zero", typ: "int", fieldName: "N", parsed: "", want: "0"},
		{name: "int64 zero", typ: "int64", fieldName: "N", parsed: "", want: "0"},
		{name: "bool zero", typ: "bool", fieldName: "On", parsed: "", want: "false"},
		{name: "float zero", typ: "float64", fieldName: "X", parsed: "", want: "0.0"},
		{name: "[]string zero", typ: "[]string", fieldName: "S", parsed: "", want: "[]string{}"},
		{name: "[]int64 zero", typ: "[]int64", fieldName: "S", parsed: "", want: "[]int64{}"},

		// Perpage: parsed default still wins when upstream supplies it.
		{name: "Perpage from parsed", typ: "int64", fieldName: "Perpage", parsed: "1000", want: "1000"},
		// Perpage fallback for the upstream outlier whose doc only describes
		// the default in prose — see comment in defaultValue.
		{name: "Perpage prose fallback to 1000", typ: "int64", fieldName: "Perpage", parsed: "", want: "1000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultValue(tt.typ, tt.fieldName, tt.parsed)
			if got != tt.want {
				t.Errorf("defaultValue(%q, %q, %q) = %q, want %q", tt.typ, tt.fieldName, tt.parsed, got, tt.want)
			}
		})
	}
}

func TestCleanFieldDocAndDefault(t *testing.T) {
	// Build CommentGroups by parsing tiny Go snippets so we exercise the same
	// AST path the generator uses.
	parseDoc := func(t *testing.T, src string) (string, string) {
		t.Helper()
		full := "package p\n" + src + "\nvar X int\n"
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "test.go", full, parser.ParseComments)
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		if len(f.Comments) == 0 {
			t.Fatal("no comments parsed")
		}
		return cleanFieldDocAndDefault(f.Comments[0])
	}

	cases := []struct {
		name    string
		src     string
		wantDoc string
		wantDef string
	}{
		{
			name: "default with quoted string",
			src: `/* Permission.

	Set to ` + "`Edit`" + ` to return folders that the user can edit

	Default: "View"
*/`,
			wantDoc: "Set to `Edit` to return folders that the user can edit",
			wantDef: `"View"`,
		},
		{
			name: "default with bare int and Format line",
			src: `/* Limit.

	Limit the maximum number of folders to return

	Format: int64
	Default: 1000
*/`,
			wantDoc: "Limit the maximum number of folders to return",
			wantDef: "1000",
		},
		{
			name: "no Default line",
			src: `/* Tag.

	Tag is a string that you can use to filter tags.
*/`,
			wantDoc: "Tag is a string that you can use to filter tags.",
			wantDef: "",
		},
		{
			name: "Default in middle of comment still extracted",
			src: `/* Page.

	Default: 1
	Page index for starting fetching folders
*/`,
			wantDoc: "Page index for starting fetching folders",
			wantDef: "1",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			gotDoc, gotDef := parseDoc(t, tt.src)
			if gotDoc != tt.wantDoc {
				t.Errorf("doc:\n  got:  %q\n  want: %q", gotDoc, tt.wantDoc)
			}
			if gotDef != tt.wantDef {
				t.Errorf("default:\n  got:  %q\n  want: %q", gotDef, tt.wantDef)
			}
		})
	}
}
