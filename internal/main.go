package internal

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	gfclient "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/models"
	"github.com/spf13/cobra"
	"mvdan.cc/xurls/v2"
)

//go:embed help.gotmpl
var helpTemplate string

//go:embed help.json
var helpJSON string

var (
	rootCmd = &cobra.Command{
		Use:               "gf",
		Short:             "CLI for Grafana API",
		Long:              "Grafana API Client for command line operations",
		DisableAutoGenTag: true,
		SilenceUsage:      true,
		Args:              cobra.NoArgs,
		RunE:              runRoot,
	}
	rootCmdFlag = struct {
		host              string
		basePath          string
		apiKey            string
		basicAuthUsername string
		basicAuthPassword string
		orgID             int64
		version           bool
		debug             bool
		noColor           bool
		colors            string
		helpJSON          bool
		timeout           time.Duration
	}{}
)

func RootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.SetHelpTemplate(helpTemplate)
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.host, "host", "localhost:3000", "Grafana server host (env: GF_HOST)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.basePath, "base-path", "/api", "Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.apiKey, "api-key", "", "API Key to authenticate to grafana server (env: GF_API_KEY)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.basicAuthUsername, "basic-user-username", "", "Basic authentication username (env: GF_BASIC_AUTH_USERNAME)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.basicAuthPassword, "basic-user-password", "", "Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)")
	rootCmd.PersistentFlags().Int64Var(&rootCmdFlag.orgID, "org-id", 0, "Organization ID (env: GF_ORG_ID)")
	rootCmd.PersistentFlags().BoolVar(&rootCmdFlag.debug, "debug", false, "Enable debug logging (env: GF_DEBUG)")
	rootCmd.PersistentFlags().DurationVar(&rootCmdFlag.timeout, "timeout", 30*time.Second, "Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT)")
	// --help-json is registered as a root-only flag (not persistent) so it does
	// not appear on every subcommand's --help. The output is a discovery index
	// (no JSON Schemas); agents fetch full body/response schemas via the
	// per-subcommand --describe-body-jsonschema / --describe-response-jsonschema
	// flags.
	rootCmd.Flags().BoolVar(&rootCmdFlag.helpJSON, "help-json", false, "Print the CLI schema index (commands, flags, body/response model types) as JSON and exit. Use --describe-body-jsonschema / --describe-response-jsonschema on a subcommand for the full JSON Schema.")
}

func runRoot(cmd *cobra.Command, args []string) error {
	if rootCmdFlag.helpJSON {
		return printHelpJSON()
	}
	if len(args) == 0 {
		cmd.PrintErrln("No subcommand specified.")
		cmd.PrintErr(cmd.UsageString())
		os.Exit(1)
	}
	return nil
}

func failIfEmptyArgs(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.PrintErrln("No subcommand specified.")
		cmd.PrintErr(cmd.UsageString())
		os.Exit(1)
	}
}

func gfClient() (*gfclient.GrafanaHTTPAPI, error) {
	cfg := gfclient.DefaultTransportConfig()

	host := os.Getenv("GF_HOST")
	if host == "" {
		host = rootCmdFlag.host
	}
	scheme, host, err := parseSchemeAndHost(host)
	if err != nil {
		return nil, err
	}
	cfg = cfg.WithSchemes([]string{scheme})
	cfg = cfg.WithHost(host)

	cfg = applyEnvString("GF_BASE_PATH", rootCmdFlag.basePath, cfg.WithBasePath)
	cfg = applyEnvBool("GF_DEBUG", rootCmdFlag.debug, func(b bool) *gfclient.TransportConfig {
		cfg.Debug = b
		return cfg
	})
	cfg = applyEnvString("GF_API_KEY", rootCmdFlag.apiKey, func(v string) *gfclient.TransportConfig {
		cfg.APIKey = v
		return cfg
	})
	basicUser := os.Getenv("GF_BASIC_AUTH_USERNAME")
	if basicUser == "" {
		basicUser = rootCmdFlag.basicAuthUsername
	}
	basicPass := os.Getenv("GF_BASIC_AUTH_PASSWORD")
	if basicPass == "" {
		basicPass = rootCmdFlag.basicAuthPassword
	}
	if basicUser != "" && basicPass != "" {
		cfg.BasicAuth = url.UserPassword(basicUser, basicPass)
	}
	api := gfclient.NewHTTPClientWithConfig(nil, cfg)
	api = applyEnvInt64("GF_ORG_ID", rootCmdFlag.orgID, api.WithOrgID)
	rawResponseBody = nil
	timeout := rootCmdFlag.timeout
	if v, ok := os.LookupEnv("GF_TIMEOUT"); ok {
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("invalid GF_TIMEOUT %q: %w", v, err)
		}
		timeout = d
	}
	api = api.WithHTTPClient(&http.Client{
		Transport: rawCaptureRoundTripper{},
		Timeout:   timeout,
	})
	return api, nil
}

// rawResponseBody holds the most recently received HTTP response body. It is
// populated by rawCaptureRoundTripper on every request and consumed by
// printRawResponse when a command's --raw flag is set. Commands run
// sequentially so there is no concurrency concern.
var rawResponseBody []byte

