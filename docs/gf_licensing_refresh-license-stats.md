## gf licensing refresh-license-stats

Refreshes license stats

### Synopsis

Refreshes license stats

You need to have a permission with action `licensing:read`.

```
gf licensing refresh-license-stats [flags]
```

### Options

```
  -h, --help   help for refresh-license-stats
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

* [gf licensing](gf_licensing.md)	 - Licensing API

