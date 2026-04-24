## gf signed-in-user unstar-dashboard-by-uid

Unstars a dashboard

### Synopsis

Unstars a dashboard

Deletes the starring of the given Dashboard for the actual user.

```
gf signed-in-user unstar-dashboard-by-uid [flags]
```

### Options

```
      --dashboard-uid string           DashboardUID
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for unstar-dashboard-by-uid
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

* [gf signed-in-user](gf_signed-in-user.md)	 - Signed in user API

