## gf devices search-devices

Lists all devices within the last 30 days

### Synopsis

Lists all devices within the last 30 days

Response schema (SearchDevicesOK.Payload):
  devices               array<object>
  devices[].clientIp    string
  devices[].createdAt   string
  devices[].deviceId    string
  devices[].lastSeenAt  string
  devices[].updatedAt   string
  devices[].userAgent   string
  page                  number
  perPage               number
  totalCount            number

```
gf devices search-devices [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for search-devices
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

* [gf devices](gf_devices.md)	 - Devices API

