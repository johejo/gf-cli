## gf library-elements get-library-element-by-uid

Gets library element by UID

### Synopsis

Gets library element by UID

Returns a library element with the given UID.

Response schema (GetLibraryElementByUIDOK.Payload):
  result                           object
  result.description               string
  result.folderId                  number  Deprecated: use FolderUID instead
  result.folderUid                 string
  result.id                        number
  result.kind                      number
  result.meta                      object
  result.meta.connectedDashboards  number
  result.meta.created              string
  result.meta.createdBy            object
  result.meta.createdBy.avatarUrl  string
  result.meta.createdBy.id         number
  result.meta.createdBy.name       string
  result.meta.folderName           string
  result.meta.folderUid            string
  result.meta.updated              string
  result.meta.updatedBy            object
  result.meta.updatedBy.avatarUrl  string
  result.meta.updatedBy.id         number
  result.meta.updatedBy.name       string
  result.model                     object
  result.name                      string
  result.orgId                     number
  result.schemaVersion             number
  result.type                      string
  result.uid                       string
  result.version                   number

```
gf library-elements get-library-element-by-uid [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-library-element-by-uid
      --library-element-uid string     LibraryElementUID
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

* [gf library-elements](gf_library-elements.md)	 - Library elements API

