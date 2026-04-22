## gf org remove-org-user-for-current-org

Deletes user in current organization

### Synopsis

Deletes user in current organization

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `org.users:remove` with scope `users:*`.

```
gf org remove-org-user-for-current-org [flags]
```

### Options

```
  -h, --help          help for remove-org-user-for-current-org
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

* [gf org](gf_org.md)	 - 

