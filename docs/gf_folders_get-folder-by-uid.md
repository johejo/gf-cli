## gf folders get-folder-by-uid

Gets folder by uid

### Synopsis

Gets folder by uid

Use: /apis/folder.grafana.app/v1/namespaces/{ns}/folders/{folder_uid}

Response schema (GetFolderByUIDOK.Payload):
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
gf folders get-folder-by-uid [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --folder-uid string              FolderUID
  -h, --help                           help for get-folder-by-uid
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

