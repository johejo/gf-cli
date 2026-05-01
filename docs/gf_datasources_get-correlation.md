## gf datasources get-correlation

Gets a correlation

### Synopsis

Gets a correlation

Response schema (GetCorrelationOK.Payload):
  config                               object
  config.field                         string         REQUIRED
                                                      Field used to attach the correlation link
  config.target                        object         REQUIRED
                                                      Target data query
  config.transformations               array<object>
  config.transformations[].expression  string
  config.transformations[].field       string
  config.transformations[].mapValue    string
  config.transformations[].type        string         enum: regex | logfmt
  config.type                          string
  description                          string         Description of the correlation
  label                                string         Label identifying the correlation
  orgId                                number         OrgID of the data source the correlation originates from
  provisioned                          boolean        Provisioned True if the correlation was created during provisioning
  sourceUID                            string         UID of the data source the correlation originates from
  targetUID                            string         UID of the data source the correlation points to
  type                                 string
  uid                                  string         Unique identifier of the correlation

```
gf datasources get-correlation [flags]
```

### Options

```
      --correlation-uid string         CorrelationUID
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-correlation
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

