## gf datasources get-datasource-by-name

Gets a single data source by name this function will be removed in the future

### Synopsis

Gets a single data source by name this function will be removed in the future

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scopes: `datasources:*`, `datasources:name:*` and `datasources:name:test_datasource` (single data source).

Response schema (GetDataSourceByNameOK.Payload):
  access            string
  accessControl     map<string, boolean>
  basicAuth         boolean
  basicAuthUser     string
  database          string
  id                number
  isDefault         boolean
  jsonData          object
  name              string
  orgId             number
  readOnly          boolean
  secureJsonFields  map<string, boolean>
  type              string
  typeLogoUrl       string
  uid               string
  url               string
  user              string
  version           number
  withCredentials   boolean

```
gf datasources get-datasource-by-name [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-datasource-by-name
      --name string                    Name
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

* [gf datasources](gf_datasources.md)	 - Datasources API

