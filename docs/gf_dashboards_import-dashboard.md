## gf dashboards import-dashboard

Imports dashboard

### Synopsis

Imports dashboard

Body schema (ImportDashboardRequest):
  dashboard          object
  folderId           number         Deprecated: use FolderUID instead
  folderUid          string
  inputs             array<object>
  inputs[].name      string
  inputs[].pluginId  string
  inputs[].type      string
  inputs[].value     string
  overwrite          boolean
  path               string
  pluginId           string

Response schema (ImportDashboardOK.Payload):
  dashboardId       number
  description       string
  folderId          number   Deprecated: use FolderUID instead
  folderUid         string
  imported          boolean
  importedRevision  number
  importedUri       string
  importedUrl       string
  path              string
  pluginId          string
  removed           boolean
  revision          number
  slug              string
  title             string
  uid               string

```
gf dashboards import-dashboard [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for import-dashboard
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