// rawCaptureRoundTripper tees response bodies into rawResponseBody so that
// callers can emit the server's actual JSON verbatim, bypassing
// grafana-openapi-client-go models whose shapes don't match Grafana's wire
// format (notably models.Frame).
type rawCaptureRoundTripper struct {
	base http.RoundTripper
}

func (r rawCaptureRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	base := r.base
	if base == nil {
		base = http.DefaultTransport
	}
	resp, err := base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	body, rerr := io.ReadAll(resp.Body)
	resp.Body.Close()
	if rerr != nil {
		return nil, rerr
	}
	rawResponseBody = body
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return resp, nil
}

func hasRawResponse() bool {
	return len(rawResponseBody) > 0
}

// printRawResponse writes the captured raw body to stdout for a successful
// request; printErrRawResponse writes it to stderr when the request failed, so
// stdout only ever carries successful output.
func printRawResponse() error    { return writeRawResponse(os.Stdout) }
func printErrRawResponse() error { return writeRawResponse(os.Stderr) }

func writeRawResponse(w io.Writer) error {
	if len(rawResponseBody) == 0 {
		return nil
	}
	// Use json.Indent (not Unmarshal + Encode) so large integer values are
	// preserved verbatim — re-encoding would round them into float64.
	var buf bytes.Buffer
	if err := json.Indent(&buf, rawResponseBody, "", "  "); err == nil {
		buf.WriteByte('\n')
		_, err := w.Write(buf.Bytes())
		return err
	}
	_, err := w.Write(rawResponseBody)
	return err
}

func parseSchemeAndHost(s string) (string, string, error) {
	if strings.HasPrefix(s, "localhost") {
		return "http", s, nil
	}
	host := xurls.Relaxed().FindString(s)
	if host == "" {
		return "", "", fmt.Errorf("gf: no valid host found in %s", s)
	}
	if strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "http://") {
		if scheme, host, ok := strings.Cut(host, "://"); ok {
			return scheme, host, nil
		}
	}
	if _, err := netip.ParseAddrPort(host); err == nil {
		// use http for raw ip addr
		return "http", host, nil
	}
	// default secure
	return "https", host, nil
}

func applyEnvString[T any](key string, flg string, f func(string) *T) *T {
	return applyEnv(key, flg, func(s string) (string, error) { return s, nil }, f)
}

func applyEnv[T any, V any](key string, flg V, parseFn func(string) (V, error), applyFn func(V) *T) *T {
	t := applyFn(flg)
	if v, ok := os.LookupEnv(key); ok {
		vv, err := parseFn(v)
		if err != nil {
			return t
		}
		return applyFn(vv)
	}
	return t
}

func applyEnvBool[T any](key string, flg bool, f func(bool) *T) *T {
	return applyEnv(key, flg, strconv.ParseBool, f)
}

func applyEnvInt64[T any](key string, flg int64, f func(int64) *T) *T {
	return applyEnv(key, flg, func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) }, f)
}

// printPayload writes a successful payload to stdout; printErrPayload writes a
// structured API error body to stderr. Keeping stdout reserved for successful
// output lets agents rely on "stdout = success JSON, stderr = errors, exit code
// = success/failure".
func printPayload(p any) error    { return encodePayload(os.Stdout, p) }
func printErrPayload(p any) error { return encodePayload(os.Stderr, p) }

func encodePayload(w io.Writer, p any) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(p)
}

func getBodyParam(flg string, dst any) error {
	// Decide whether flg is inline JSON or a file path by its first
	// non-space byte. A leading '{' or '[' means the caller intended JSON, so
	// any parse failure is reported as a JSON error rather than being silently
	// retried as a file path (which yields a misleading "no such file" error).
	trimmed := strings.TrimSpace(flg)
	if strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[") {
		if err := json.Unmarshal([]byte(trimmed), dst); err != nil {
			return fmt.Errorf("--body is not valid JSON: %w", err)
		}
		return nil
	}
	b, err := os.ReadFile(flg)
	if err != nil {
		return fmt.Errorf("--body is neither inline JSON (it does not start with '{' or '[') nor a readable file: %w", err)
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return fmt.Errorf("--body file %s does not contain valid JSON: %w", flg, err)
	}
	return nil
}

type getPayloadError interface {
	GetPayload() *models.ErrorResponseBody
}

// describeBodyJSONSchema prints the JSON Schema for --describe-body-jsonschema
// and exits. os.Exit bypasses cobra's post-PreRun required-flag validation so
// the user does not need to also pass --body / path-param flags just to view
// the schema; it mirrors how --help and --version typically short-circuit.
func describeBodyJSONSchema(schema string) {
	fmt.Println(schema)
	os.Exit(0)
}

func describeResponseJSONSchema(schema string) {
	fmt.Println(schema)
	os.Exit(0)
}

// printHelpJSON writes the embedded helpJSON to stdout in compact form.
// The embedded string is pretty-printed for repo diff readability;
// json.Compact strips that whitespace at runtime. Invoked by --help-json
// on the root command.
func printHelpJSON() error {
	var buf bytes.Buffer
	if err := json.Compact(&buf, []byte(helpJSON)); err != nil {
		_, werr := os.Stdout.WriteString(helpJSON)
		return werr
	}
	buf.WriteByte('\n')
	_, err := os.Stdout.Write(buf.Bytes())
	return err
}
