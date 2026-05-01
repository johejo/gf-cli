## gf folders create-folder

Creates folder

### Synopsis

Creates folder

If nested folders are enabled then it additionally expects the parent folder UID.

Use: /apis/folder.grafana.app/v1/namespaces/{ns}/folders/{folder_uid}

Body schema (CreateFolderCommand):
  description  string
  parentUid    string
  title        string
  uid          string

Response schema (CreateFolderOK.Payload):
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
gf folders create-folder [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for create-folder
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

* [gf folders](gf_folders.md)	 - Folders API

