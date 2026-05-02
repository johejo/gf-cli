## gf datasources update-correlation

Updates a correlation

### Synopsis

Updates a correlation

Body schema (UpdateCorrelationCommand):
  config                               object
  config.field                         string               Field used to attach the correlation link
  config.target                        map<string, object>  Target data query
  config.transformations               array<object>        Source data transformations
  config.transformations[].expression  string
  config.transformations[].field       string
  config.transformations[].mapValue    string
  config.transformations[].type        string               enum: regex | logfmt
  description                          string               Optional description of the correlation
  label                                string               Optional label identifying the correlation
  type                                 string

Response schema (UpdateCorrelationOK.Payload):
  message                                     string
  result                                      object
  result.config                               object
  result.config.field                         string               REQUIRED
                                                                   Field used to attach the correlation link
  result.config.target                        map<string, object>  REQUIRED
                                                                   Target data query
  result.config.transformations               array<object>
  result.config.transformations[].expression  string
  result.config.transformations[].field       string
  result.config.transformations[].mapValue    string
  result.config.transformations[].type        string               enum: regex | logfmt
  result.config.type                          string
  result.description                          string               Description of the correlation
  result.label                                string               Label identifying the correlation
  result.orgId                                number               OrgID of the data source the correlation originates from
  result.provisioned                          boolean              Provisioned True if the correlation was created during provisioning
  result.sourceUID                            string               UID of the data source the correlation originates from
  result.targetUID                            string               UID of the data source the correlation points to
  result.type                                 string
  result.uid                                  string               Unique identifier of the correlation

```
gf datasources update-correlation [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --correlation-uid string         CorrelationUID
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-correlation
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --source-uid string              SourceUID
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

