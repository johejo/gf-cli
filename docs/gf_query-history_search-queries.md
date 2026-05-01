## gf query-history search-queries

Queries history search

### Synopsis

Queries history search

Returns a list of queries in the query history that matches the search criteria. Query history search supports pagination. Use the `limit` parameter to control the maximum number of queries returned; the default limit is 100. You can also use the `page` query parameter to fetch queries from any page other than the first one.

```
gf query-history search-queries [flags]
```

### Options

```
      --datasource-uid strings         List of data source UIDs to search for
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --from int                       From range for the query history search
  -h, --help                           help for search-queries
      --limit int                      Limit the number of returned results
      --only-starred                   Flag indicating if only starred queries should be returned
      --page int                       Use this parameter to access hits beyond limit. Numbering starts at 1. limit param acts as page size.
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --search-string string           Text inside query or comments that is searched for
      --sort string                    Sort method Default: "time-desc"
      --to int                         To range for the query history search
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

* [gf query-history](gf_query-history.md)	 - Query history API

