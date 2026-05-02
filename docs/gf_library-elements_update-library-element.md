## gf library-elements update-library-element

Updates library element

### Synopsis

Updates library element

Updates an existing library element identified by uid.

Body schema (PatchLibraryElementCommand):
  folderId   number  ID of the folder where the library element is stored.
                     Deprecated: use FolderUID instead
  folderUid  string  UID of the folder where the library element is stored.
  kind       number  Kind of element to create, Use 1 for library panels or 2 for c.
                     1 - library panels, enum: 1
  model      object  The JSON model for the library element.
  name       string  Name of the library element.
  uid        string
  version    number  Version of the library element you are updating.

Response schema (UpdateLibraryElementOK.Payload):
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
gf library-elements update-library-element [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-library-element
      --library-element-uid string     LibraryElementUID [required]
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

