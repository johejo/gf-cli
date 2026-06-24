## gf dashboards get-dashboard-versions-by-uid

Gets all existing versions for the dashboard using UID

### Synopsis

Gets all existing versions for the dashboard using UID

Response schema (GetDashboardVersionsByUIDOK.Payload):
  continueToken             string
  versions                  array<object>
  versions[].created        string
  versions[].createdBy      string
  versions[].dashboardId    number
  versions[].data           object
  versions[].id             number
  versions[].message        string
  versions[].parentVersion  number
  versions[].restoredFrom   number
  versions[].uid            string
  versions[].version        number

```
gf dashboards get-dashboard-versions-by-uid [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-dashboard-versions-by-uid
      --limit int                      Maximum number of results to return
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --start int                      Version to start from when returning queries
      --uid string                     UID [required]
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf dashboards](gf_dashboards.md)	 - Dashboards API

