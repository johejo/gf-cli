## gf admin-users

Admin users API

```
gf admin-users [flags]
```

### Options

```
  -h, --help   help for admin-users
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

* [gf](gf.md)	 - CLI for Grafana API
* [gf admin-users admin-create-user](gf_admin-users_admin-create-user.md)	 - Creates new user
* [gf admin-users admin-delete-user](gf_admin-users_admin-delete-user.md)	 - Deletes global user
* [gf admin-users admin-disable-user](gf_admin-users_admin-disable-user.md)	 - Disables user
* [gf admin-users admin-enable-user](gf_admin-users_admin-enable-user.md)	 - Enables user
* [gf admin-users admin-get-user-auth-tokens](gf_admin-users_admin-get-user-auth-tokens.md)	 - Returns a list of all auth tokens devices that the user currently have logged in from
* [gf admin-users admin-logout-user](gf_admin-users_admin-logout-user.md)	 - Logouts user revokes all auth tokens devices for the user user of issued auth tokens devices will no longer be logged in and will be required to authenticate again upon next activity
* [gf admin-users admin-revoke-user-auth-token](gf_admin-users_admin-revoke-user-auth-token.md)	 - Revokes auth token for user
* [gf admin-users admin-update-user-password](gf_admin-users_admin-update-user-password.md)	 - Sets password for user
* [gf admin-users admin-update-user-permissions](gf_admin-users_admin-update-user-permissions.md)	 - Sets permissions for user

