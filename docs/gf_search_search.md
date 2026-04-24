## gf search search



```
gf search search [flags]
```

### Options

```
      --dashboard-ids dashboardUIDs    List of dashboard id’s to search for This is deprecated: users should use the dashboardUIDs query parameter instead (default [])
      --dashboard-uids strings         List of dashboard uid’s to search for
      --deleted                        Flag indicating if only soft deleted Dashboards should be returned
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --folder-ids 0                   List of folder id’s to search in for dashboards If it's 0 then it will query for the top level folders This is deprecated: users should use the `folderUIDs` query parameter instead (default [])
      --folder-uids strings            List of folder UID’s to search in for dashboards If it's an empty string then it will query for the top level folders
  -h, --help                           help for search
      --limit int                      Limit the number of returned results (max 5000)
      --page int                       Use this parameter to access hits beyond limit. Numbering starts at 1. limit param acts as page size. Only available in Grafana v6.2+.
      --permission Edit                Set to Edit to return dashboards/folders that the user can edit Default: "View"
      --query string                   Search Query
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --sort string                    Sort method; for listing all the possible sort methods use the search sorting endpoint. Default: "alpha-asc"
      --starred                        Flag indicating if only starred Dashboards should be returned
      --tag strings                    List of tags to search for
      --type string                    Type to search for, dash-folder or dash-db
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

* [gf search](gf_search.md)	 - Search API

