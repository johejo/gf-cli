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

// TestRawMode_PropagatesHTTPError verifies that on a non-2xx response, raw mode
//  1. returns an error (non-zero exit) rather than swallowing the failure, and
//  2. writes the server error body to stderr, keeping stdout empty so agents
//     can treat "stdout = success output" as an invariant.
func TestRawMode_PropagatesHTTPError(t *testing.T) {
	canned := `{"message":"unauthorized"}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, canned)
	}))
	t.Cleanup(srv.Close)

	bodyFile := writeTempBody(t, `{"from":"now-5m","to":"now","queries":[{"refId":"A","expr":"up"}]}`)
	stdout, stderr, execErr := runGFCapture(t, srv.URL,
		"datasources", "query-metrics-with-expressions",
		"--body", bodyFile,
		"--raw=true",
	)
	if execErr == nil {
		t.Fatalf("expected non-nil error for 401 response, got nil; stderr: %s", stderr)
	}
	if strings.Contains(stdout, "unauthorized") {
		t.Fatalf("error body must not go to stdout, got stdout: %s", stdout)
	}
	if !strings.Contains(stderr, "unauthorized") {
		t.Fatalf("expected server body on stderr, got: %s", stderr)
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
	stdout, _, execErr := runGFCapture(t, serverURL, args...)
	return stdout, execErr
}

// runGFCapture is like runGFAllowErr but also captures stderr. The CLI writes
// successful payloads to os.Stdout and error payloads/diagnostics to os.Stderr,
// so both real OS streams are replaced with pipes here.
func runGFCapture(t *testing.T, serverURL string, args ...string) (string, string, error) {
	t.Helper()
	t.Setenv("GF_HOST", serverURL)
	t.Setenv("GF_BASE_PATH", "/api")
	t.Setenv("GF_API_KEY", "")
	t.Setenv("GF_BASIC_AUTH_USERNAME", "")
	t.Setenv("GF_BASIC_AUTH_PASSWORD", "")

	oldStdout, oldStderr := os.Stdout, os.Stderr
	outR, outW, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	errR, errW, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout, os.Stderr = outW, errW
	capture := func(r *os.File) chan []byte {
		ch := make(chan []byte, 1)
		go func() {
			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			ch <- buf.Bytes()
		}()
		return ch
	}
	outDone, errDone := capture(outR), capture(errR)

	cmd := internal.RootCmd()
	cmd.SetArgs(args)
	// Route cobra's own diagnostics to os.Stderr (the captured pipe) so the
	// test sees the same stream split a real invocation produces.
	cmd.SetErr(os.Stderr)
	execErr := cmd.Execute()

	_ = outW.Close()
	_ = errW.Close()
	os.Stdout, os.Stderr = oldStdout, oldStderr
	stdout, stderr := <-outDone, <-errDone

	return string(stdout), string(stderr), execErr
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
