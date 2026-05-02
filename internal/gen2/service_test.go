package main

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestSplitDoc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		in        string
		wantShort string
		wantLong  string
	}{
		{
			name:      "single line comment",
			in:        "GetAccessControlStatus gets status",
			wantShort: "Gets status",
			wantLong:  "",
		},
		{
			name: "multiple paragraphs",
			in: "AddUserRole adds a user role assignment\n\n" +
				"Assign a role to a specific user. For bulk updates consider Set user role assignments.\n\n" +
				"You need to have a permission with action `users.roles:add` and scope `permissions:type:delegate`.",
			wantShort: "Adds a user role assignment",
			wantLong: "Assign a role to a specific user. For bulk updates consider Set user role assignments.\n\n" +
				"You need to have a permission with action `users.roles:add` and scope `permissions:type:delegate`.",
		},
		{
			// Multi-line first paragraphs are joined and used as Short when
			// the joined text fits inside shortMaxLen. The earlier suppression
			// rule produced empty Shorts for ten subcommands; this case is
			// the canonical regression guard for that fix.
			name: "multi-line first paragraph is joined into Short",
			in: "CreateDashboardSnapshot whens creating a snapshot using the API\n" +
				"you have to provide the full dashboard payload\n\n" +
				"Snapshot public mode should be enabled or authentication is required.",
			wantShort: "Whens creating a snapshot using the API you have to provide the full dashboard payload",
			wantLong:  "Snapshot public mode should be enabled or authentication is required.",
		},
		{
			// Short single-line summaries pass through verbatim even if the
			// grammar is awkward; we can't reliably tell them apart from
			// valid summaries without brittle heuristics.
			name: "single-line short summary stays as Short",
			in: "PostDashboard creates update dashboard\n\n" +
				"Creates a new dashboard or updates an existing dashboard.",
			wantShort: "Creates update dashboard",
			wantLong:  "Creates a new dashboard or updates an existing dashboard.",
		},
		{
			// When the first paragraph contains multiple sentences and the
			// first one fits inside shortMaxLen, split there: first sentence
			// (with trailing period) becomes Short, the rest stays in Long.
			name: "first sentence within shortMaxLen splits into Short",
			in: "GetHealth apiHealthHandler will return ok if Grafana's web server is running and it\n" +
				"can access the database. If the database cannot be accessed it will return\n" +
				"http status code 503.",
			wantShort: "ApiHealthHandler will return ok if Grafana's web server is running and it can access the database.",
			wantLong:  "If the database cannot be accessed it will return http status code 503.",
		},
		{
			// Single-line first paragraphs that exceed shortMaxLen and have
			// no sentence boundary fall back to word-boundary truncation;
			// the full paragraph is preserved in Long. The truncation cuts
			// at the last whitespace within shortMaxLen ("...grafana when ")
			// and trims the trailing space.
			name: "long single-line first paragraph truncates into Short",
			in: "GetUserFromLDAP finds an user based on a username in LDAP this helps illustrate how would the particular user be mapped in grafana when synced\n\n" +
				"If you are running Grafana Enterprise and have Fine-grained access control enabled.",
			wantShort: "Finds an user based on a username in LDAP this helps illustrate how would the particular user be mapped in grafana when",
			wantLong: "Finds an user based on a username in LDAP this helps illustrate how would the particular user be mapped in grafana when synced\n\n" +
				"If you are running Grafana Enterprise and have Fine-grained access control enabled.",
		},
		{
			// A first paragraph that is exactly shortMaxLen long fits without
			// truncation and produces no first-paragraph Long entry.
			name:      "first paragraph at exactly shortMaxLen passes through",
			in:        "GetThing " + strings.Repeat("a", shortMaxLen-1),
			wantShort: "A" + strings.Repeat("a", shortMaxLen-2),
			wantLong:  "",
		},
		{
			// When ". " is present but the first sentence itself exceeds
			// shortMaxLen, splitFirstSentence rejects the boundary and we
			// fall through to word-boundary truncation. Without this case
			// the "sentence too long" branch in splitFirstSentence is
			// unexercised. Construction: "A " (2) + "a " * 69 (138) +
			// "ok. then more text." First sentence is 143 chars; the word
			// boundary closest to shortMaxLen=120 lands after "A " + 58
			// reps of "a " + "a", i.e. 119 chars.
			name:      "sentence boundary exists but first sentence exceeds shortMaxLen",
			in:        "GetThing " + strings.Repeat("a ", 70) + "ok. then more text.",
			wantShort: "A " + strings.Repeat("a ", 58) + "a",
			wantLong:  "A " + strings.Repeat("a ", 69) + "ok. then more text.",
		},
		{
			name:      "empty",
			in:        "",
			wantShort: "",
			wantLong:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := splitDoc(tt.in)
			if got.Short != tt.wantShort {
				t.Fatalf("Short = %q, want %q", got.Short, tt.wantShort)
			}
			if got.Long != tt.wantLong {
				t.Fatalf("Long = %q, want %q", got.Long, tt.wantLong)
			}
		})
	}
}

