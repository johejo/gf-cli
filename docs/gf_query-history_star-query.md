## gf query-history star-query

Adds star to query in query history

### Synopsis

Adds star to query in query history as specified by the UID.

Response schema (StarQueryOK.Payload):
  result                object
  result.comment        string
  result.createdAt      number
  result.createdBy      number
  result.datasourceUid  string
  result.queries        object
  result.starred        boolean
  result.uid            string

```
gf query-history star-query [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for star-query
      --query-history-uid string       QueryHistoryUID [required]
      --raw                            Print the raw HTTP response body instead of the decoded payload
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

* [gf query-history](gf_query-history.md)	 - Query history API

