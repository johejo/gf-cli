## gf folders get-folders

Gets all folders

### Synopsis

Gets all folders

It returns all folders that the authenticated user has permission to view. If nested folders are enabled, it expects an additional query parameter with the parent folder UID and returns the immediate subfolders that the authenticated user has permission to view. If the parameter is not supplied then it returns immediate subfolders under the root that the authenticated user has permission to view.

Use: /apis/folder.grafana.app/v1/namespaces/{ns}/folders

```
gf folders get-folders [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-folders
      --limit int                      Limit the maximum number of folders to return (default 1000)
      --page int                       Page index for starting fetching folders (default 1)
      --parent-uid string              The parent folder UID
      --permission Edit                Set to Edit to return folders that the user can edit (default "View")
      --raw                            Print the raw HTTP response body instead of the decoded payload
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

* [gf folders](gf_folders.md)	 - Folders API

