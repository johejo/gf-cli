## gf provisioning put-mute-timing

Replaces an existing mute timing

### Synopsis

Replaces an existing mute timing

Body schema (MuteTimeInterval):
  name                                 string
  time_intervals                       array<object>
  time_intervals[].days_of_month       array<string>
  time_intervals[].location            string
  time_intervals[].months              array<string>
  time_intervals[].times               array<object>
  time_intervals[].times[].end_time    string
  time_intervals[].times[].start_time  string
  time_intervals[].weekdays            array<string>
  time_intervals[].years               array<string>

Response schema (PutMuteTimingAccepted.Payload):
  name                                 string
  time_intervals                       array<object>
  time_intervals[].days_of_month       array<string>
  time_intervals[].location            string
  time_intervals[].months              array<string>
  time_intervals[].times               array<object>
  time_intervals[].times[].end_time    string
  time_intervals[].times[].start_time  string
  time_intervals[].weekdays            array<string>
  time_intervals[].years               array<string>

```
gf provisioning put-mute-timing [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for put-mute-timing
      --name string                    Mute timing name
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

