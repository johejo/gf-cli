## gf dashboards create-dashboard-snapshot

Whens creating a snapshot using the API you have to provide the full dashboard payload including the snapshot data this endpoint is designed for the grafana UI

### Synopsis

Whens creating a snapshot using the API you have to provide the full dashboard payload including the snapshot data this endpoint is designed for the grafana UI

Snapshot public mode should be enabled or authentication is required.

Body schema (CreateDashboardSnapshotCommand):
{
  "apiVersion": string,
  "dashboard": any,
  "deleteKey": string,
  "expires": number,
  "external": boolean,
  "key": string,
  "kind": string,
  "name": string
}
  dashboard                required

```
gf dashboards create-dashboard-snapshot [flags]
```

### Options

```
      --body string                The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema   Print the JSON Schema of the request body and exit without calling the API
  -h, --help                       help for create-dashboard-snapshot
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

