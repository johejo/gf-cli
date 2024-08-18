package internal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/models"
	"github.com/itchyny/gojq"
	"github.com/johejo/gf-cli/internal/cli"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "gf",
		Short:             "CLI for Grafana API",
		Long:              "Grafana API Client for command line operations with shell completions",
		DisableAutoGenTag: true,
	}
	rootCmdFlag = struct {
		Host              string
		BasePath          string
		APIKey            string
		BasicAuthUsername string
		BasicAuthPassword string
		OrgID             int64
		Version           bool
		Debug             bool
		jq                string
		noColor           bool
		colors            string
	}{}
)

func RootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.Host, "host", "localhost:3000", "Grafana server host (env: GF_HOST)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.BasePath, "base-path", "/api", "Base path for server: useful when using sever behind reverse proxy (env: GF_BASE_PATH)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.APIKey, "api-key", "", "API Key to authenticate to grafana server (env: GF_API_KEY)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.BasicAuthUsername, "basic-user-username", "", "Basic authentication username (env: GF_BASIC_AUTH_PASSWORD)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.BasicAuthPassword, "basic-user-password", "", "Basic authentication password (env: GF_BASIC_AUTH_USERNAME)")
	rootCmd.PersistentFlags().Int64Var(&rootCmdFlag.OrgID, "org-id", 0, "Organization ID (env: GF_ORG_ID)")
	rootCmd.PersistentFlags().BoolVar(&rootCmdFlag.Debug, "debug", false, "Enable debug logging (env: GF_DEBUG)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.jq, "jq", ".", "Filter JSON output using a jq `expression` (env: GF_JQ)")
	rootCmd.PersistentFlags().BoolVar(&rootCmdFlag.noColor, "no-color", false, "Disable colored output (env: GF_NO_COLOR or NO_COLOR)")
}

func gfClient() *client.GrafanaHTTPAPI {
	cfg := client.DefaultTransportConfig()
	cfg = applyEnvString(cfg, "GF_HOST", rootCmdFlag.Host, cfg.WithHost)
	cfg = applyEnvString(cfg, "GF_BASE_PATH", rootCmdFlag.BasePath, cfg.WithBasePath)
	cfg = applyEnvBool(cfg, "GF_DEBUG", rootCmdFlag.Debug, func(b bool) *client.TransportConfig {
		cfg.Debug = b
		return cfg
	})
	cfg = applyEnvString(cfg, "GF_API_KEY", rootCmdFlag.APIKey, func(v string) *client.TransportConfig {
		cfg.APIKey = v
		return cfg
	})
	u := os.Getenv("GF_BASIC_AUTH_USERNAME")
	if u == "" {
		u = rootCmdFlag.BasicAuthUsername
	}
	p := os.Getenv("GF_BASIC_AUTH_PASSWORD")
	if p == "" {
		p = rootCmdFlag.BasicAuthPassword
	}
	if u != "" && p != "" {
		cfg.BasicAuth = url.UserPassword(u, p)
	}
	api := client.NewHTTPClientWithConfig(nil, cfg)
	api = applyEnvInt64(api, "GF_ORG_ID", rootCmdFlag.OrgID, api.WithOrgID)
	return api
}

func applyEnvString[T any](t *T, key string, flg string, f func(string) *T) *T {
	return applyEnv(t, key, flg, func(s string) (string, error) { return s, nil }, f)
}

func applyEnv[T any, V any](t *T, key string, flg V, parseFn func(string) (V, error), applyFn func(V) *T) *T {
	t = applyFn(flg)
	if v, ok := os.LookupEnv(key); ok {
		vv, err := parseFn(v)
		if err != nil {
			return t
		}
		t = applyFn(vv)
		return t
	}
	return t
}

func applyEnvBool[T any](t *T, key string, flg bool, f func(bool) *T) *T {
	return applyEnv(t, key, flg, strconv.ParseBool, f)
}

func applyEnvInt64[T any](t *T, key string, flg int64, f func(int64) *T) *T {
	return applyEnv(t, key, flg, func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) }, f)
}

func printPayload(p any) error {
	filter := rootCmdFlag.jq
	if v, ok := os.LookupEnv("GF_JQ"); ok {
		filter = v
	}
	if filter == "" {
		filter = "."
	}
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	p, err = jq(rootCmdFlag.jq, v)
	if err != nil {
		return err
	}
	noColor := rootCmdFlag.noColor || os.Getenv("GF_NO_COLOR") != "" || os.Getenv("NO_COLOR") != ""
	if !noColor && isatty.IsTerminal(os.Stdout.Fd()) {
		m := cli.NewMarshaler(false, 2)
		return m.Marshal(p, os.Stdout)
	}
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	return e.Encode(p)
}

func getBodyParam(flg string, dst any) error {
	b := []byte(flg)
	if json.Valid(b) {
		if err := json.Unmarshal(b, dst); err != nil {
			return err
		}
		return nil
	}
	b, err := os.ReadFile(flg)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}
	return nil
}

type getPayloadError interface {
	GetPayload() *models.ErrorResponseBody
}

func jq(s string, payload any) (any, error) {
	q, err := gojq.Parse(s)
	if err != nil {
		return nil, err
	}
	iter := q.Run(payload)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
				break
			}
			return v, err
		}
		return v, nil
	}
	return nil, fmt.Errorf("gf: jq filter error filter=%s, payload", s, payload)
}
