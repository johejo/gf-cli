## gf access-control get-access-control-status

Gets status

### Synopsis

Gets status

Returns an indicator to check if fine-grained access control is enabled or not.

You need to have a permission with action `status:accesscontrol` and scope `services:accesscontrol`.

```
gf access-control get-access-control-status [flags]
```

### Options

```
  -h, --help   help for get-access-control-status
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

* [gf access-control](gf_access-control.md)	 - Access control API

