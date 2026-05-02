package main

import (
	"encoding/json"
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

func TestBuildHelpJSON_CommonFlags(t *testing.T) {
	// Four fakes covering the gating matrix that gen.gotmpl applies for the
	// universal flags. The asserted-registered set per fake is what gen.gotmpl
	// would actually emit (computed directly from BodyField/Response state),
	// and we verify the agent contract derived from commonFlags + per-action
	// hasJSONSchema reproduces the same set.
	type fake struct {
		key                string
		act                *Action
		expectRaw          bool // always true today, kept explicit to catch contract drift
		expectDescribeBody bool
		expectDescribeResp bool
	}
	fakes := []fake{
		{
			key: "svc body-and-response",
			act: &Action{
				CmdName:    "body-and-response",
				HTTPMethod: "POST",
				BodyField:  &BodyFieldInfo{ModelType: "models.Foo", JSONSchema: "{}"},
				Response:   &ResponseInfo{TypeName: "FooOK", HasPayload: true, JSONSchema: "{}"},
			},
			expectRaw:          true,
			expectDescribeBody: true,
			expectDescribeResp: true,
		},
		{
			key: "svc body-no-schema",
			act: &Action{
				CmdName:   "body-no-schema",
				BodyField: &BodyFieldInfo{ModelType: "any", IsInterface: true},
				Response:  &ResponseInfo{TypeName: "BarOK", HasPayload: true, JSONSchema: "{}"},
			},
			expectRaw:          true,
			expectDescribeBody: false,
			expectDescribeResp: true,
		},
		{
			key: "svc response-only",
			act: &Action{
				CmdName:  "response-only",
				Response: &ResponseInfo{TypeName: "BazOK", HasPayload: true, JSONSchema: "{}"},
			},
			expectRaw:          true,
			expectDescribeBody: false,
			expectDescribeResp: true,
		},
		{
			key: "svc neither",
			act: &Action{
				CmdName:  "neither",
				Response: &ResponseInfo{NumReturns: 1},
			},
			expectRaw:          true,
			expectDescribeBody: false,
			expectDescribeResp: false,
		},
		{
			key: "svc frame-response",
			act: &Action{
				CmdName: "frame-response",
				Response: &ResponseInfo{
					TypeName:        "FrameOK",
					HasPayload:      true,
					ContainsFrame:   true,
					JSONSchema:      "{}",
					RootDescription: "Response wire format differs from the Go type; consider --raw.",
				},
			},
			expectRaw:          true,
			expectDescribeBody: false,
			expectDescribeResp: true,
		},
	}

	svc := &Service{CmdName: "svc", Actions: make([]*Action, 0, len(fakes))}
	for _, f := range fakes {
		svc.Actions = append(svc.Actions, f.act)
	}
	raw, err := buildHelpJSON([]*Service{svc})
	if err != nil {
		t.Fatalf("buildHelpJSON: %v", err)
	}
	var doc helpDoc
	if err := json.Unmarshal([]byte(raw), &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	// commonFlags presence and shape.
	if got, want := len(doc.CommonFlags), 3; got != want {
		t.Fatalf("commonFlags length: got %d, want %d", got, want)
	}
	wantNames := map[string]string{
		"raw":                          "always",
		"describe-body-jsonschema":     "body",
		"describe-response-jsonschema": "response",
	}
	seenNames := map[string]bool{}
	for _, cf := range doc.CommonFlags {
		want, ok := wantNames[cf.Name]
		if !ok {
			t.Errorf("unexpected commonFlag name %q", cf.Name)
			continue
		}
		if cf.AppliesWhen != want {
			t.Errorf("commonFlag %q appliesWhen: got %q, want %q", cf.Name, cf.AppliesWhen, want)
		}
		if cf.Type != "bool" {
			t.Errorf("commonFlag %q type: got %q, want \"bool\"", cf.Name, cf.Type)
		}
		seenNames[cf.Name] = true
	}
	for n := range wantNames {
		if !seenNames[n] {
			t.Errorf("commonFlag %q missing", n)
		}
	}
	// raw must carry the per-action default rule; the others must not.
	for _, cf := range doc.CommonFlags {
		switch cf.Name {
		case "raw":
			if cf.DefaultRule != "response.containsFrame" {
				t.Errorf("raw defaultRule: got %q, want %q", cf.DefaultRule, "response.containsFrame")
			}
		default:
			if cf.DefaultRule != "" {
				t.Errorf("%s defaultRule: got %q, want empty", cf.Name, cf.DefaultRule)
			}
		}
	}

	// Agent contract: predict the registered flag set per command from the
	// commonFlags + per-action body/response, then compare against the gate
	// expressed in gen.gotmpl (encoded above as expect*).
	predict := func(c *helpActionOutput) (raw, body, resp bool) {
		for _, cf := range doc.CommonFlags {
			switch cf.AppliesWhen {
			case "always":
				if cf.Name == "raw" {
					raw = true
				}
			case "body":
				if cf.Name == "describe-body-jsonschema" && c.Body != nil && c.Body.HasJSONSchema {
					body = true
				}
			case "response":
				if cf.Name == "describe-response-jsonschema" && c.Response != nil && c.Response.HasJSONSchema {
					resp = true
				}
			}
		}
		return
	}
	for _, f := range fakes {
		c, ok := doc.Commands[f.key]
		if !ok {
			t.Errorf("command %q missing from output", f.key)
			continue
		}
		gotRaw, gotBody, gotResp := predict(c)
		if gotRaw != f.expectRaw || gotBody != f.expectDescribeBody || gotResp != f.expectDescribeResp {
			t.Errorf("%s: predicted (raw=%v body=%v resp=%v), want (raw=%v body=%v resp=%v)",
				f.key, gotRaw, gotBody, gotResp, f.expectRaw, f.expectDescribeBody, f.expectDescribeResp)
		}
	}

	// response.note round-trips Response.RootDescription: present iff the
	// upstream signal is set (today, iff ContainsFrame is true). This is what
	// the agent contract documented above promises about the field.
	const frameNote = "Response wire format differs from the Go type; consider --raw."
	for _, f := range fakes {
		c, ok := doc.Commands[f.key]
		if !ok {
			continue
		}
		wantNote := ""
		if f.act.Response != nil {
			wantNote = f.act.Response.RootDescription
		}
		var gotNote string
		if c.Response != nil {
			gotNote = c.Response.Note
		}
		if gotNote != wantNote {
			t.Errorf("%s: response.note: got %q, want %q", f.key, gotNote, wantNote)
		}
	}
	// Sanity-check the frame-response fake actually carries the canonical text;
	// without this the empty-vs-empty case above would silently pass.
	if c, ok := doc.Commands["svc frame-response"]; !ok || c.Response == nil || c.Response.Note != frameNote {
		t.Errorf("svc frame-response: response.note must equal %q (got %+v)", frameNote, c)
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
