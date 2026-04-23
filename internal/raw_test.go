package internal_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/johejo/gf-cli/internal"
)

// TestRawMode_QueryMetricsWithExpressions verifies that:
//  1. The --raw flag defaults to true for datasources query-metrics-with-expressions
//     (the response type contains models.Frame so the typed decode loses data), and
//     the command prints the server body verbatim through the JSON pretty-printer.
//  2. Passing --raw=false falls back to the old typed output, which drops the
//     wire-format data frames (surfacing as {"Fields": null}).
func TestRawMode_QueryMetricsWithExpressions(t *testing.T) {
	canned := `{"results":{"A":{"frames":[{"schema":{"refId":"A","fields":[{"name":"Time","type":"time"},{"name":"up","type":"number","labels":{"__name__":"up"}}]},"data":{"values":[[1776841240574],[1]]}}]}}}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/ds/query" || r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, canned)
	}))
	t.Cleanup(srv.Close)

	bodyFile := writeTempBody(t, `{"from":"now-5m","to":"now","queries":[{"refId":"A","expr":"up"}]}`)

	t.Run("default_raw_true", func(t *testing.T) {
		out := runGF(t, srv.URL,
			"datasources", "query-metrics-with-expressions",
			"--body", bodyFile,
		)
		if !jsonEquivalent(t, out, canned) {
			t.Fatalf("raw output does not match server body\nwant: %s\ngot: %s", canned, out)
		}
	})

	t.Run("raw_false_uses_typed_decode", func(t *testing.T) {
		out := runGF(t, srv.URL,
			"datasources", "query-metrics-with-expressions",
			"--body", bodyFile,
			"--raw=false",
		)
		if !strings.Contains(out, `"Fields": null`) {
			t.Fatalf("expected typed decode with lossy {\"Fields\": null} frames, got: %s", out)
		}
		if strings.Contains(out, `"values"`) {
			t.Fatalf("typed decode should not contain raw wire-format fields, got: %s", out)
		}
	})
}

// TestRawMode_OptInOnNonLossyCommand verifies that --raw is default-false on
// commands whose response type does not contain models.Frame, and that passing
// --raw flips output to the verbatim server body.
func TestRawMode_OptInOnNonLossyCommand(t *testing.T) {
	canned := `[{"id":1,"name":"ds","uid":"abc","extraWireField":"preserved"}]`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/datasources" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, canned)
	}))
	t.Cleanup(srv.Close)

	t.Run("default_typed", func(t *testing.T) {
		out := runGF(t, srv.URL, "datasources", "get-datasources")
		if strings.Contains(out, "extraWireField") {
			t.Fatalf("typed decode should not expose unknown wire fields, got: %s", out)
		}
	})

	t.Run("raw_flag_prints_verbatim", func(t *testing.T) {
		out := runGF(t, srv.URL, "datasources", "get-datasources", "--raw")
		if !strings.Contains(out, "extraWireField") {
			t.Fatalf("raw output should preserve unknown wire fields, got: %s", out)
		}
	})
}

// TestRawMode_PreservesLargeIntegers guards against re-encoding JSON numbers
// through float64, which would round integers above 2^53. Raw mode must emit
// large integer values (e.g. nanosecond timestamps) byte-for-byte.
func TestRawMode_PreservesLargeIntegers(t *testing.T) {
	const bigInt = "9223372036854775000" // > 2^53, safely below math.MaxInt64
	canned := `{"results":{"A":{"frames":[{"data":{"values":[[` + bigInt + `]]}}]}}}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, canned)
	}))
	t.Cleanup(srv.Close)

	bodyFile := writeTempBody(t, `{"from":"now-5m","to":"now","queries":[{"refId":"A","expr":"up"}]}`)
	out, _ := runGFAllowErr(t, srv.URL,
		"datasources", "query-metrics-with-expressions",
		"--body", bodyFile,
		"--raw=true",
	)
	if !strings.Contains(out, bigInt) {
		t.Fatalf("large integer %s not preserved verbatim in raw output:\n%s", bigInt, out)
	}
}

// TestRawMode_PropagatesHTTPError verifies that raw mode still returns an
// error (non-zero exit) when the server responds with a non-2xx status, even
// though the body has already been printed. Without this, --raw swallows
// failures and commands look successful.
func TestRawMode_PropagatesHTTPError(t *testing.T) {
	canned := `{"message":"unauthorized"}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, canned)
	}))
	t.Cleanup(srv.Close)

	bodyFile := writeTempBody(t, `{"from":"now-5m","to":"now","queries":[{"refId":"A","expr":"up"}]}`)
	out, execErr := runGFAllowErr(t, srv.URL,
		"datasources", "query-metrics-with-expressions",
		"--body", bodyFile,
		"--raw=true",
	)
	if execErr == nil {
		t.Fatalf("expected non-nil error for 401 response, got nil; output: %s", out)
	}
	if !strings.Contains(out, "unauthorized") {
		t.Fatalf("expected server body in raw output, got: %s", out)
	}
}

// runGF executes the root cobra command against the given httptest server and
// returns captured stdout. It replaces os.Stdout with a pipe because the CLI
// writes via os.Stdout (not cmd.OutOrStdout).
func runGF(t *testing.T, serverURL string, args ...string) string {
	t.Helper()
	out, execErr := runGFAllowErr(t, serverURL, args...)
	if execErr != nil {
		t.Fatalf("cmd.Execute: %v\nstdout: %s", execErr, out)
	}
	return out
}

func runGFAllowErr(t *testing.T, serverURL string, args ...string) (string, error) {
	t.Helper()
	t.Setenv("GF_HOST", serverURL)
	t.Setenv("GF_BASE_PATH", "/api")
	t.Setenv("GF_API_KEY", "")
	t.Setenv("GF_BASIC_AUTH_USERNAME", "")
	t.Setenv("GF_BASIC_AUTH_PASSWORD", "")

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w
	done := make(chan []byte, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.Bytes()
	}()

	cmd := internal.RootCmd()
	cmd.SetArgs(args)
	cmd.SetErr(io.Discard)
	execErr := cmd.Execute()

	_ = w.Close()
	os.Stdout = oldStdout
	out := <-done

	return string(out), execErr
}

func writeTempBody(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "body-*.json")
	if err != nil {
		t.Fatalf("temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("write temp body: %v", err)
	}
	_ = f.Close()
	return f.Name()
}

func jsonEquivalent(t *testing.T, a, b string) bool {
	t.Helper()
	var av, bv any
	if err := json.Unmarshal([]byte(a), &av); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &bv); err != nil {
		return false
	}
	ab, _ := json.Marshal(av)
	bb, _ := json.Marshal(bv)
	return bytes.Equal(ab, bb)
}
