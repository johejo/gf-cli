## gf library-elements get-library-elements



```
gf library-elements get-library-elements [flags]
```

### Options

```
      --exclude-uid string      Element UID to exclude from search results.
      --folder-filter string    A comma separated list of folder ID(s) to filter the elements by.
  -h, --help                    help for get-library-elements
      --kind int                Kind of element to search for.
      --page int                The page for a set of records, given that only perPage records are returned at a time. Numbering starts at 1. Default: 1
      --per-page int            The number of results per page. Default: 100
      --raw                     Print the raw HTTP response body instead of the decoded payload
      --search-string string    Part of the name or description searched for.
      --sort-direction string   Sort order of elements.
      --type-filter string      A comma separated list of types to filter the elements by
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

* [gf library-elements](gf_library-elements.md)	 - Library elements API

