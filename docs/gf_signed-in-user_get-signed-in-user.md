## gf signed-in-user get-signed-in-user

Get (current authenticated user)

### Synopsis

Get (current authenticated user)

Response schema (GetSignedInUserOK.Payload):
  accessControl                   map<string, boolean>
  authLabels                      array<string>
  avatarUrl                       string
  createdAt                       string
  email                           string
  id                              number
  isDisabled                      boolean
  isExternal                      boolean
  isExternallySynced              boolean
  isGrafanaAdmin                  boolean
  isGrafanaAdminExternallySynced  boolean
  isProvisioned                   boolean
  login                           string
  name                            string
  orgId                           number
  theme                           string
  uid                             string
  updatedAt                       string

```
gf signed-in-user get-signed-in-user [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-signed-in-user
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

* [gf signed-in-user](gf_signed-in-user.md)	 - Signed in user API

