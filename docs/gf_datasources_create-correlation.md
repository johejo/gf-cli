## gf datasources create-correlation

Adds correlation

### Synopsis

Adds correlation

Body schema (CreateCorrelationCommand):
{
  "config": {
    "field": string,
    "target": any,
    "transformations": [object],  // models.Transformations
    "type": string
  },
  "description": string,
  "label": string,
  "provisioned": boolean,
  "targetUID": string,
  "type": string
}
  config.field             REQUIRED
                           Field used to attach the correlation link
  config.target            REQUIRED
                           Target data query
  description              Optional description of the correlation
  label                    Optional label identifying the correlation
  provisioned              True if correlation was created with provisioning. This makes it read-only.
  targetUID                Target data source UID to which the correlation is created. required if type = query

```
gf datasources create-correlation [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block is a type reference; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for create-correlation
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

