package main

import "testing"

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
