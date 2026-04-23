package main

import (
	"encoding/json"
	"os/exec"
	"testing"
)

func openAPIClientDir(t *testing.T) string {
	t.Helper()
	out, err := exec.Command("go", "list", "-m", "-json", "github.com/grafana/grafana-openapi-client-go").Output()
	if err != nil {
		t.Fatalf("go list module: %v", err)
	}
	var info struct {
		Dir string
	}
	if err := json.Unmarshal(out, &info); err != nil {
		t.Fatalf("parse go list output: %v", err)
	}
	if info.Dir == "" {
		t.Fatal("empty module dir from go list")
	}
	return info.Dir
}

func TestParseResponse_ContainsFrame(t *testing.T) {
	baseDir := openAPIClientDir(t)

	tests := []struct {
		name       string
		pkg        string
		resp       string
		wantFrame  bool
		wantHasPld bool
	}{
		{"datasources query", "datasources", "QueryMetricsWithExpressionsOK", true, true},
		{"dashboards query public", "dashboards", "QueryPublicDashboardOK", true, true},
		{"datasources get-datasources", "datasources", "GetDataSourcesOK", false, true},
		{"health get-health", "health", "GetHealthOK", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rf, err := ParseResponse(baseDir, tt.pkg, tt.resp)
			if err != nil {
				t.Fatalf("ParseResponse: %v", err)
			}
			if rf.HasPayload != tt.wantHasPld {
				t.Errorf("HasPayload = %v, want %v", rf.HasPayload, tt.wantHasPld)
			}
			if rf.ContainsFrame != tt.wantFrame {
				t.Errorf("ContainsFrame = %v, want %v", rf.ContainsFrame, tt.wantFrame)
			}
		})
	}
}
