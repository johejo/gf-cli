## gf access-control set-resource-permissions-for-built-in-role



```
gf access-control set-resource-permissions-for-built-in-role [flags]
```

### Options

```
      --body string            The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --built-in-role string   BuiltInRole
  -h, --help                   help for set-resource-permissions-for-built-in-role
      --resource string        Resource
      --resource-id string     ResourceID
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

* [gf access-control](gf_access-control.md)	 - 

