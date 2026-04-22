## gf library-elements update-library-element

Updates library element

### Synopsis

Updates library element

Updates an existing library element identified by uid.

Body schema (PatchLibraryElementCommand):
{
  "folderId": number,
  "folderUid": string,
  "kind": number,
  "model": any,
  "name": string,
  "uid": string,
  "version": number
}
  folderId                 ID of the folder where the library element is stored. Deprecated: use FolderUID instead
  folderUid                UID of the folder where the library element is stored.
  kind                     Kind of element to create, Use 1 for library panels or 2 for c. Description: 1 - library panels, enum: 1
  model                    The JSON model for the library element.
  name                     Name of the library element.
  version                  Version of the library element you are updating.

```
gf library-elements update-library-element [flags]
```

### Options

```
      --body string                  The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema     Print the JSON Schema of the request body and exit without calling the API
  -h, --help                         help for update-library-element
      --library-element-uid string   LibraryElementUID
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

