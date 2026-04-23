## gf access-control remove-team-role

Removes team role

### Synopsis

Removes team role

You need to have a permission with action `teams.roles:remove` and scope `permissions:type:delegate`.

```
gf access-control remove-team-role [flags]
```

### Options

```
  -h, --help              help for remove-team-role
      --raw               Print the raw HTTP response body instead of the decoded payload
      --role-uid string   RoleUID
      --team-id int       TeamID
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

* [gf access-control](gf_access-control.md)	 - Access control API

