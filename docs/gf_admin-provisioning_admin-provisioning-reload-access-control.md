## gf admin-provisioning admin-provisioning-reload-access-control

Yous need to have a permission with action provisioning reload with scope provisioners accesscontrol

### Synopsis

Yous need to have a permission with action provisioning reload with scope provisioners accesscontrol

Response schema (AdminProvisioningReloadAccessControlAccepted.Payload):
  error    string  Error An optional detailed description of the actual error. Only included if running in developer mode.
  message  string  REQUIRED
                   a human readable version of the error
  status   string  Status An optional status to denote the cause of the error.
                   For example, a 412 Precondition Failed error may include additional information of why that error happened.

```
gf admin-provisioning admin-provisioning-reload-access-control [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for admin-provisioning-reload-access-control
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

* [gf admin-provisioning](gf_admin-provisioning.md)	 - Admin provisioning API

