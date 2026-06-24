## gf folders move-folder

Moves folder

### Synopsis

Moves folder

Use: /apis/folder.grafana.app/v1/namespaces/{ns}/folders/{folder_uid}, Changing the parent folder annotation

Body schema (MoveFolderCommand):
  parentUid  string

Response schema (MoveFolderOK.Payload):
  accessControl  map<string, boolean>
  canAdmin       boolean
  canDelete      boolean
  canEdit        boolean
  canSave        boolean
  created        string
  createdBy      string
  hasAcl         boolean
  id             number                Deprecated: use UID instead
  managedBy      string
  orgId          number
  parentUid      string                only used if nested folders are enabled
  parents        array<object>         the parent folders starting from the root going down
  title          string
  uid            string
  updated        string
  updatedBy      string
  url            string
  version        number

```
gf folders move-folder [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --folder-uid string              FolderUID [required]
  -h, --help                           help for move-folder
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

* [gf folders](gf_folders.md)	 - Folders API

