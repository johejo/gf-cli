## gf sso-settings update-provider-settings

Updates s s o settings

### Synopsis

Updates s s o settings

Inserts or updates the SSO Settings for a provider.

You need to have a permission with action `settings:write` and scope `settings:auth.<provider>:*`.

Body schema (UpdateProviderSettingsParamsBody):
  id        string
  provider  string
  settings  map<string, object>

Response schema (UpdateProviderSettingsNoContent.Payload):
  message  string

```
gf sso-settings update-provider-settings [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-provider-settings
      --key string                     Key
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

* [gf sso-settings](gf_sso-settings.md)	 - Sso settings API

