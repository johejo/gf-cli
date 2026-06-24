## gf enterprise get-datasource-cache-config

Get cache config for a single data source

### Synopsis

Get cache config for a single data source

Response schema (GetDataSourceCacheConfigOK.Payload):
  created         string
  dataSourceID    number   Fields that can be set by the API caller - read/write
  dataSourceUID   string
  defaultTTLMs    number   These are returned by the HTTP API, but are managed internally - read-only
                           Note: 'created' and 'updated' are special properties managed automatically by xorm, but we are setting them manually
  enabled         boolean
  message         string
  ttlQueriesMs    number   TTL MS, or "time to live", is how long a cached item will stay in the cache before it is removed (in milliseconds)
  ttlResourcesMs  number
  updated         string
  useDefaultTTL   boolean  If UseDefaultTTL is enabled, then the TTLQueriesMS and TTLResourcesMS in this object is always sent as the default TTL located in grafana.ini

```
gf enterprise get-datasource-cache-config [flags]
```

### Options

```
      --data-source-type string        DataSourceType
      --data-source-uid string         DataSourceUID [required]
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-datasource-cache-config
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

* [gf enterprise](gf_enterprise.md)	 - Enterprise API

