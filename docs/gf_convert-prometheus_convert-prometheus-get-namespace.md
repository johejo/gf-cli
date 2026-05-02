## gf convert-prometheus convert-prometheus-get-namespace

Gets grafana managed alert rules that were imported from prometheus compatible sources for a specified namespace folder

### Synopsis

Gets grafana managed alert rules that were imported from prometheus compatible sources for a specified namespace folder

Response schema (ConvertPrometheusGetNamespaceOK.Payload):
  Body  map<string, array<object>>  in: body

```
gf convert-prometheus convert-prometheus-get-namespace [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for convert-prometheus-get-namespace
      --namespace-title string         NamespaceTitle [required]
      --raw                            Print the raw HTTP response body instead of the decoded payload
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using sever behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_USERNAME)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_PASSWORD)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
```

### SEE ALSO

* [gf convert-prometheus](gf_convert-prometheus.md)	 - Convert prometheus API

