## gf access-control list-roles

Gets all roles

### Synopsis

Gets all roles

Gets all existing roles. The response contains all global and organization local roles, for the organization which user is signed in.

You need to have a permission with action `roles:read` and scope `roles:*`.

The `delegatable` flag reduces the set of roles to only those for which the signed-in user has permissions to assign.

```
gf access-control list-roles [flags]
```

### Options

```
      --delegatable                    Delegatable
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for list-roles
      --include-hidden                 IncludeHidden
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --target-org-id int              TargetOrgID
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

* [gf access-control](gf_access-control.md)	 - Access control API

