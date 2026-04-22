## gf dashboards post-dashboard

Creates update dashboard

### Synopsis

Creates update dashboard

Creates a new dashboard or updates an existing dashboard. Note: This endpoint is not intended for creating folders, use `POST /api/folders` for that.

Body schema (SaveDashboardCommand):
{
  "UpdatedAt": string,
  "dashboard": any,
  "folderId": number,
  "folderUid": string,
  "isFolder": boolean,
  "message": string,
  "overwrite": boolean,
  "userId": number
}

```
gf dashboards post-dashboard [flags]
```

### Options

```
      --body string                The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema   Print the JSON Schema of the request body and exit without calling the API
  -h, --help                       help for post-dashboard
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

* [gf dashboards](gf_dashboards.md)	 - Dashboards API

