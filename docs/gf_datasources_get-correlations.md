## gf datasources get-correlations



```
gf datasources get-correlations [flags]
```

### Options

```
  -h, --help                 help for get-correlations
      --limit int            Limit the maximum number of correlations to return per page Default: 100
      --page int             Page index for starting fetching correlations Default: 1
      --source-uid strings   Source datasource UID filter to be applied to correlations
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

* [gf datasources](gf_datasources.md)	 - Datasources API

