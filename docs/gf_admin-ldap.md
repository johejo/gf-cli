## gf admin-ldap

Admin ldap API

```
gf admin-ldap [flags]
```

### Options

```
  -h, --help   help for admin-ldap
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

* [gf](gf.md)	 - CLI for Grafana API
* [gf admin-ldap get-ldap-status](gf_admin-ldap_get-ldap-status.md)	 - Attempts to connect to all the configured LDAP servers and returns information on whenever they re available or not
* [gf admin-ldap get-sync-status](gf_admin-ldap_get-sync-status.md)	 - Returns the current state of the LDAP background sync integration
* [gf admin-ldap get-user-from-ldap](gf_admin-ldap_get-user-from-ldap.md)	 - Finds an user based on a username in LDAP this helps illustrate how would the particular user be mapped in grafana when
* [gf admin-ldap post-sync-user-with-ldap](gf_admin-ldap_post-sync-user-with-ldap.md)	 - Enables a single grafana user to be synchronized against LDAP
* [gf admin-ldap reload-ldap-cfg](gf_admin-ldap_reload-ldap-cfg.md)	 - Reloads the LDAP configuration

