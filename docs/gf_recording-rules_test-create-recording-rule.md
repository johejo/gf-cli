## gf recording-rules test-create-recording-rule

Tests a recording rule

### Synopsis

Tests a recording rule

Body schema (RecordingRuleJSON):
  active                boolean
  count                 boolean
  description           string
  dest_data_source_uid  string
  id                    string
  interval              number
  name                  string
  prom_name             string
  queries               array<map<string, object>>
  range                 number
  target_ref_id         string

Response schema (TestCreateRecordingRuleOK.Payload):
  message  string

```
gf recording-rules test-create-recording-rule [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for test-create-recording-rule
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

* [gf recording-rules](gf_recording-rules.md)	 - Recording rules API

