## gf admin-ldap get-user-from-ldap

Finds an user based on a username in LDAP this helps illustrate how would the particular user be mapped in grafana when

### Synopsis

Finds an user based on a username in LDAP this helps illustrate how would the particular user be mapped in grafana when synced

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `ldap.user:read`.

Response schema (GetUserFromLDAPOK.Payload):
  message  string

```
gf admin-ldap get-user-from-ldap [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-user-from-ldap
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --user-name string               UserName [required]
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

* [gf admin-ldap](gf_admin-ldap.md)	 - Admin ldap API

