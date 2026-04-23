## gf admin-ldap get-ldap-status

Attempts to connect to all the configured LDAP servers and returns information on whenever they re available or not

### Synopsis

Attempts to connect to all the configured LDAP servers and returns information on whenever they re available or not

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `ldap.status:read`.

```
gf admin-ldap get-ldap-status [flags]
```

### Options

```
  -h, --help   help for get-ldap-status
      --raw    Print the raw HTTP response body instead of the decoded payload
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

