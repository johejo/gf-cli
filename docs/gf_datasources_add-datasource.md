## gf datasources add-datasource

Creates a data source

### Synopsis

Creates a data source

By defining `password` and `basicAuthPassword` under secureJsonData property Grafana encrypts them securely as an encrypted blob in the database. The response then lists the encrypted fields under secureJsonFields.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:create`

Body schema (AddDataSourceCommand):
  access           string
  basicAuth        boolean
  basicAuthUser    string
  database         string
  isDefault        boolean
  jsonData         object
  name             string
  secureJsonData   map<string, string>
  type             string
  uid              string
  url              string
  user             string
  withCredentials  boolean

Response schema (AddDataSourceOK.Payload):
  datasource                   object                REQUIRED
  datasource.access            string
  datasource.accessControl     map<string, boolean>
  datasource.basicAuth         boolean
  datasource.basicAuthUser     string
  datasource.database          string
  datasource.id                number
  datasource.isDefault         boolean
  datasource.jsonData          object
  datasource.name              string
  datasource.orgId             number
  datasource.readOnly          boolean
  datasource.secureJsonFields  map<string, boolean>
  datasource.type              string
  datasource.typeLogoUrl       string
  datasource.uid               string
  datasource.url               string
  datasource.user              string
  datasource.version           number
  datasource.withCredentials   boolean
  id                           number                REQUIRED
                                                     ID Identifier of the new data source.
  message                      string                REQUIRED
                                                     Message Message of the deleted dashboard.
  name                         string                REQUIRED
                                                     Name of the new data source.

```
gf datasources add-datasource [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for add-datasource
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

* [gf datasources](gf_datasources.md)	 - Datasources API

