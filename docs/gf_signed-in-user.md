## gf signed-in-user

Signed in user API

```
gf signed-in-user [flags]
```

### Options

```
  -h, --help   help for signed-in-user
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
* [gf signed-in-user change-user-password](gf_signed-in-user_change-user-password.md)	 - Changes password
* [gf signed-in-user clear-help-flags](gf_signed-in-user_clear-help-flags.md)	 - Clears user help flag
* [gf signed-in-user get-signed-in-user](gf_signed-in-user_get-signed-in-user.md)	 - Get (current authenticated user)
* [gf signed-in-user get-signed-in-user-org-list](gf_signed-in-user_get-signed-in-user-org-list.md)	 - Organizations of the actual user
* [gf signed-in-user get-signed-in-user-team-list](gf_signed-in-user_get-signed-in-user-team-list.md)	 - Teams that the actual user is member of
* [gf signed-in-user get-user-auth-tokens](gf_signed-in-user_get-user-auth-tokens.md)	 - Auths tokens of the actual user
* [gf signed-in-user get-user-preferences](gf_signed-in-user_get-user-preferences.md)	 - Gets user preferences
* [gf signed-in-user patch-user-preferences](gf_signed-in-user_patch-user-preferences.md)	 - Patches user preferences
* [gf signed-in-user revoke-user-auth-token](gf_signed-in-user_revoke-user-auth-token.md)	 - Revokes an auth token of the actual user
* [gf signed-in-user set-help-flag](gf_signed-in-user_set-help-flag.md)	 - Sets user help flag
* [gf signed-in-user star-dashboard-by-uid](gf_signed-in-user_star-dashboard-by-uid.md)	 - Stars a dashboard
* [gf signed-in-user unstar-dashboard-by-uid](gf_signed-in-user_unstar-dashboard-by-uid.md)	 - Unstars a dashboard
* [gf signed-in-user update-signed-in-user](gf_signed-in-user_update-signed-in-user.md)	 - Updates signed in user
* [gf signed-in-user update-user-preferences](gf_signed-in-user_update-user-preferences.md)	 - Updates user preferences
* [gf signed-in-user user-set-using-org](gf_signed-in-user_user-set-using-org.md)	 - Switches user context for signed in user

