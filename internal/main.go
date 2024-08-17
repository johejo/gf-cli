package internal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/grafana/grafana-openapi-client-go/client"
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
		OrgID             string
		Version           bool
	}{}
)

func RootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.Host, "host", "localhost:3000", "Grafana server host (env: GF_HOST)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.BasePath, "base-path", "", "Base path for server: useful when using sever behind reverse proxy (env: GF_BASE_PATH)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.APIKey, "api-key", "", "API Key to authenticate grafana server (env: GF_API_KEY)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.BasicAuthUsername, "basic-user-username", "", "Basic authentication username (env: GF_BASIC_AUTH_PASSWORD)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.BasicAuthPassword, "basic-user-password", "", "Basic authentication password (env: GF_BASIC_AUTH_USERNAME)")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlag.OrgID, "org-id", "", "Organization ID (env: GF_ORG_ID)")
}

func gfClient() *client.GrafanaHTTPAPI {
	cfg := client.DefaultTransportConfig()
	cfg = applyEnv(cfg, "GF_HOST", rootCmdFlag.Host, cfg.WithHost)
	cfg = applyEnv(cfg, "GF_BASE_PATH", rootCmdFlag.BasePath, cfg.WithBasePath)
	cfg = applyEnv(cfg, "GF_API_KEY", rootCmdFlag.APIKey, func(v string) *client.TransportConfig {
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

func applyEnv[T any](t *T, key string, flg string, f func(string) *T) *T {
	if v, ok := os.LookupEnv(key); ok {
		t = f(v)
	}
	if flg != "" {
		t = f(flg)
	}
	return t
}

func applyEnvInt64[T any](t *T, key string, flg string, f func(int64) *T) *T {
	return applyEnv(t, key, flg, func(s string) *T {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return t
		}
		return f(i)
	})
}

func printPayload(p any) error {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
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
