## gf users search-users-with-paging

Gets users with paging

### Synopsis

Gets users with paging

Response schema (SearchUsersWithPagingOK.Payload):
  page                   number
  perPage                number
  totalCount             number
  users                  array<object>
  users[].authLabels     array<string>
  users[].avatarUrl      string
  users[].created        string
  users[].email          string
  users[].id             number
  users[].isAdmin        boolean
  users[].isDisabled     boolean
  users[].isProvisioned  boolean
  users[].lastSeenAt     string
  users[].lastSeenAtAge  string
  users[].login          string
  users[].name           string
  users[].uid            string

```
gf users search-users-with-paging [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for search-users-with-paging
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

* [gf users](gf_users.md)	 - Users API

