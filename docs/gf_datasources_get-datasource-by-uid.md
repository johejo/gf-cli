## gf datasources get-datasource-by-uid

Gets a single data source by UID

### Synopsis

Gets a single data source by UID

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scopes: `datasources:*`, `datasources:uid:*` and `datasources:uid:kLtEtcRGk` (single data source).

```
gf datasources get-datasource-by-uid [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-datasource-by-uid
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

* [gf datasources](gf_datasources.md)	 - Datasources API

