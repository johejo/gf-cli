## gf provisioning get-mute-timing

Gets a mute timing

### Synopsis

Gets a mute timing

Response schema (GetMuteTimingOK.Payload):
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
gf provisioning get-mute-timing [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-mute-timing
      --name string                    Mute timing name [required]
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

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

