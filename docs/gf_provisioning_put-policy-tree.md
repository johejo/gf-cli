## gf provisioning put-policy-tree

Sets the notification policy tree

### Synopsis

Sets the notification policy tree

Body schema (Route):
  active_time_intervals  array<string>
  continue               boolean
  group_by               array<string>
  group_interval         string
  group_wait             string
  match                  map<string, string>   Deprecated. Remove before v1.0 release.
  match_re               map<string, string>
  matchers               array<object>
  matchers[].isEqual     boolean
  matchers[].isRegex     boolean               REQUIRED
  matchers[].name        string                REQUIRED
  matchers[].value       string                REQUIRED
  mute_time_intervals    array<string>
  object_matchers        array<array<string>>
  provenance             string
  receiver               string
  repeat_interval        string
  routes                 array<object>

```
gf provisioning put-policy-tree [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for put-policy-tree
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --x-disable-provenance string    XDisableProvenance
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

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

