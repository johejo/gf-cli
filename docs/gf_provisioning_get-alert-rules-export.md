## gf provisioning get-alert-rules-export



```
gf provisioning get-alert-rules-export [flags]
```

### Options

```
      --download             Whether to initiate a download of the file or not.
      --folder-uid strings   UIDs of folders from which to export rules
      --format string        Format of the downloaded file. Supported yaml, json or hcl. Accept header can also be used, but the query parameter will take precedence. Default: "yaml"
      --group string         Name of group of rules to export. Must be specified only together with a single folder UID
  -h, --help                 help for get-alert-rules-export
      --rule-uid string      UID of alert rule to export. If specified, parameters folderUid and group must be empty.
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

* [gf provisioning](gf_provisioning.md)	 - 

