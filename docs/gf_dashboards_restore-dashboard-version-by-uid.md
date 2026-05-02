## gf dashboards restore-dashboard-version-by-uid

Restores a dashboard to a given dashboard version using UID

### Synopsis

Restores a dashboard to a given dashboard version using UID

This API will be removed when /apis/dashboards.grafana.app/v1 is released. You can restore a dashboard by reading it from history, then creating it again.

Body schema (RestoreDashboardVersionCommand):
  version  number

Response schema (RestoreDashboardVersionByUIDOK.Payload):
  folderUid  string  FolderUID The unique identifier (uid) of the folder the dashboard belongs to.
  id         number  REQUIRED
                     ID The unique identifier (id) of the created/updated dashboard.
  status     string  REQUIRED
                     Status status of the response.
  title      string  REQUIRED
                     Slug The slug of the dashboard.
  uid        string  REQUIRED
                     UID The unique identifier (uid) of the created/updated dashboard.
  url        string  REQUIRED
                     URL The relative URL for accessing the created/updated dashboard.
  version    number  REQUIRED
                     Version The version of the dashboard.

```
gf dashboards restore-dashboard-version-by-uid [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for restore-dashboard-version-by-uid
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --uid string                     UID
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

