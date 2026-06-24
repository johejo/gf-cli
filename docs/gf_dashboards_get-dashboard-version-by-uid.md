## gf dashboards get-dashboard-version-by-uid

Gets a specific dashboard version using UID

### Synopsis

Gets a specific dashboard version using UID

Response schema (GetDashboardVersionByUIDOK.Payload):
  created        string
  createdBy      string
  dashboardId    number
  data           object
  id             number
  message        string
  parentVersion  number
  restoredFrom   number
  uid            string
  version        number

```
gf dashboards get-dashboard-version-by-uid [flags]
```

### Options

```
      --dashboard-version-id int       DashboardVersionID [required]
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-dashboard-version-by-uid
      --raw                            Print the raw HTTP response body instead of the decoded payload
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