func TestBuildMethodDocMap_RecoversBlankLineSeparated(t *testing.T) {
	t.Parallel()

	src := `package fake

type Client struct{}

/*
WithBlankLine handles separated docs

Long paragraph here.
*/

func (a *Client) WithBlankLine() error { return nil }

/*
NoBlankLine handles attached docs
*/
func (a *Client) NoBlankLine() error { return nil }

func (a *Client) NoDoc() error { return nil }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "fake.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}

	m := buildMethodDocMap(f)

	if got, want := m["WithBlankLine"].Short, "Handles separated docs"; got != want {
		t.Errorf("WithBlankLine.Short = %q, want %q", got, want)
	}
	if got, want := m["WithBlankLine"].Long, "Long paragraph here."; got != want {
		t.Errorf("WithBlankLine.Long = %q, want %q", got, want)
	}
	if got, want := m["NoBlankLine"].Short, "Handles attached docs"; got != want {
		t.Errorf("NoBlankLine.Short = %q, want %q", got, want)
	}
	if _, ok := m["NoDoc"]; ok {
		t.Errorf("NoDoc should have no entry, got %+v", m["NoDoc"])
	}
}

func TestExtractOperationInfo(t *testing.T) {
	t.Parallel()

	src := `package fake

type Client struct{}
type FakeParams struct{}

func (a *Client) GetUserFromLDAPWithParams(params *FakeParams) error {
	if params == nil {
		params = &FakeParams{}
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserFromLDAP",
		Method:             "GET",
		PathPattern:        "/admin/ldap/{user_name}",
		ProducesMediaTypes: []string{"application/json"},
	}
	_ = op
	return nil
}

func (a *Client) NoOp() error { return nil }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "fake.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}

	method, path := extractOperationInfo(f, "GetUserFromLDAPWithParams")
	if method != "GET" {
		t.Errorf("method = %q, want %q", method, "GET")
	}
	if path != "/admin/ldap/{user_name}" {
		t.Errorf("path = %q, want %q", path, "/admin/ldap/{user_name}")
	}

	method, path = extractOperationInfo(f, "DoesNotExist")
	if method != "" || path != "" {
		t.Errorf("missing method should return empty, got (%q, %q)", method, path)
	}

	method, path = extractOperationInfo(f, "NoOp")
	if method != "" || path != "" {
		t.Errorf("body without ClientOperation should return empty, got (%q, %q)", method, path)
	}
}

