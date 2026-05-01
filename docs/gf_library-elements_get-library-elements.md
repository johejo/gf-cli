## gf library-elements get-library-elements

Gets all library elements

### Synopsis

Gets all library elements

Returns a list of all library elements the authenticated user has permission to view. Use the `perPage` query parameter to control the maximum number of library elements returned; the default limit is `100`. You can also use the `page` query parameter to fetch library elements from any page other than the first one.

Response schema (GetLibraryElementsOK.Payload):
  result                                      object
  result.elements                             array<object>
  result.elements[].description               string
  result.elements[].folderId                  number         Deprecated: use FolderUID instead
  result.elements[].folderUid                 string
  result.elements[].id                        number
  result.elements[].kind                      number
  result.elements[].meta                      object
  result.elements[].meta.connectedDashboards  number
  result.elements[].meta.created              string
  result.elements[].meta.createdBy            object
  result.elements[].meta.createdBy.avatarUrl  string
  result.elements[].meta.createdBy.id         number
  result.elements[].meta.createdBy.name       string
  result.elements[].meta.folderName           string
  result.elements[].meta.folderUid            string
  result.elements[].meta.updated              string
  result.elements[].meta.updatedBy            object
  result.elements[].meta.updatedBy.avatarUrl  string
  result.elements[].meta.updatedBy.id         number
  result.elements[].meta.updatedBy.name       string
  result.elements[].model                     object
  result.elements[].name                      string
  result.elements[].orgId                     number
  result.elements[].schemaVersion             number
  result.elements[].type                      string
  result.elements[].uid                       string
  result.elements[].version                   number
  result.page                                 number
  result.perPage                              number
  result.totalCount                           number

```
gf library-elements get-library-elements [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --exclude-uid string             Element UID to exclude from search results.
      --folder-filter string           A comma separated list of folder ID(s) to filter the elements by. Deprecated: Use FolderFilterUIDs instead.
      --folder-filter-uids string      A comma separated list of folder UID(s) to filter the elements by.
  -h, --help                           help for get-library-elements
      --kind int                       Kind of element to search for.
      --page int                       The page for a set of records, given that only perPage records are returned at a time. Numbering starts at 1. Default: 1
      --per-page int                   The number of results per page. Default: 100
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --search-string string           Part of the name or description searched for.
      --sort-direction string          Sort order of elements.
      --type-filter string             A comma separated list of types to filter the elements by
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

