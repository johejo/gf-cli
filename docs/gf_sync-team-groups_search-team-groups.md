## gf sync-team-groups search-team-groups

Searches for team groups with optional filtering and pagination

### Synopsis

Searches for team groups with optional filtering and pagination

Response schema (SearchTeamGroupsOK.Payload):
  page                  number
  perPage               number
  teamGroups            array<object>
  teamGroups[].groupId  string
  teamGroups[].orgId    number
  teamGroups[].teamId   number
  teamGroups[].uid      string
  totalCount            number

```
gf sync-team-groups search-team-groups [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for search-team-groups
      --name string                    Filter by exact name match
      --page int                       Default: 1
      --perpage int                    Number of items per page Default: 1000 (default 1000)
      --query string                   If set it will return results where the query value is contained in the name field. Query values with spaces need to be URL encoded.
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --team-id int                    TeamID
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

* [gf sync-team-groups](gf_sync-team-groups.md)	 - Sync team groups API

