## gf provisioning export-mute-timing

Exports a mute timing in provisioning format

```
gf provisioning export-mute-timing [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --download                       Whether to initiate a download of the file or not.
      --format string                  Format of the downloaded file. Supported yaml, json or hcl. Accept header can also be used, but the query parameter will take precedence. Default: "yaml"
  -h, --help                           help for export-mute-timing
      --name string                    Mute timing name
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

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

