## gf users update-user



```
gf users update-user [flags]
```

### Options

```
      --body string   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help          help for update-user
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
      --jq expression                Filter JSON output using a jq expression (env: GF_JQ) (default ".")
      --no-color                     Disable colored output (env: GF_NO_COLOR or NO_COLOR)
      --org-id int                   Organization ID (env: GF_ORG_ID)
```

### SEE ALSO

* [gf users](gf_users.md)	 - 

