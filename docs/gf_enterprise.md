## gf enterprise

Enterprise API

```
gf enterprise [flags]
```

### Options

```
  -h, --help   help for enterprise
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
* [gf enterprise clean-datasource-cache](gf_enterprise_clean-datasource-cache.md)	 - Clean cache for a single data source
* [gf enterprise disable-datasource-cache](gf_enterprise_disable-datasource-cache.md)	 - Disable cache for a single data source
* [gf enterprise enable-datasource-cache](gf_enterprise_enable-datasource-cache.md)	 - Enable cache for a single data source
* [gf enterprise get-datasource-cache-config](gf_enterprise_get-datasource-cache-config.md)	 - Get cache config for a single data source
* [gf enterprise get-team-lbac-rules-api](gf_enterprise_get-team-lbac-rules-api.md)	 - Retrieves l b a c rules for a team
* [gf enterprise search-result](gf_enterprise_search-result.md)	 - Debugs permissions
* [gf enterprise set-datasource-cache-config](gf_enterprise_set-datasource-cache-config.md)	 - Set cache config for a single data source
* [gf enterprise update-team-lbac-rules-api](gf_enterprise_update-team-lbac-rules-api.md)	 - Updates l b a c rules for a team