func TestActionLongParts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		act  *Action
		want []string
	}{
		{
			name: "short only does not force long help",
			act: &Action{
				Short: "Gets a single data source by name",
			},
			want: nil,
		},
		{
			name: "long help starts with short summary",
			act: &Action{
				Short: "Gets a single data source by name",
				Long:  "If you are running Grafana Enterprise and have Fine-grained access control enabled.",
			},
			want: []string{
				"Gets a single data source by name",
				"If you are running Grafana Enterprise and have Fine-grained access control enabled.",
			},
		},
		{
			name: "long help preserves multiple long paragraphs",
			act: &Action{
				Short: "Creates a data source",
				Long:  "First long paragraph.\n\nSecond long paragraph.",
			},
			want: []string{
				"Creates a data source",
				"First long paragraph.",
				"Second long paragraph.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := actionLongParts(tt.act)
			if len(got) != len(tt.want) {
				t.Fatalf("len(actionLongParts) = %d, want %d: %#v", len(got), len(tt.want), got)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("actionLongParts[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestActionBodySchema(t *testing.T) {
	t.Parallel()

	const schema = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "CreateTeamCommand"
}`

	tests := []struct {
		name    string
		act     *Action
		wantHas bool
	}{
		{
			name:    "no body field",
			act:     &Action{Short: "Creates a team"},
			wantHas: false,
		},
		{
			name:    "body field without JSON schema",
			act:     &Action{BodyField: &BodyFieldInfo{ModelType: "models.CreateTeamCommand"}},
			wantHas: false,
		},
		{
			name: "body field with JSON schema",
			act: &Action{BodyField: &BodyFieldInfo{
				ModelType:   "models.CreateTeamCommand",
				JSONSchema:  schema,
				Annotations: "  name                     string         REQUIRED",
			}},
			wantHas: true,
		},
		{
			name: "body field without annotations",
			act: &Action{BodyField: &BodyFieldInfo{
				ModelType:  "models.CreateTeamCommand",
				JSONSchema: schema,
			}},
			wantHas: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := actionBodySchema(tt.act)
			if tt.wantHas && got == "" {
				t.Fatalf("actionBodySchema() = empty, want non-empty")
			}
			if !tt.wantHas && got != "" {
				t.Fatalf("actionBodySchema() = %q, want empty", got)
			}
			if tt.wantHas && !strings.HasPrefix(got, "Body schema (CreateTeamCommand):\n") {
				t.Fatalf("actionBodySchema() = %q, want header prefix 'Body schema (CreateTeamCommand):\\n'", got)
			}
		})
	}
}

func TestActionResponseSchema(t *testing.T) {
	t.Parallel()

	const schema = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "CreateTeamOK.Payload"
}`

	tests := []struct {
		name    string
		act     *Action
		wantHas bool
	}{
		{
			name:    "no response",
			act:     &Action{},
			wantHas: false,
		},
		{
			name:    "response without JSON schema",
			act:     &Action{Response: &ResponseInfo{TypeName: "CreateTeamOK"}},
			wantHas: false,
		},
		{
			name: "response with JSON schema",
			act: &Action{Response: &ResponseInfo{
				TypeName:    "CreateTeamOK",
				JSONSchema:  schema,
				Annotations: "  id                       integer        REQUIRED",
			}},
			wantHas: true,
		},
		{
			name: "response with frame advisory only",
			act: &Action{Response: &ResponseInfo{
				TypeName:        "CreateTeamOK",
				JSONSchema:      schema,
				RootDescription: "Response wire format differs from the Go type; consider --raw.",
			}},
			wantHas: true,
		},
		{
			name: "response without annotations or advisory",
			act: &Action{Response: &ResponseInfo{
				TypeName:   "CreateTeamOK",
				JSONSchema: schema,
			}},
			wantHas: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := actionResponseSchema(tt.act)
			if tt.wantHas && got == "" {
				t.Fatalf("actionResponseSchema() = empty, want non-empty")
			}
			if !tt.wantHas && got != "" {
				t.Fatalf("actionResponseSchema() = %q, want empty", got)
			}
			if tt.wantHas && !strings.HasPrefix(got, "Response schema (CreateTeamOK.Payload):\n") {
				t.Fatalf("actionResponseSchema() = %q, want header prefix 'Response schema (CreateTeamOK.Payload):\\n'", got)
			}
		})
	}
}

func TestActionResponseSchemaCombinedLayout(t *testing.T) {
	t.Parallel()

	act := &Action{Response: &ResponseInfo{
		TypeName:        "QueryPublicDashboardOK",
		JSONSchema:      `{"$schema":"https://json-schema.org/draft/2020-12/schema"}`,
		Annotations:     "  results                  map<string, object>",
		RootDescription: "Response wire format differs from the Go type; consider --raw.",
	}}
	want := "Response schema (QueryPublicDashboardOK.Payload):\n" +
		"  Note: Response wire format differs from the Go type; consider --raw.\n" +
		"  results                  map<string, object>"
	if got := actionResponseSchema(act); got != want {
		t.Fatalf("actionResponseSchema() =\n%s\nwant:\n%s", got, want)
	}
}
