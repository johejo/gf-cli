## gf admin-users admin-logout-user

Logouts user revokes all auth tokens devices for the user user of issued auth tokens devices will no longer be logged in and will be required to authenticate again upon next activity

### Synopsis

Logouts user revokes all auth tokens devices for the user user of issued auth tokens devices will no longer be logged in and will be required to authenticate again upon next activity

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.logout` and scope `global.users:*`.

```
gf admin-users admin-logout-user [flags]
```

### Options

```
  -h, --help          help for admin-logout-user
      --raw           Print the raw HTTP response body instead of the decoded payload
      --user-id int   UserID
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

* [gf admin-users](gf_admin-users.md)	 - Admin users API

