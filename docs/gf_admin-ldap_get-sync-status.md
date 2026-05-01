## gf admin-ldap get-sync-status

Returns the current state of the LDAP background sync integration

### Synopsis

Returns the current state of the LDAP background sync integration

You need to have a permission with action `ldap.status:read`.

Response schema (GetSyncStatusOK.Payload):
  enabled                       boolean
  nextSync                      string
  prevSync                      object
  prevSync.Elapsed              number
  prevSync.FailedUsers          array<object>
  prevSync.FailedUsers[].Error  string
  prevSync.FailedUsers[].Login  string
  prevSync.MissingUserIds       array<number>
  prevSync.Started              string
  prevSync.UpdatedUserIds       array<number>
  schedule                      string

```
gf admin-ldap get-sync-status [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-sync-status
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

* [gf admin-ldap](gf_admin-ldap.md)	 - Admin ldap API

