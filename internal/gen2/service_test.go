package main

import "testing"

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
			name: "first paragraph has multiple lines",
			in: "CreateDashboardSnapshot whens creating a snapshot using the API\n" +
				"you have to provide the full dashboard payload\n\n" +
				"Snapshot public mode should be enabled or authentication is required.",
			wantShort: "Whens creating a snapshot using the API you have to provide the full dashboard payload",
			wantLong:  "Snapshot public mode should be enabled or authentication is required.",
		},
		{
			name:      "empty",
			in:        "",
			wantShort: "",
			wantLong:  "",
		},
	}

	for _, tt := range tests {
		tt := tt
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
		tt := tt
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
