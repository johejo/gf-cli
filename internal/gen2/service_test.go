package main

import (
	"go/parser"
	"go/token"
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
			// A multi-line first paragraph means the upstream OpenAPI summary
			// almost certainly leaked a description into the summary slot. We
			// suppress Short and promote the paragraph into Long.
			name: "first paragraph has multiple lines suppresses Short",
			in: "CreateDashboardSnapshot whens creating a snapshot using the API\n" +
				"you have to provide the full dashboard payload\n\n" +
				"Snapshot public mode should be enabled or authentication is required.",
			wantShort: "",
			wantLong: "Whens creating a snapshot using the API you have to provide the full dashboard payload\n\n" +
				"Snapshot public mode should be enabled or authentication is required.",
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
			// A single line that exceeds shortMaxLen is treated like the
			// multi-line case: suppress Short, promote into Long.
			name: "very long single-line first paragraph suppresses Short",
			in: "CreateDashboardSnapshot whens creating a snapshot using the API you have to provide the full dashboard payload including the snapshot data this endpoint is designed for the grafana UI\n\n" +
				"Snapshot public mode should be enabled or authentication is required.",
			wantShort: "",
			wantLong: "Whens creating a snapshot using the API you have to provide the full dashboard payload including the snapshot data this endpoint is designed for the grafana UI\n\n" +
				"Snapshot public mode should be enabled or authentication is required.",
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

	schema := &BodySchemaInfo{
		TypeName: "CreateTeamCommand",
		Fields: []*ModelField{
			{JSONName: "name", GoType: "string", JSONType: "string"},
		},
	}

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
			name:    "body field without schema",
			act:     &Action{BodyField: &BodyFieldInfo{}},
			wantHas: false,
		},
		{
			name:    "body field with schema",
			act:     &Action{BodyField: &BodyFieldInfo{Schema: schema}},
			wantHas: true,
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
		})
	}
}
