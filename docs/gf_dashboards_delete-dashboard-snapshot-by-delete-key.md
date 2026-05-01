## gf dashboards delete-dashboard-snapshot-by-delete-key

Deletes snapshot by delete key

### Synopsis

Deletes snapshot by delete key

Snapshot public mode should be enabled or authentication is required.

Response schema (DeleteDashboardSnapshotByDeleteKeyOK.Payload):
  message  string

```
gf dashboards delete-dashboard-snapshot-by-delete-key [flags]
```

### Options

```
      --delete-key string              DeleteKey
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for delete-dashboard-snapshot-by-delete-key
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

* [gf dashboards](gf_dashboards.md)	 - Dashboards API

