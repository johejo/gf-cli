## gf service-accounts search-org-service-accounts-with-paging

Searches service accounts with paging

### Synopsis

Searches service accounts with paging

Required permissions (See note in the [introduction](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api) for an explanation): action: `serviceaccounts:read` scope: `serviceaccounts:*`

Response schema (SearchOrgServiceAccountsWithPagingOK.Payload):
  page                             number
  perPage                          number
  serviceAccounts                  array<object>
  serviceAccounts[].accessControl  map<string, boolean>
  serviceAccounts[].avatarUrl      string
  serviceAccounts[].id             number
  serviceAccounts[].isDisabled     boolean
  serviceAccounts[].isExternal     boolean
  serviceAccounts[].login          string
  serviceAccounts[].name           string
  serviceAccounts[].orgId          number
  serviceAccounts[].role           string
  serviceAccounts[].tokens         number
  serviceAccounts[].uid            string
  totalCount                       number                It can be used for pagination of the user list
                                                         E.g. if totalCount is equal to 100 users and
                                                         the perpage parameter is set to 10 then there are 10 pages of users.

```
gf service-accounts search-org-service-accounts-with-paging [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --disabled                       Disabled
      --expired-tokens                 ExpiredTokens
  -h, --help                           help for search-org-service-accounts-with-paging
      --page int                       The default value is 1.
      --perpage int                    The default value is 1000. (default 1000)
      --query string                   It will return results where the query value is contained in one of the name. Query values with spaces need to be URL encoded.
      --raw                            Print the raw HTTP response body instead of the decoded payload
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

* [gf service-accounts](gf_service-accounts.md)	 - Service accounts API

