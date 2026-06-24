## gf licensing delete-license-token

Removes license from database

### Synopsis

Removes license from database

Removes the license stored in the Grafana database. Available in Grafana Enterprise v7.4+.

You need to have a permission with action `licensing:delete`.

Body schema (DeleteTokenCommand):
  instance  string

Response schema (DeleteLicenseTokenAccepted.Payload):
  error    string  Error An optional detailed description of the actual error. Only included if running in developer mode.
  message  string  REQUIRED
                   a human readable version of the error
  status   string  Status An optional status to denote the cause of the error.
                   For example, a 412 Precondition Failed error may include additional information of why that error happened.

```
gf licensing delete-license-token [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for delete-license-token
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

* [gf licensing](gf_licensing.md)	 - Licensing API

