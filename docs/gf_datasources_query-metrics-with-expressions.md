## gf datasources query-metrics-with-expressions

Data source query metrics with expressions

### Synopsis

Data source query metrics with expressions

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:query`.

Body schema (MetricRequest):
  debug    boolean
  from     string         REQUIRED
                          From Start time in epoch timestamps in milliseconds or relative using Grafana time units.
  queries  array<object>  REQUIRED
                          queries.refId – Specifies an identifier of the query. Is optional and default to “A”.
                          queries.datasourceId – Specifies the data source to be queried. Each query in the request must have an unique datasourceId.
                          queries.maxDataPoints - Species maximum amount of data points that dashboard panel can render. Is optional and default to 100.
                          queries.intervalMs - Specifies the time interval in milliseconds of time series. Is optional and defaults to 1000.
  to       string         REQUIRED
                          To End time in epoch timestamps in milliseconds or relative using Grafana time units.

Response schema (QueryMetricsWithExpressionsOK.Payload):
  Note: Response wire format differs from the Go type; consider --raw.
  results  map<string, object>

```
gf datasources query-metrics-with-expressions [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for query-metrics-with-expressions
      --raw                            Print the raw HTTP response body instead of the decoded payload (default true)
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

* [gf datasources](gf_datasources.md)	 - Datasources API

