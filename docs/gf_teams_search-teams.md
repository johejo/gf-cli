## gf teams search-teams

Teams search with paging

### Synopsis

Teams search with paging

Response schema (SearchTeamsOK.Payload):
  page                   number
  perPage                number
  teams                  array<object>
  teams[].accessControl  map<string, boolean>
  teams[].avatarUrl      string
  teams[].email          string
  teams[].externalUID    string
  teams[].id             number                REQUIRED
                                               @deprecated Use UID instead
  teams[].isProvisioned  boolean               REQUIRED
  teams[].memberCount    number                REQUIRED
  teams[].name           string                REQUIRED
  teams[].orgId          number                REQUIRED
  teams[].permission     number
  teams[].uid            string                REQUIRED
  totalCount             number

```
gf teams search-teams [flags]
```

### Options

```
      --accesscontrol                  Accesscontrol
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for search-teams
      --name string                    Name
      --page int                       Page (default 1)
      --perpage int                    Number of items per page The totalCount field in the response can be used for pagination list E.g. if totalCount is equal to 100 teams and the perpage parameter is set to 10 then there are 10 pages of teams. (default 1000)
      --query string                   If set it will return results where the query value is contained in the name field. Query values with spaces need to be URL encoded.
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --sort string                    Sort
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf teams](gf_teams.md)	 - Teams API

